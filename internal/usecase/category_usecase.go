package usecase

import (
	"Notes_TodoList/internal/domain"
)

type categoryUsecase struct {
	repo domain.CategoryRepository
}

func NewCategoryUsecase(repo domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUsecase{repo: repo}
}

func (uc *categoryUsecase) Create(myCategory domain.Category) (domain.Category, error) {
	return uc.repo.Create(myCategory)
}

func (uc *categoryUsecase) GetByID(id int) (domain.Category, error) {
	return uc.repo.GetByID(id)
}

func (uc *categoryUsecase) GetAll() ([]domain.Category, error) {
	return uc.repo.GetAll()
}

func (uc *categoryUsecase) Update(id int, myCategory domain.Category) (domain.Category, error) {
	return uc.repo.Update(id, myCategory)
}

func (uc *categoryUsecase) Delete(id int) error {
	return uc.repo.Delete(id)
}
