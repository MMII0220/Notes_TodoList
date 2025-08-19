package models

import "time"

type Note struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name" binding:"required"`
	Description string    `db:"description" json:"description"`
	CategoryID  *int      `db:"category_id" json:"category_id"`
	Priority    int       `db:"priority" json:"priority"`
	IsDone      bool      `db:"is_done" json:"is_done"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
