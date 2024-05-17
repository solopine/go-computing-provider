package computing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/robfig/cron/v3"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/models"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var deployingChan = make(chan models.Job)
var TaskMap sync.Map

type CronTask struct {
	nodeId       string
	ownerAddress string
}

func NewCronTask(nodeId string) *CronTask {
	ownerAddress, _, err := GetOwnerAddressAndWorkerAddress()
	if err != nil {
		logs.GetLogger().Errorf("get owner address failed, error: %v", err)
		return nil
	}
	return &CronTask{nodeId: nodeId, ownerAddress: ownerAddress}
}

func (task *CronTask) RunTask() {
	addNodeLabel()
	checkJobStatus(task.ownerAddress)
	task.checkCollateralBalance()
	task.cleanAbnormalDeployment()
	task.setFailedUbiTaskStatus()
	task.watchNameSpaceForDeleted()
	task.updateUbiTaskReward()
}

func checkJobStatus(ownerAddress string) {
	go func() {
		for {
			select {
			case job := <-deployingChan:
				TaskMap.Store(job.Uuid, &job)
			case <-time.After(3 * time.Second):
				TaskMap.Range(func(key, value any) bool {
					jobUuid := key.(string)
					job := value.(*models.Job)
					if reportJobStatus(jobUuid, job.Status, ownerAddress) && job.Status == models.DEPLOY_TO_K8S {
						TaskMap.Delete(jobUuid)
					}
					return true
				})
			}
		}
	}()
}

func (task *CronTask) watchNameSpaceForDeleted() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/50 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("watchNameSpaceForDeleted catch panic error: %+v", err)
			}
		}()
		service := NewK8sService()
		namespaces, err := service.ListNamespace(context.TODO())
		if err != nil {
			logs.GetLogger().Errorf("Failed get all namespace, error: %+v", err)
			return
		}

		for _, namespace := range namespaces {
			getPods, err := service.GetPods(namespace, "")
			if err != nil {
				logs.GetLogger().Errorf("Failed get pods form namespace,namepace: %s, error: %+v", namespace, err)
				continue
			}
			if !getPods && (strings.HasPrefix(namespace, constants.K8S_NAMESPACE_NAME_PREFIX) || strings.HasPrefix(namespace, "ubi-task")) {
				if err = service.DeleteNameSpace(context.TODO(), namespace); err != nil {
					logs.GetLogger().Errorf("Failed delete namespace, namepace: %s, error: %+v", namespace, err)
				}
			}
		}
		NewDockerService().CleanResource()
	})
	c.Start()
}

func (task *CronTask) watchExpiredTask() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/5 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("watchExpiredTask catch panic error: %+v", err)
			}
		}()

		jobList, err := NewJobService().GetJobList()
		if err != nil {
			logs.GetLogger().Errorf("Failed watchExpiredTask get job data, error: %+v", err)
			return
		}

		var deleteSpaceIds []string
		for _, job := range jobList {
			namespace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(job.WalletAddress)
			if len(strings.TrimSpace(job.TaskUuid)) == 0 {
				taskStatus, err := checkTaskStatusByHub(job.TaskUuid, task.nodeId)
				if err != nil {
					logs.GetLogger().Errorf("Failed check task status by Orchestrator service, error: %+v", err)
					return
				}
				if strings.Contains(taskStatus, "no task found") {
					logs.GetLogger().Infof("task_uuid: %s, task not found on the orchestrator service, starting to delete it.", job.TaskUuid)
					deleteJob(namespace, job.SpaceUuid)
					deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
					continue
				}
				if strings.Contains(taskStatus, "Terminated") || strings.Contains(taskStatus, "Terminated") ||
					strings.Contains(taskStatus, "Cancelled") || strings.Contains(taskStatus, "Failed") {
					logs.GetLogger().Infof("task_uuid: %s, current status is %s, starting to delete it.", job.TaskUuid, taskStatus)
					if err = deleteJob(namespace, job.SpaceUuid); err == nil {
						deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
						continue
					}
				}
			}

			if time.Now().Unix() > job.ExpireTime {
				expireTimeStr := time.Unix(job.ExpireTime, 0).Format("2006-01-02 15:04:05")
				logs.GetLogger().Infof("<timer-task> spaceUuid: %s, expireTime: %s. the job starting terminated", job.SpaceUuid, expireTimeStr)
				if err = deleteJob(namespace, job.SpaceUuid); err == nil {
					deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
					continue
				}
			}

			k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(job.WalletAddress)
			deployName := constants.K8S_DEPLOY_NAME_PREFIX + job.SpaceUuid
			service := NewK8sService()
			if _, err = service.k8sClient.AppsV1().Deployments(k8sNameSpace).Get(context.TODO(), deployName, metav1.GetOptions{}); err != nil && errors.IsNotFound(err) {
				deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
				continue
			}
		}

	})
	c.Start()
}

