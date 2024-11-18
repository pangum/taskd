package mysql

import (
	"time"

	"github.com/goexl/id"
	"github.com/goexl/task"
	"github.com/pangum/db"
	"github.com/pangum/taskd/internal/internal/internal/column"
	"github.com/pangum/taskd/internal/internal/model"
	"github.com/pangum/taskd/internal/internal/repository/internal/get"
	"xorm.io/builder"
)

type Task struct {
	id id.Generator
	db *db.Engine
	tx *db.Transaction

	table *model.Task
}

func NewTask(tx get.Transaction) *Task {
	return &Task{
		id: tx.Id,
		db: tx.DB,
		tx: tx.Transaction,

		table: new(model.Task),
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

func (t *Task) GetsRunnable(retries uint32, excludes ...*model.Task) (tasks *[]*model.Tasker, err error) {
	now := time.Now()
	required := builder.Lte{
		column.Times.String(): retries, // 还没有到最大重试次数
	}

	// 可被运行的条件一：运行时间已到且状态是被认可可被重新执行
	created := builder.Lte{
		column.Next.String(): now, // 运行时间已到
	}.And(builder.Eq{
		column.Status.String(): task.StatusCreated, // 刚创建的任务
	}.Or(builder.Eq{
		column.Status.String(): task.StatusFailed, // 已经处于错误状态的任务
	}))

	// 可被运行的条件二：任务已完成执行，但需要重启执行（可被循环执行的任务，达到下一次执行的条件）
	recyclable := builder.Eq{ // 已经完成的任务，需要重新执行
		column.Status.String(): task.StatusSuccess, // 已完成
	}.And(builder.Eq{
		column.Times.String(): 0, // 次数被重置
	})

	// 可被运行的条件三：因各种问题中断执行
	interrupted := builder.Lte{
		column.Stop.String(): now, // 超过最大运行时间段
	}.And(builder.Eq{
		column.Status.String(): task.StatusRunning, // 运行中
	}.Or(builder.Eq{
		column.Status.String(): task.StatusRetrying, // 重试中
	}))

	// 排除
	excludeTasks := builder.NewCond()
	for _, exclude := range excludes {
		excludeTasks = excludeTasks.And(builder.Neq{
			column.Id.String(): exclude.Id,
		})
	}

	entities := make([]*model.Tasker, 0)
	tasks = &entities
	cond := required.And(created.Or(recyclable).Or(interrupted)).And(excludeTasks)

	session := t.db.Table(t.table).Where(cond)
	session.Limit(1024) // 最大取1024个数据
	session.Join("INNER", "schedule", "schedule.id = task.schedule")
	err = session.Find(tasks)

	return
}

func (t *Task) Update(task *model.Task, columns ...string) (int64, error) {
	return t.db.ID(task.Id).MustCols(columns...).Update(task)
}

func (t *Task) Archive(task *model.Task) (int64, error) {
	return t.tx.Do(t.delete(task))
}

func (t *Task) Delete(task *model.Task) (int64, error) {
	return t.tx.Do(t.delete(task))
}

func (t *Task) delete(task *model.Task) func(session *db.Session) (int64, error) {
	return func(session *db.Session) (affected int64, err error) {
		deleted := new(model.Task)
		deleted.Id = task.Id
		if adt, dse := session.Delete(deleted); nil != dse { // 删除计划本身
			err = dse
		} else if ads, dte := t.deleteSchedule(session, task); nil != dte { // 删除对应的任务
			err = dte
		} else {
			affected = adt + ads
		}

		return
	}
}

func (t *Task) deleteSchedule(session *db.Session, task *model.Task) (affected int64, err error) {
	deleted := new(model.Schedule)
	deleted.Id = task.Schedule
	affected, err = session.Delete(deleted)

	return
}
