package models

const (
	SOURCE_TYPE_CPU = 0
	SOURCE_TYPE_GPU = 1
)

const (
	ZK_TYPE_FIL_32  = "fil-c2-32G"
	ZK_TYPE_FIL_512 = "fil-c2-512M"
	ZK_TYPE_ALEO    = "aleo_proof"
)

const (
	TASK_RECEIVED_STATUS = iota + 1
	TASK_RUNNING_STATUS
	TASK_SUCCESS_STATUS
	TASK_FAILED_STATUS
)

func TaskStatusStr(status int) string {
	var statusStr string
	switch status {
	case TASK_RECEIVED_STATUS:
		statusStr = "received"
	case TASK_RUNNING_STATUS:
		statusStr = "running"
	case TASK_SUCCESS_STATUS:
		statusStr = "success"
	case TASK_FAILED_STATUS:
		statusStr = "failed"
	}
	return statusStr
}

func GetSourceTypeStr(resourceType int) string {
	switch resourceType {
	case SOURCE_TYPE_CPU:
		return "CPU"
	case SOURCE_TYPE_GPU:
		return "GPU"
	}
	return ""
}

type TaskEntity struct {
	Id           int64   `json:"id" gorm:"primaryKey;id"`
	ZkType       string  `json:"zk_type" gorm:"zk_type"`
	Name         string  `json:"name" gorm:"name"`
	Contract     string  `json:"contract" gorm:"name"`
	ResourceType int     `json:"resource_type" gorm:"resource_type"` // 1
	InputParam   string  `json:"input_param" gorm:"input_param"`
	TxHash       string  `json:"tx_hash" gorm:"tx_hash"`
	Status       int     `json:"status" gorm:"status"`
	Reward       float64 `json:"reward" gorm:"reward; default:0"`
	CreateTime   int64   `json:"create_time" gorm:"create_time"`
	EndTime      int64   `json:"end_time" gorm:"end_time"`
	Error        string  `json:"error" gorm:"error"`
}

func (task *TaskEntity) TableName() string {
	return "t_task"
}

const (
	DEPLOY_RECEIVE_JOB = iota + 1
	DEPLOY_DOWNLOAD_SOURCE
	DEPLOY_UPLOAD_RESULT
	DEPLOY_BUILD_IMAGE
	DEPLOY_PUSH_IMAGE
	DEPLOY_PULL_IMAGE
	DEPLOY_TO_K8S
)

func GetDeployStatusStr(deployStatus int) string {
	var statusStr string
	switch deployStatus {
	case DEPLOY_DOWNLOAD_SOURCE:
		statusStr = "downloadSource"
	case DEPLOY_UPLOAD_RESULT:
		statusStr = "uploadResult"
	case DEPLOY_BUILD_IMAGE:
		statusStr = "buildImage"
	case DEPLOY_PUSH_IMAGE:
		statusStr = "pushImage"
	case DEPLOY_PULL_IMAGE:
		statusStr = "pullImage"
	case DEPLOY_TO_K8S:
		statusStr = "deployToK8s"
	}
	return statusStr
}

type JobEntity struct {
	Id              uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Source          string `json:"source" gorm:"source"`
	Name            string `json:"name" gorm:"name"`
	SpaceUuid       string `json:"space_uuid"`
	JobUuid         string `json:"job_uuid"`
	TaskUuid        string `json:"task_uuid"`
	ResourceType    string `json:"resource_type"`
	SpaceType       int    `json:"space_type"` // 0: public; 1: private
	SourceUrl       string `json:"source_url" gorm:"source_url"`
	Hardware        string `json:"hardware" gorm:"hardware"`
	Duration        int    `json:"duration"`
	DeployStatus    int    `json:"deploy_status" gorm:"deploy_status"`
	WalletAddress   string `json:"wallet_address"`
	Status          int    `json:"status" gorm:"status"`
	ResultUrl       string `json:"result_url" gorm:"result_url"`
	RealUrl         string `json:"real_url"`
	K8sDeployName   string `json:"k8s_deploy_name" gorm:"k8s_deploy_name"`
	K8sResourceType string `json:"k8s_resource_type" gorm:"k8s_resource_type"`
	NameSpace       string `json:"name_space" gorm:"name_space"`
	ImageName       string `json:"image_name" gorm:"image_name"`
	BuildLog        string `json:"build_log" gorm:"build_log"`
	ContainerLog    string `json:"container_log" gorm:"container_log"`
	ExpireTime      int64  `json:"expire_time" gorm:"expire_time"`
	CreateTime      int64  `json:"create_time" gorm:"create_time"`
	Error           string `json:"error" gorm:"error"`
}

func (*JobEntity) TableName() string {
	return "t_job"
}