func (task *CronTask) checkCollateralBalance() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0/15 * * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [checkCollateralBalance], error: %+v", err)
			}
		}()

		url := fmt.Sprintf("%s/cp/collateral/%s", conf.GetConfig().HUB.ServerUrl, task.ownerAddress)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			logs.GetLogger().Errorf("create req failed: %+v", err)
			return
		}
		req.Header.Set("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logs.GetLogger().Errorf("send req failed: %+v", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logs.GetLogger().Errorf("read response failed: %+v", err)
			return
		}
		var collateral struct {
			Data struct {
				Balance float64 `json:"balance"`
			} `json:"data"`
			Message string `json:"message"`
			Status  string `json:"status"`
		}
		err = json.Unmarshal(body, &collateral)
		if err != nil {
			logs.GetLogger().Errorf("json conversion failed: %+v", err)
			return
		}

		result := collateral.Data.Balance / 1e18
		result = math.Round(result*1000) / 1000
		if result <= conf.GetConfig().HUB.BalanceThreshold {
			logs.GetLogger().Warnf("No sufficient collateral Balance, the current collateral balance is: %0.3f. Please run: computing-provider collateral [fromWalletAddress] [amount]", result)
		}
	})
	c.Start()
}

func (task *CronTask) cleanAbnormalDeployment() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* 0/30 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [cleanAbnormalDeployment], error: %+v", err)
			}
		}()

		k8sService := NewK8sService()
		namespaces, err := k8sService.ListNamespace(context.TODO())
		if err != nil {
			logs.GetLogger().Errorf("Failed get all namespace, error: %+v", err)
			return
		}

		for _, namespace := range namespaces {
			if strings.HasPrefix(namespace, constants.K8S_NAMESPACE_NAME_PREFIX) {
				deployments, err := k8sService.k8sClient.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
				if err != nil {
					logs.GetLogger().Errorf("Error getting deployments in namespace %s: %v\n", namespace, err)
					continue
				}

				for _, deployment := range deployments.Items {
					creationTimestamp := deployment.ObjectMeta.CreationTimestamp.Time
					currentTime := time.Now()
					age := currentTime.Sub(creationTimestamp)
					if (deployment.Status.AvailableReplicas == 0 && age.Hours() >= 2) || age.Hours() > 24*15 {
						logs.GetLogger().Infof("Cleaning up deployment %s in namespace %s", deployment.Name, namespace)
						err := k8sService.k8sClient.AppsV1().Deployments(namespace).Delete(context.TODO(), deployment.Name, metav1.DeleteOptions{})
						if err != nil {
							if errors.IsNotFound(err) {
								logs.GetLogger().Errorf("Deployment %s not found. Ignoring", deployment.Name)
							} else {
								logs.GetLogger().Errorf("Error deleting deployment %s: %v", deployment.Name, err)
							}
						} else {
							logs.GetLogger().Errorf("Deployment %s deleted successfully.", deployment.Name)
						}
					}
				}
			}
		}
	})
	c.Start()
}

func (task *CronTask) setFailedUbiTaskStatus() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/10 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [cleanAbnormalDeployment], error: %+v", err)
			}
		}()

		var taskList []models.TaskEntity
		oneHourAgo := time.Now().Add(-1 * time.Hour).Unix()
		err := NewTaskService().Model(&models.TaskEntity{}).Where("status !=? and status !=? and create_time <?", models.TASK_SUCCESS_STATUS, models.TASK_FAILED_STATUS, oneHourAgo).Or("tx_hash==''").Find(&taskList).Error
		if err != nil {
			logs.GetLogger().Errorf("Failed get task list, error: %+v", err)
			return
		}

		for _, entity := range taskList {
			ubiTask := entity

			JobName := strings.ToLower(ubiTask.ZkType) + "-" + strconv.Itoa(int(ubiTask.Id))
			k8sNameSpace := "ubi-task-" + strconv.Itoa(int(ubiTask.Id))
			service := NewK8sService()
			service.k8sClient.BatchV1().Jobs(k8sNameSpace).Delete(context.TODO(), JobName, metav1.DeleteOptions{})
			ubiTask.Status = models.TASK_FAILED_STATUS
			NewTaskService().SaveTaskEntity(&ubiTask)
		}
	})
	c.Start()
}

