package computing

import (
	"github.com/google/wire"
	"github.com/swanchain/go-computing-provider/internal/db"
	"github.com/swanchain/go-computing-provider/internal/models"
	"gorm.io/gorm"
	"time"
)

type TaskService struct {
	*gorm.DB
}

func (taskServ TaskService) GetAllTask() (list []*models.TaskEntity, err error) {
	err = taskServ.Model(&models.TaskEntity{}).Order("create_time").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}

func (taskServ TaskService) GetTaskList(taskStatus int) (list []*models.TaskEntity, err error) {
	err = taskServ.Where(&models.TaskEntity{Status: taskStatus}).Order("create_time").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}

func (taskServ TaskService) SaveTaskEntity(task *models.TaskEntity) (err error) {
	if task.Status == models.TASK_FAILED_STATUS || task.Status == models.TASK_SUCCESS_STATUS {
		task.EndTime = time.Now().Unix()
	}
	return taskServ.Save(task).Error
}

func (taskServ TaskService) GetTaskEntity(taskId int64) (*models.TaskEntity, error) {
	var taskEntity models.TaskEntity
	err := taskServ.First(&taskEntity, taskId).Error
	return &taskEntity, err
}

func (taskServ TaskService) GetTaskListNoReward() (list []*models.TaskEntity, err error) {
	err = taskServ.Where("reward ==0").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}

type JobService struct {
	*gorm.DB
}

func (jobServ JobService) SaveJobEntity(job *models.JobEntity) (err error) {
	return jobServ.Debug().Save(job).Error
}

func (jobServ JobService) UpdateJobEntityBySpaceUuid(job *models.JobEntity) (err error) {
	return jobServ.Debug().Where("space_uuid=?", job.SpaceUuid).Updates(job).Error
}

func (jobServ JobService) UpdateJobEntityByJobUuid(job *models.JobEntity) (err error) {
	return jobServ.Debug().Where("job_uuid=?", job.JobUuid).Updates(job).Error
}

func (jobServ JobService) GetJobEntityByTaskUuid(taskUuid string) (models.JobEntity, error) {
	var job models.JobEntity
	err := jobServ.Debug().Where("task_uuid=?", taskUuid).Find(&job).Error
	return job, err
}

func (jobServ JobService) GetJobEntityBySpaceUuid(spaceUuid string) (models.JobEntity, error) {
	var job models.JobEntity
	err := jobServ.Debug().Where("space_uuid=?", spaceUuid).Find(&job).Error
	return job, err
}

func (jobServ JobService) DeleteJobEntityBySpaceUuId(spaceUuid string) error {
	return jobServ.Debug().Where("space_uuid=?", spaceUuid).Delete(&models.JobEntity{}).Error
}

func (jobServ JobService) GetJobList() (list []*models.JobEntity, err error) {
	err = jobServ.Model(&models.JobEntity{}).Find(&list).Error
	return
}

func (jobServ JobService) DeleteJobs(spaceIds []string) (err error) {
	return jobServ.Where("space_uuid in ?", spaceIds).Delete(&models.JobEntity{}).Error
}

var taskSet = wire.NewSet(db.NewDbService, wire.Struct(new(TaskService), "*"))
var jobSet = wire.NewSet(db.NewDbService, wire.Struct(new(JobService), "*"))
