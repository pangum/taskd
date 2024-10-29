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
		tx: database.Tx,
	}
}

func (t *Schedule) Add(schedule task.Schedule) error {
	return t.tx.Do(t.add(schedule))
}

func (t *Schedule) Get(schedule *model.Schedule, columns ...string) (bool, error) {
	return t.db.Cols(columns...).Get(schedule)
}

func (t *Schedule) Update(schedule *model.Schedule, columns ...string) (int64, error) {
	return t.db.ID(schedule.Id).MustCols(columns...).Update(schedule)
}

func (t *Schedule) Delete(schedule *model.Schedule) (int64, error) {
	return t.db.Delete(schedule)
}

func (t *Schedule) add(schedule task.Schedule) func(session *db.Session) error {
	return func(session *db.Session) (err error) {
		if _, ase := t.addSchedule(session, schedule); nil != ase {
			err = ase
		} else if _, ate := t.addTask(session, schedule); nil != ate {
			err = ate
		}

		return
	}
}

func (t *Schedule) addSchedule(session *db.Session, schedule task.Schedule) (int64, error) {
	inserted := new(model.Schedule)
	if 0 == schedule.Id() {
		inserted.Id = t.id.Next().Value()
	} else {
		inserted.Id = t.id.Next().Value()
	}

	inserted.Type = schedule.Type()
	inserted.Subtype = schedule.Subtype()
	inserted.Target = schedule.Target()
	inserted.Data = schedule.Data()

	return session.Insert(inserted)
}

func (t *Schedule) addTask(session *db.Session, schedule task.Schedule) (int64, error) {
	inserted := new(model.Task)
	inserted.Id = t.id.Next().Value()
	inserted.Next = schedule.Next()
	inserted.Status = task.StatusCreated

	now := time.Now()
	inserted.Start = now
	inserted.Stop = now.Add(schedule.Elapsed())

	return session.Insert(inserted)
}
