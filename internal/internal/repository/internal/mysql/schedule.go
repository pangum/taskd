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

func (s *Schedule) Add(runtime *model.Runtime, runtimes ...*model.Runtime) (successes *[]*model.Tasker, err error) {
	models := make([]*model.Tasker, 0, 1+len(runtimes))
	saves := append([]*model.Runtime{runtime}, runtimes...)
	if _, err = s.tx.Do(s.add(&saves, &models)); nil == err {
		successes = &models
	}

	return
}

func (s *Schedule) Get(schedule *model.Schedule, columns ...string) (bool, error) {
	return s.db.Cols(columns...).Get(schedule)
}

func (s *Schedule) Update(schedule *model.Schedule, columns ...string) (int64, error) {
	return s.db.ID(schedule.Id).MustCols(columns...).Update(schedule)
}

func (s *Schedule) Delete(schedule *model.Schedule) (int64, error) {
	return s.tx.Do(s.delete(schedule))
}

func (s *Schedule) delete(schedule *model.Schedule) func(session *db.Session) (int64, error) {
	return func(session *db.Session) (affected int64, err error) {
		deleted := new(model.Schedule)
		deleted.Id = schedule.Id
		if ads, dse := session.Delete(deleted); nil != dse { // 删除计划本身
			err = dse
		} else if adt, dte := s.deleteTask(session, schedule); nil != dte { // 删除对应的任务
			err = dte
		} else {
			affected = ads + adt
		}

		return
	}
}

func (s *Schedule) add(runtimes *[]*model.Runtime, successes *[]*model.Tasker) func(session *db.Session) (int64, error) {
	return func(session *db.Session) (affected int64, err error) {
		schedules := make([]any, 0, len(*runtimes))
		for _, runtime := range *runtimes {
			schedule := &runtime.Schedule
			if 0 == schedule.Id {
				schedule.Id = s.id.Next().Value()
			}
			schedules = append(schedules, schedule)
		}

		if ais, ie := session.Insert(schedules...); nil != ie {
			err = ie
		} else if aat, ate := s.addTasks(session, runtimes, successes); nil != ate {
			err = ate
		} else {
			affected = ais + aat
		}

		return
	}
}

func (s *Schedule) addTasks(
	session *db.Session,
	runtimes *[]*model.Runtime, successes *[]*model.Tasker,
) (affected int64, err error) {
	tasks := make([]any, 0, len(*runtimes))
	for _, runtime := range *runtimes {
		_task := new(model.Task)
		_task.Id = s.id.Next().Value()
		_task.Schedule = runtime.Id
		_task.Next = runtime.Next
		_task.Status = task.StatusCreated

		now := time.Now()
		_task.Start = now
		_task.Stop = now.Add(runtime.Elapsed)

		tasks = append(tasks, _task)
	}

	if affected, err = session.Insert(tasks...); nil == err {
		s.parseTasks(&tasks, runtimes, successes)
	}

	return
}

func (s *Schedule) deleteTask(session *db.Session, schedule *model.Schedule) (affected int64, err error) {
	deleted := new(model.Task)
	deleted.Schedule = schedule.Id
	affected, err = session.Delete(deleted)

	return
}

func (s *Schedule) parseTasks(tasks *[]any, runtimes *[]*model.Runtime, successes *[]*model.Tasker) {
	for index, _task := range *tasks {
		if converted, ok := _task.(*model.Task); ok {
			success := new(model.Tasker)
			success.Id = converted.Id
			success.Start = converted.Start
			success.Next = converted.Next
			success.Stop = converted.Stop
			success.Retries = converted.Times
			success.Status = task.StatusCreated

			schedule := (*runtimes)[index]
			success.Target = schedule.Target
			success.Type = schedule.Type
			success.Subtype = schedule.Subtype
			success.Elapsed = schedule.Elapsed
			success.Data = schedule.Data

			*successes = append(*successes, success)
		}
	}
}
