package domain

import (
	"errors"
	"time"
)

type Work struct {
	StartedAt  time.Time  `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
}

func NewWork(startedAt time.Time) *Work {
	return &Work{
		StartedAt: startedAt,
	}
}

func (w *Work) Finish(finishedAt time.Time) error {
	if w.FinishedAt != nil {
		return errors.New("error.work.already_finished")
	}
	if finishedAt.Before(w.StartedAt) {
		return errors.New("error.work.finish_date_before_start_date")
	}
	w.FinishedAt = &finishedAt
	return nil
}
