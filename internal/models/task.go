package models

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrTitleRequired      = errors.New("title is required")
	ErrTitleTooLong       = errors.New("title must be less than 200 characters")
	ErrDescriptionTooLong = errors.New("description must be less than 1000 characters")
	ErrTaskNotFound       = errors.New("task not found")
	ErrInvalidID          = errors.New("invalid task ID")
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" example:"Buy groceries"`
	Description string    `json:"description" example:"Milk, eggs, bread"`
	Completed   bool      `json:"completed" example:"false"`
	CreatedAt   time.Time `json:"created_at" example:"2024-01-01T10:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-01T10:00:00Z"`
}

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required" example:"Buy groceries"`
	Description string `json:"description" example:"Milk, eggs, bread"`
}

type UpdateTaskRequest struct {
	Title       *string `json:"title,omitempty" example:"Buy groceries"`
	Description *string `json:"description,omitempty" example:"Milk, eggs, bread"`
	Completed   *bool   `json:"completed,omitempty" example:"true"`
}

func (r *CreateTaskRequest) Validate() error {

	r.Title = strings.TrimSpace(r.Title)
	r.Description = strings.TrimSpace(r.Description)

	if r.Title == "" {
		return ErrTitleRequired
	}

	if len(r.Title) > 200 {
		return ErrTitleTooLong
	}

	if len(r.Description) > 1000 {
		return ErrDescriptionTooLong
	}

	return nil
}

func (r *UpdateTaskRequest) Validate() error {
	// Validate title if provided
	if r.Title != nil {
		trimmed := strings.TrimSpace(*r.Title)
		if trimmed == "" {
			return ErrTitleRequired
		}
		if len(trimmed) > 200 {
			return ErrTitleTooLong
		}
		*r.Title = trimmed
	}

	// Validate description if provided
	if r.Description != nil {
		trimmed := strings.TrimSpace(*r.Description)
		if len(trimmed) > 1000 {
			return ErrDescriptionTooLong
		}
		*r.Description = trimmed
	}

	return nil
}