func (task *CronTask) updateUbiTaskReward() {
	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	nodeId := GetNodeId(cpRepoPath)

	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/10 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [updateUbiTaskReward], error: %+v", err)
			}
		}()

		taskList, err := NewTaskService().GetTaskListNoReward()
		if err != nil {
			logs.GetLogger().Errorf("Failed get task list, error: %+v", err)
			return
		}

		for _, entity := range taskList {
			ubiTask := entity
			reward, err := getReward(nodeId, strconv.Itoa(int(ubiTask.Id)))
			if err != nil {
				logs.GetLogger().Errorf("get ubi task reward failed, taskId: %d, error: %v", ubiTask.Id, err)
				continue
			}
			ubiTask.Reward = reward
			NewTaskService().SaveTaskEntity(ubiTask)
		}
	})
	c.Start()
}

func addNodeLabel() {
	k8sService := NewK8sService()
	nodes, err := k8sService.k8sClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	nodeGpuInfoMap, err := k8sService.GetResourceExporterPodLog(context.TODO())
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Infof("collect all node: %d", len(nodes.Items))
	for _, node := range nodes.Items {
		cpNode := node
		if collectInfo, ok := nodeGpuInfoMap[cpNode.Name]; ok {
			for _, detail := range collectInfo.Gpu.Details {
				if err = k8sService.AddNodeLabel(cpNode.Name, detail.ProductName); err != nil {
					logs.GetLogger().Errorf("add node label, nodeName %s, gpuName: %s, error: %+v", cpNode.Name, detail.ProductName, err)
					continue
				}
			}
			err := k8sService.AddNodeLabel(cpNode.Name, collectInfo.CpuName)
			if err != nil {
				logs.GetLogger().Errorf("nodeName: %s, error: %v", cpNode.Name, err)
			}
		}
	}
}

func checkTaskStatusByHub(taskUuid, nodeId string) (string, error) {
	url := fmt.Sprintf("%s/check_task_status_with_node_id/%s/%s", conf.GetConfig().HUB.ServerUrl, taskUuid, nodeId)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("offset", "0")
	req.Header.Add("limit", "10")
	req.Header.Add("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var taskStatus struct {
		Data struct {
			JobStatus  string `json:"job_status"`
			TaskStatus string `json:"task_status"`
		} `json:"data"`
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	err = json.Unmarshal(respBody, &taskStatus)
	if err != nil {
		logs.GetLogger().Errorf("check_task_status_with_node_id resp: %s", string(respBody))
		return "", err
	}
	if taskStatus.Status == "failed" {
		return taskStatus.Message, nil
	}
	return taskStatus.Status, nil
}

func reportJobStatus(jobUuid string, deployStatus int, ownerAddress string) bool {
	var job = new(models.JobEntity)
	job.JobUuid = jobUuid
	job.Status = deployStatus
	if err := NewJobService().UpdateJobEntityByJobUuid(job); err != nil {
		logs.GetLogger().Errorf("update job info by jobUuid failed, error: %v", err)
	}

	reqParam := map[string]interface{}{
		"job_uuid":       jobUuid,
		"status":         models.GetDeployStatusStr(deployStatus),
		"public_address": ownerAddress,
	}

	payload, err := json.Marshal(reqParam)
	if err != nil {
		logs.GetLogger().Errorf("Failed convert to json, error: %+v", err)
		return false
	}

	client := &http.Client{}
	url := conf.GetConfig().HUB.ServerUrl + "/job/status"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		logs.GetLogger().Errorf("Error creating request: %v", err)
		return false
	}
	req.Header.Set("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logs.GetLogger().Errorf("Failed send a request, error: %+v", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	logs.GetLogger().Debugf("report job status successfully. uuid: %s, status: %s", jobUuid, models.GetDeployStatusStr(deployStatus))
	return true
}
