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

func (r *CategoryRepository) FindThisCategory(c context.Context, id int) (Structs.Category, error) {
	var category Structs.Category
	sqlRequest := `SELECT id, category_title FROM category WHERE id = $1`
	rows := r.db.QueryRow(c, sqlRequest, id)
	err := rows.Scan(&category.ID, &category.CategoryTitle)
	if err != nil {
		return Structs.Category{}, err
	}
	return category, nil
}

func (r *CategoryRepository) FindAllCategories(c context.Context) ([]Structs.Category, error) {
	var categories []Structs.Category
	sqlRequest := `SELECT id, category_title FROM category`
	rows, err := r.db.Query(c, sqlRequest)
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

func (r *CategoryRepository) CreateCategory(c context.Context, category Structs.Category) (int, error) {
	var id int
	sqlRequest := `INSERT INTO category
	(category_title)
	VALUES($1)
	returning id`
	rows := r.db.QueryRow(c, sqlRequest, category.CategoryTitle)
	err := rows.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *CategoryRepository) UpdateCategory(c context.Context, category Structs.Category) error {
	sqlRequest := `UPDATE category
	SET category_title=$2
	WHERE id=$1`
	_, err := r.db.Exec(c, sqlRequest, category.ID, category.CategoryTitle)
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) DeleteCategory(c context.Context, id int) error {
	sqlRequest := `DELETE FROM category_movie WHERE category_ids = $1`
	_, err := r.db.Exec(c, sqlRequest, id)
	if err != nil {
		return err
	}
	sqlRequest = `DELETE FROM category WHERE id = $1`
	_, err = r.db.Exec(c, sqlRequest, id)
	if err != nil {
		return err
	}
	return nil
}
