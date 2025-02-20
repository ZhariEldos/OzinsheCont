package Repository

import (
	"context"
	"ozinsheproject/Structs"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(conn *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{db: conn}
}

func (r *CategoryRepository) FindCategoryByID(c context.Context, id int) (Structs.Category, error) {
	var category Structs.Category
	sqlRequst := `SELECT id, category_title FROM category WHERE id = $1`
	rows := r.db.QueryRow(c, sqlRequst, id)
	err := rows.Scan(&category.ID, &category.CategoryTitle)
	if err != nil {
		return Structs.Category{}, err
	}
	return category, nil
}

func (r *CategoryRepository) FindAllCategories(c context.Context) ([]Structs.Category, error) {
	var categories []Structs.Category
	sqlRequst := `SELECT id, category_title FROM category`
	rows, err := r.db.Query(c, sqlRequst)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c Structs.Category
		err = rows.Scan(&c.ID, &c.CategoryTitle)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}
