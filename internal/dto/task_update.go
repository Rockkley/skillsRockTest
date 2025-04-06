package dto

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type TaskStatus string

const (
	StatusNew        TaskStatus = "new"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "done"
)

type TaskUpdateDTO struct {
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status,omitempty"`
}

func (t *TaskUpdateDTO) Validate() error {
	if title := strings.TrimSpace(t.Title); title != "" {
		if l := utf8.RuneCountInString(title); l < 3 || l > 100 {

			return errors.New("title must be between 3 and 100 characters")
		}
	}

	if desc := strings.TrimSpace(t.Description); desc != "" {
		if utf8.RuneCountInString(desc) < 3 {

			return errors.New("description must be at least 3 characters")
		}
	}

	if t.Status != "" && !IsValidTaskStatus(t.Status) {

		return errors.New("status must be one of: new, in_progress, done")
	}

	return nil
}

func IsValidTaskStatus(status TaskStatus) bool {
	switch status {
	case StatusNew, StatusInProgress, StatusDone:
		return true
	default:
		return false
	}
}
