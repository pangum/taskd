package mysql

import (
	"time"

	"github.com/goexl/id"
	"github.com/goexl/task"
	"github.com/pangum/db"
	"github.com/pangum/taskd/internal/internal/model"
	"github.com/pangum/taskd/internal/internal/repository/internal/get"
)

type Schedule struct {
	id id.Generator
	db *db.Engine
	tx *db.Transaction
}

func NewSchedule(database get.Transaction) *Schedule {
	return &Schedule{
		id: database.Id,
		db: database.DB,
		tx: database.Transaction,
	}
}

func (s *Schedule) Add(schedule *model.Schedule, next time.Time) (runnable *model.Tasker, err error) {
	runnable = new(model.Tasker)
	err = s.tx.Do(s.add(schedule, runnable, next))

	return
}

func (s *Schedule) Get(schedule *model.Schedule, columns ...string) (bool, error) {
	return s.db.Cols(columns...).Get(schedule)
}

func (s *Schedule) Update(schedule *model.Schedule, columns ...string) (int64, error) {
	return s.db.ID(schedule.Id).MustCols(columns...).Update(schedule)
}

func (s *Schedule) Delete(schedule *model.Schedule) (affected int64, err error) {
	if err = s.tx.Do(s.delete(schedule)); nil == err {
		affected = 1
	}

	return
}

func (s *Schedule) delete(schedule *model.Schedule) func(session *db.Session) error {
	return func(session *db.Session) (err error) {
		deleted := new(model.Schedule)
		deleted.Id = schedule.Id
		if _, dse := session.Delete(deleted); nil != dse { // 删除计划本身
			err = dse
		} else if _, dte := s.deleteTask(session, schedule); nil != dte { // 删除对应的任务
			err = dte
		}

		return
	}
}

func (s *Schedule) add(
	schedule *model.Schedule,
	runnable *model.Tasker, next time.Time,
) func(session *db.Session) error {
	return func(session *db.Session) (err error) {
		if 0 == schedule.Id {
			schedule.Id = s.id.Next().Value()
		}

		if _, ie := session.Insert(schedule); nil != ie {
			err = ie
		} else {
			_, err = s.addTask(session, schedule, runnable, next)
		}

		return
	}
}

func (s *Schedule) addTask(
	session *db.Session,
	schedule *model.Schedule,
	runnable *model.Tasker, next time.Time,
) (affected int64, err error) {
	saved := new(model.Task)
	saved.Id = s.id.Next().Value()
	saved.Schedule = schedule.Id
	saved.Next = next
	saved.Status = task.StatusCreated

	now := time.Now()
	saved.Start = now
	saved.Stop = now.Add(schedule.Elapsed)

	if affected, err = session.Insert(saved); nil == err {
		runnable.Id = saved.Id
		runnable.Start = saved.Start
		runnable.Next = saved.Next
		runnable.Stop = saved.Stop
		runnable.Retries = saved.Times
		runnable.Status = task.StatusCreated

		runnable.Target = schedule.Target
		runnable.Type = schedule.Type
		runnable.Subtype = schedule.Subtype
		runnable.Elapsed = schedule.Elapsed
		runnable.Data = schedule.Data
	}

	return
}

func (s *Schedule) deleteTask(session *db.Session, schedule *model.Schedule) (affected int64, err error) {
	deleted := new(model.Task)
	deleted.Schedule = schedule.Id
	affected, err = session.Delete(deleted)

	return
}
