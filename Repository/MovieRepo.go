package Repository

import (
	"context"
	"ozinsheproject/Structs"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MovieRepository struct {
	db *pgxpool.Pool
}

func NewMovieRepository(conn *pgxpool.Pool) *MovieRepository {
	return &MovieRepository{db: conn}
}

func (r *MovieRepository) FindAllMovies(c context.Context) ([]Structs.Movie, error) {
	sqlRequst := `SELECT
	m.id, 
	m.movie_title,
	m.director,
	m.producer,
	m.description,
	m.realesed,
	c.id,
	c.category_title
	FROM movies m 
	JOIN category_movie mc ON mc.movie_ids = m.id 
	JOIN category c ON mc.category_ids = c.id`
	rows, err := r.db.Query(c, sqlRequst)
	if err != nil {
		return nil, err
	}
	var Movies []Structs.Movie
	bufMovies := make(map[int]*Structs.Movie)
	for rows.Next() {
		var mov Structs.Movie
		var cat Structs.Category
		err = rows.Scan(
			&mov.ID,
			&mov.MovieTitle,
			&mov.Director,
			&mov.Producer,
			&mov.Description,
			&mov.Realesed,
			&cat.ID,
			&cat.CategoryTitle)
		if err != nil {
			return nil, err
		}
		if _, exist := bufMovies[mov.ID]; !exist {
			bufMovies[mov.ID] = &mov
			Movies = append(Movies, mov)
		}
		bufMovies[mov.ID].Category = append(bufMovies[mov.ID].Category, cat)
	}
	for i := range Movies {
		Movies[i].Category = bufMovies[i+1].Category
	}
	return Movies, nil
}

func (r *MovieRepository) FindThisMovie(c context.Context, id int) (Structs.Movie, error) {
	sqlRequst := `SELECT
	m.id, 
	m.movie_title,
	m.director,
	m.producer,
	m.description,
	m.realesed,
	c.id,
	c.category_title
	FROM movies m 
	JOIN category_movie mc ON mc.movie_ids = m.id 
	JOIN category c ON mc.category_ids = c.id
	WHERE m.id = $1`
	rows, err := r.db.Query(c, sqlRequst, id)
	if err != nil {
		return Structs.Movie{}, err
	}
	var Movie Structs.Movie
	for rows.Next() {
		var cat Structs.Category
		err = rows.Scan(
			&Movie.ID,
			&Movie.MovieTitle,
			&Movie.Director,
			&Movie.Producer,
			&Movie.Description,
			&Movie.Realesed,
			&cat.ID,
			&cat.CategoryTitle)
		if err != nil {
			return Structs.Movie{}, err
		}
		Movie.Category = append(Movie.Category, cat)
	}
	return Movie, nil
}

func (r *MovieRepository) CreateMovie(c context.Context, movie Structs.Movie) (int, error) {
	var id int
	var categories []Structs.Category = movie.Category
	sqlRequst := `INSERT INTO movies
		(movie_title, director, producer, description, realesed)
		VALUES($1, $2, $3, $4, $5)
		returning id`

	rows := r.db.QueryRow(c, sqlRequst,
		movie.MovieTitle,
		movie.Director,
		movie.Producer,
		movie.Description,
		movie.Realesed,
	)
	err := rows.Scan(&id)
	if err != nil {
		return 0, err
	}
	for _, v := range categories {
		_, err = r.db.Exec(c, `INSERT INTO public.category_movie (category_ids, movie_ids) VALUES($1, $2)`, v.ID, id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (r *MovieRepository) UpdateMovie(c context.Context, movie Structs.Movie) error {
	var categories []Structs.Category = movie.Category
	sqlRequst := `UPDATE movies
		SET movie_title=$2, director=$3, producer=$4, description=$5, realesed=$6
		WHERE id = $1`
	_, err := r.db.Exec(c, sqlRequst, movie.ID, movie.MovieTitle, movie.Director, movie.Producer, movie.Description, movie.Realesed)
	if err != nil {
		return err
	}
	sqlRequst = `DELETE FROM category_movie WHERE movie_ids = $1`
	_, err = r.db.Exec(c, sqlRequst, movie.ID)
	if err != nil {
		return err
	}
	sqlRequst = `INSERT INTO category_movie (category_ids, movie_ids) VALUES($1, $2)`
	for _, v := range categories {
		_, err = r.db.Exec(c, sqlRequst, v.ID, movie.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *MovieRepository) DeleteMovie(c context.Context, id int) error {
	sqlRequst := `DELETE FROM category_movie WHERE movie_ids = $1`
	_, err := r.db.Exec(c, sqlRequst, id)
	if err != nil {
		return err
	}
	sqlRequst = `DELETE FROM movies WHERE id = $1`
	_, err = r.db.Exec(c, sqlRequst, id)
	if err != nil {
		return err
	}
	return nil
}
