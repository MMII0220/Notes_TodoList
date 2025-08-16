package postgres

import (
	"Notes_TodoList/internal/domain"

	"github.com/jmoiron/sqlx"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) domain.CategoryRepository {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) Create(myCategory domain.Category) (domain.Category, error) {
	query := `INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	err := r.db.QueryRowx(query, myCategory.Name, myCategory.Description).
		Scan(&myCategory.ID, &myCategory.CreatedAt, &myCategory.UpdatedAt)
	return myCategory, err
}

func (r *categoryRepo) GetByID(id int) (domain.Category, error) {
	var myCategory domain.Category
	query := `SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1`
	err := r.db.Get(&myCategory, query, id)
	return myCategory, err
}

func (r *categoryRepo) Update(id int, myCategory domain.Category) (domain.Category, error) {
	query := `UPDATE categories SET name=$1, description=$2, updated_at=NOW() WHERE id=$3 RETURNING created_at, updated_at`
	err := r.db.QueryRowx(query, myCategory.Name, myCategory.Description, id).
		Scan(&myCategory.CreatedAt, &myCategory.UpdatedAt)
	myCategory.ID = id
	return myCategory, err
}

func (r *categoryRepo) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *categoryRepo) GetAll() ([]domain.Category, error) {
	rows, err := r.db.Query("SELECT id, name, description, created_at, updated_at FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var c domain.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
