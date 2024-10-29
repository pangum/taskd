package kernel

import "github.com/pangum/taskd/internal/internal/model"

var _ task.Task = (*TaskSchedule)(nil)

type TaskSchedule struct {
	schedule *model.Schedule
	task     *model.Task
}

func NewTaskSchedule(schedule *model.Schedule, task *model.Task) *TaskSchedule {
	return &TaskSchedule{
		schedule: schedule,
		task:     task,
	}
}
