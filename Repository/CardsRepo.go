package Repository

import (
	"context"
	"ozinsheproject/Structs"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CardsRepository struct {
	db *pgxpool.Pool
}

func NewCardsRepository(conn *pgxpool.Pool) *CardsRepository {
	return &CardsRepository{conn}
}

// TODO: Correct to CRUD

func (r *CardsRepository) FindAllCards(c context.Context) ([]Structs.Cards, error) {
	sqlRequest := `SELECT id, card_title, url_picture FROM cards`
	rows, err := r.db.Query(c, sqlRequest)
	if err != nil {
		return nil, err
	}
	var cards []Structs.Cards
	for rows.Next() {
		var card Structs.Cards
		err = rows.Scan(&card.ID, &card.CardsTitle, &card.URLPicture)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (r *CardsRepository) FindThisCard(c context.Context, id int) (Structs.Cards, error) {
	var card Structs.Cards

	sqlRequest := `SELECT id, card_title, url_picture FROM cards WHERE id = $1`
	rows := r.db.QueryRow(c, sqlRequest, id)
	err := rows.Scan(&card.ID, &card.CardsTitle, &card.URLPicture)
	if err != nil {
		return Structs.Cards{}, err
	}

	return card, nil
}

func (r *CardsRepository) CreateCard(c context.Context, card Structs.Cards) (int, error) {
	sqlRequest := `INSERT INTO public.cards
		(card_title, url_picture)
		VALUES($1, $2)
		RETURNING id`
	rows := r.db.QueryRow(c, sqlRequest, card.CardsTitle, card.URLPicture)
	var id int
	err := rows.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CardsRepository) UpdateCard(c context.Context, card Structs.Cards) error {
	sqlRequest := `UPDATE public.cards SET card_title=$2, url_picture=$3 WHERE id = $1`
	_, err := r.db.Exec(c, sqlRequest, card.ID, card.CardsTitle, card.URLPicture)
	if err != nil {
		return err
	}
	return nil
}

func (r *CardsRepository) DeleteCards(c context.Context, id int) error {
	sqlRequest := `DELETE FROM public.cards WHERE id = $1`
	_, err := r.db.Exec(c, sqlRequest, id)
	if err != nil {
		return err
	}
	return nil
}
