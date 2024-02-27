package card

import (
	"context"
	"database/sql"
	"errors"

	"github.com/TomasCanevaro206/business-case-tarjetas-mp.git/internal/domain"
)

// Errors
var (
	RepoErrNoRows   = errors.New("there are no rows in the database that match the given data")
	RepoErrInternal = errors.New("there was an unexpected error while trying to fetch data")
)

// Repository encapsulates the storage of a Card.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Card, error)
	Get(ctx context.Context, id int) (domain.Card, error)
	Exists(ctx context.Context, cardId int) bool
	Save(ctx context.Context, c domain.Card) (int, error)
	Update(ctx context.Context, c domain.Card) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Card, error) {
	query := "SELECT * FROM cards;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var cards []domain.Card

	for rows.Next() {
		c := domain.Card{}
		_ = rows.Scan(&c.CardID, &c.CardNumber, &c.CardType, &c.ExpirationDate, &c.CardState, &c.TimestampCreation, &c.TimestampModification)
		cards = append(cards, c)
	}

	return cards, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Card, error) {
	query := "SELECT * FROM cards WHERE card_id=?;"
	row := r.db.QueryRow(query, id)
	c := domain.Card{}
	err := row.Scan(&c.CardID, &c.CardNumber, &c.CardType, &c.ExpirationDate, &c.CardState, &c.TimestampCreation, &c.TimestampModification)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Card{}, RepoErrNoRows
		}
		return domain.Card{}, RepoErrInternal
	}

	return c, nil
}

func (r *repository) Exists(ctx context.Context, cardId int) bool {
	query := "SELECT card_id FROM cards WHERE card_id=?;"
	row := r.db.QueryRow(query, cardId)
	err := row.Scan(&cardId)
	return err == nil
}

func (r *repository) Save(ctx context.Context, c domain.Card) (int, error) {
	query := "INSERT INTO cards(card_id,card_number,card_type,expiration_date,card_state,timestamp_creation,timestamp_modificaction) VALUES (?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(c.CardID, c.CardNumber, c.CardType, c.ExpirationDate, c.CardState, c.TimestampCreation, c.TimestampModification)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, c domain.Card) error {
	query := "UPDATE cards SET card_id=?, card_number=?, card_type=?, expiration_date=?, card_state=?, timestamp_creation=?, timestamp_modificaction=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(c.CardID, c.CardNumber, c.CardType, c.ExpirationDate, c.CardState, c.TimestampCreation, c.TimestampModification)
	if err != nil {
		return err
	}

	n_affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n_affected < 1 {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM cards WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return RepoErrNoRows
	}

	return nil
}
