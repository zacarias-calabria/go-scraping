package domain

import (
	"testing"
	"time"
)

func TestWork_NewWork(t *testing.T) {
	t.Run("Should create a new work successfully", func(t *testing.T) {
		now := time.Now()
		work := NewWork(now)
		if work.StartedAt != now {
			t.Errorf("Expected StartedAt to be %v, got %v", now, work.StartedAt)
		}
		if work.FinishedAt != nil {
			t.Error("Expected FinishedAt to be nil for new work")
		}
	})
}

func TestWork_Finish(t *testing.T) {
	t.Run("should finish work successfully", func(t *testing.T) {
		startTime := time.Now()
		finishTime := startTime.Add(1 * time.Hour)
		work := NewWork(startTime)
		err := work.Finish(finishTime)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if work.FinishedAt == nil {
			t.Error("Expected FinishedAt to be set")
		}
		if *work.FinishedAt != finishTime {
			t.Errorf("Expected FinishedAt to be %v, got %v", finishTime, *work.FinishedAt)
		}
	})

	t.Run("should not finish work that is already finished", func(t *testing.T) {
		// Arrange
		startTime := time.Now()
		firstFinishTime := startTime.Add(1 * time.Hour)
		secondFinishTime := startTime.Add(2 * time.Hour)
		work := NewWork(startTime)
		work.Finish(firstFinishTime)
		err := work.Finish(secondFinishTime)
		if err == nil {
			t.Error("Expected error for finishing already finished work")
		}
		if err.Error() != "error.work.already_finished" {
			t.Errorf("Expected error message 'error.work_already_finished', got '%v'", err.Error())
		}
		if *work.FinishedAt != firstFinishTime {
			t.Errorf("Expected FinishedAt to remain %v, got %v", firstFinishTime, *work.FinishedAt)
		}
	})

	t.Run("should not finish work with finish time before start time", func(t *testing.T) {
		// Arrange
		startTime := time.Now()
		finishTime := startTime.Add(-1 * time.Hour)
		work := NewWork(startTime)
		err := work.Finish(finishTime)
		if err == nil {
			t.Error("Expected error for finish time before start time")
		}
		if err.Error() != "error.work.finish_date_before_start_date" {
			t.Errorf("Expected error message 'error.work.finish_date_before_start_date', got '%v'", err.Error())
		}
		if work.FinishedAt != nil {
			t.Error("Expected FinishedAt to remain nil")
		}
	})
}
