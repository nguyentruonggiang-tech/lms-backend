package dto

import (
	"lms-api/ent"
	"time"
)

type LessonWithProgress struct {
	Lesson      *ent.Lessons `json:"lesson"`
	IsCompleted bool         `json:"is_completed"`
	CompletedAt *time.Time   `json:"completed_at,omitempty"`
}
