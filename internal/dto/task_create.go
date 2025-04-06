package dto

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type TaskCreateDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t *TaskCreateDTO) Validate() error {
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

	return nil
}
