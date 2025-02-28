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
