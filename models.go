type Category struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
package main

import "time"

type Task struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
