package repository

import (
	"Notes_TodoList/internal/domain"
	"context"
)

type CategoryRepository interface {
	Create(ctx context.Context, myCategory *domain.Category) error
	Update(ctx context.Context, myCategory *domain.Category) error
	GetByID(ctx context.Context, id int) (*domain.Category, error)
	Delete(ctx context.Context, id int) error
	GetAll() ([]domain.Category, error)
}
