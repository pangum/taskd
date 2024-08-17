package mysql

import (
	"time"

	"github.com/dnsabre/core/kernel"
	"github.com/goexl/id"
	"github.com/pangum/db"
	"github.com/pangum/taskd/internal/internal/internal/column"
	"github.com/pangum/taskd/internal/internal/model"
	"github.com/pangum/taskd/internal/internal/repository/internal/get"
	"xorm.io/builder"
)

type Task struct {
	id id.Generator
	db *db.Engine
}

func NewTask(database get.Database) *Task {
	return &Task{
		id: database.Id,
		db: database.DB,
	}
}

func (t *Task) Add(task *model.Task) (int64, error) {
	if 0 == task.Id {
		task.Id = t.id.Next().Value()
	}

	return t.db.InsertOne(task)
}

func (t *Task) Get(task *model.Task, columns ...string) (bool, error) {
	return t.db.Cols(columns...).Get(task)
}

func (t *Task) GetsRunnable(
	count int, times uint8, maximum time.Duration, excludes ...*model.Task,
) (tasks *[]*model.Task, err error) {
	required := builder.Lte{
		column.Times.String(): times, // 还没有到最大重试次数
	}

	// 可被运行的条件一：运行时间已到且状态是被认可可被重新执行
	created := builder.Lte{
		column.Next.String(): time.Now(), // 运行时间已到
	}.And(builder.Eq{
		column.Status.String(): kernel.TaskStatus_TASK_STATUS_CREATED, // 刚创建的任务
	}.Or(builder.Eq{
		column.Status.String(): kernel.TaskStatus_TASK_STATUS_FAILED, // 已经处于错误状态的任务
	}))
	// 可被运行的条件二：任务已完成执行，但需要重启执行（可被循环执行的任务，达到下一次执行的条件）
	recyclable := builder.Eq{ // 已经完成的任务，需要重新执行
		column.Status.String(): kernel.TaskStatus_TASK_STATUS_SUCCESS, // 已完成
	}.And(builder.Eq{
		column.Times.String(): 0, // 次数被重置
	})
	// 可被运行的条件三：因各种问题中断执行
	interrupted := builder.Lte{
		column.Next.String(): time.Now().Add(-maximum), // 超过配置时间的最大运行时间段
	}.And(builder.Eq{
		column.Status.String(): kernel.TaskStatus_TASK_STATUS_RUNNING, // 运行中
	}.Or(builder.Eq{
		column.Status.String(): kernel.TaskStatus_TASK_STATUS_RETRYING, // 重试中
	}))

	// 排除
	excludeTasks := builder.NewCond()
	for _, exclude := range excludes {
		excludeTasks = excludeTasks.And(builder.Neq{
			column.Id.String(): exclude.Id,
		})
	}

	entities := make([]*model.Task, 0)
	tasks = &entities
	cond := required.And(created.Or(recyclable).Or(interrupted)).And(excludeTasks)
	err = t.db.Where(cond).Limit(count).OrderBy(column.Created.Asc()).Find(tasks)

	return
}

func (t *Task) Update(task *model.Task, columns ...string) (int64, error) {
	return t.db.ID(task.Id).MustCols(columns...).Update(task)
}

func (t *Task) Delete(task *model.Task) (int64, error) {
	return t.db.Delete(task)
}
