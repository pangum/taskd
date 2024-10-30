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

func (s *Schedule) Add(schedule *model.Schedule, next time.Time) (runnable *model.Task, err error) {
	runnable = new(model.Task)
	err = s.tx.Do(s.add(schedule, runnable, next))

	return
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

func (s *Schedule) add(schedule *model.Schedule, runnable *model.Task, next time.Time) func(session *db.Session) error {
	return func(session *db.Session) (err error) {
		if 0 == schedule.Id {
			schedule.Id = s.id.Next().Value()
		}

		if _, ie := session.Insert(schedule); nil != ie {
			err = ie
		} else {
			runnable.Id = s.id.Next().Value()
			runnable.Next = next
			runnable.Status = task.StatusCreated

			now := time.Now()
			runnable.Start = now
			runnable.Stop = now.Add(schedule.Elapsed)

			_, err = session.Insert(runnable)
		}

		return
	}
}
