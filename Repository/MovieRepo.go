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
	sqlRequest := `SELECT
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
	rows, err := r.db.Query(c, sqlRequest)
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
		Movies[i].Category = bufMovies[Movies[i].ID].Category
	}
	return Movies, nil
}

// WARN: If this function couldn't find any movies with this id, he will SEND MOVIE WITH ID = -1!
func (r *MovieRepository) FindThisMovie(c context.Context, id int) (Structs.Movie, error) {
	sqlRequest := `SELECT
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
	rows, err := r.db.Query(c, sqlRequest, id)
	if err != nil {
		return Structs.Movie{ID: -1}, err
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
			return Structs.Movie{ID: -1}, err
		}
		Movie.Category = append(Movie.Category, cat)
	}
	if isThisANullComplexStruct(Movie) {
		return Structs.Movie{ID: -1}, nil
	}
	return Movie, nil
}

func (r *MovieRepository) CreateMovie(c context.Context, movie Structs.Movie) (int, error) {
	var id int
	var categories []Structs.Category = movie.Category
	sqlRequest := `INSERT INTO movies
		(movie_title, director, producer, description, realesed)
		VALUES($1, $2, $3, $4, $5)
		returning id`

	rows := r.db.QueryRow(c, sqlRequest,
		movie.MovieTitle,
		movie.Director,
		movie.Producer,
		movie.Description,
		movie.Realesed,
	)
	err := rows.Scan(&id)
	if err != nil {
		return -1, err
	}
	for _, v := range categories {
		_, err = r.db.Exec(c, `INSERT INTO public.category_movie (category_ids, movie_ids) VALUES($1, $2)`, v.ID, id)
		if err != nil {
			return -1, err
		}
	}
	return id, nil
}

func (r *MovieRepository) UpdateMovie(c context.Context, movie Structs.Movie) error {
	var categories []Structs.Category = movie.Category
	sqlRequest := `UPDATE movies
		SET movie_title=$2, director=$3, producer=$4, description=$5, realesed=$6
		WHERE id = $1`
	_, err := r.db.Exec(c, sqlRequest, movie.ID, movie.MovieTitle, movie.Director, movie.Producer, movie.Description, movie.Realesed)
	if err != nil {
		return err
	}
	sqlRequest = `DELETE FROM category_movie WHERE movie_ids = $1`
	_, err = r.db.Exec(c, sqlRequest, movie.ID)
	if err != nil {
		return err
	}
	sqlRequest = `INSERT INTO category_movie (category_ids, movie_ids) VALUES($1, $2)`
	for _, v := range categories {
		_, err = r.db.Exec(c, sqlRequest, v.ID, movie.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *MovieRepository) DeleteMovie(c context.Context, id int) error {
	sqlRequest := `DELETE FROM category_movie WHERE movie_ids = $1`
	_, err := r.db.Exec(c, sqlRequest, id)
	if err != nil {
		return err
	}
	sqlRequest = `DELETE FROM movies WHERE id = $1`
	_, err = r.db.Exec(c, sqlRequest, id)
	if err != nil {
		return err
	}
	return nil
}

func isThisANullComplexStruct(Movie Structs.Movie) bool {
	if (Movie.ID == 0) && (Movie.MovieTitle == "") && (Movie.Director == "") && (Movie.Producer == "") && (Movie.Description == "") && (Movie.Realesed == 0) && (Movie.Category == nil) && (Movie.Cards == nil) {
		return true
	}
	return false
}
