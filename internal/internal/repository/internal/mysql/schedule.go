package mysql

import (
	"github.com/pangum/taskd/internal/internal/kernel"
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
		tx: database.Tx,
	}
}

func (s *Schedule) Add(schedule task.Schedule) (task task.Task, err error) {
	return s.tx.Do(s.add(schedule))
}

func (s *Schedule) Get(schedule *model.Schedule, columns ...string) (bool, error) {
	return s.db.Cols(columns...).Get(schedule)
}

func (s *Schedule) Update(schedule *model.Schedule, columns ...string) (int64, error) {
	return s.db.ID(schedule.Id).MustCols(columns...).Update(schedule)
}

func (s *Schedule) Delete(schedule *model.Schedule) (int64, error) {
	return s.db.Delete(schedule)
}

func (s *Schedule) add(schedule task.Schedule, task task.Task) func(session *db.Session) error {
	return func(session *db.Session) (err error) {
		if _, saved, ase := s.addSchedule(session, schedule); nil != ase {
			err = ase
		} else if _, ate := s.addTask(session, saved, task); nil != ate {
			err = ate
		}

		return
	}
}

func (s *Schedule) addSchedule(session *db.Session, schedule task.Schedule) (affected int64, saved *model.Schedule, err error) {
	saved = new(model.Schedule)
	if 0 == schedule.Id() {
		saved.Id = s.id.Next().Value()
	} else {
		saved.Id = s.id.Next().Value()
	}

	saved.Type = schedule.Type()
	saved.Subtype = schedule.Subtype()
	saved.Target = schedule.Target()
	saved.Data = schedule.Data()

	return session.Insert(saved)
}

func (s *Schedule) addTask(session *db.Session, schedule *model.Schedule, task task.Task) (affected int64, err error) {
	saved := new(model.Task)
	saved.Id = s.id.Next().Value()
	saved.Next = schedule.Next()
	saved.Status = task.StatusCreated

	now := time.Now()
	saved.Start = now
	saved.Stop = now.Add(schedule.Elapsed())

	if affected, err = session.Insert(saved); nil == err {
		task = kernel.NewTaskSchedule(schedule, saved)
	}

	return
}
