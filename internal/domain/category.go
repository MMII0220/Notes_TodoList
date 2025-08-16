package domain

import "time"

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryUsecase interface {
	Create(myCategory Category) (Category, error)
	GetByID(id int) (Category, error)
	GetAll() ([]Category, error)
	Update(id int, myCategory Category) (Category, error)
	Delete(id int) error
}

type CategoryRepository interface {
	Create(myCategory Category) (Category, error)
	GetByID(id int) (Category, error)
	GetAll() ([]Category, error)
	Update(id int, myCategory Category) (Category, error)
	Delete(id int) error
}
