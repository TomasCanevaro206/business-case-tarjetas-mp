package card

import (
	"context"
	"errors"

	"github.com/TomasCanevaro206/business-case-tarjetas-mp.git/internal/domain"
)

// Errors
var (
	ErrNotFound    = errors.New("card not found")
	ErrExists      = errors.New("card already exists")
	ErrInvalidBody = errors.New("invalid body data")
	ErrUnexpected  = errors.New("unexpected server error")
)

type Service interface {
	GetAll(c context.Context) ([]domain.Card, error)
	Get(c context.Context, id int) (domain.Card, error)
	Create(c context.Context, newCard domain.Card) (int, error)
	Update(c context.Context, card domain.Card, id int) (domain.Card, error)
	Delete(c context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetAll(c context.Context) ([]domain.Card, error) {
	return s.repository.GetAll(c)
}

func (s *service) Get(c context.Context, id int) (domain.Card, error) {
	result, err := s.repository.Get(c, id)
	if err != nil {
		switch err {
		case RepoErrNoRows:
			return domain.Card{}, ErrNotFound
		default:
			return domain.Card{}, ErrUnexpected
		}
	}
	return result, nil
}

func (s *service) Create(c context.Context, newCard domain.Card) (int, error) {
	if s.repository.Exists(c, newCard.CardID) {
		return 0, ErrExists
	}

	id, err := s.repository.Save(c, newCard)
	if err != nil {
		return 0, err
	}

	newCard.CardID = id

	return id, nil
}

func (s *service) Update(c context.Context, newCard domain.Card, id int) (updatedCard domain.Card, err error) {
	updatedCard, err = s.Get(c, id)
	if err != nil {
		if err == ErrNotFound {
			return domain.Card{}, ErrNotFound
		} else {
			return domain.Card{}, ErrUnexpected
		}
	}

	if newCard.CardID <= 0 && newCard.CardNumber <= 0 && newCard.CardType == "" && newCard.ExpirationDate == "" && newCard.CardState == "" && newCard.TimestampCreation == "" && newCard.TimestampModification == "" {
		return domain.Card{}, ErrInvalidBody
	}

	updatedCardPtr := &updatedCard
	if newCard.CardID > 0 {
		updatedCardPtr.CardID = newCard.CardID
	}
	if newCard.CardNumber > 0 {
		updatedCardPtr.CardNumber = newCard.CardNumber
	}
	if newCard.CardType != "" {
		updatedCardPtr.CardType = newCard.CardType
	}
	if newCard.ExpirationDate != "" {
		updatedCardPtr.ExpirationDate = newCard.ExpirationDate

	}
	if newCard.CardState != "" {
		updatedCardPtr.CardState = newCard.CardState
	}
	if newCard.TimestampCreation != "" {
		updatedCardPtr.TimestampCreation = newCard.TimestampCreation
	}
	if newCard.TimestampModification != "" {
		updatedCardPtr.TimestampModification = newCard.TimestampModification
	}

	err = s.repository.Update(c, updatedCard)
	if err != nil {
		switch err {
		case RepoErrNoRows:
			return domain.Card{}, ErrNotFound
		default:
			return domain.Card{}, ErrUnexpected
		}
	}

	return
}

func (s *service) Delete(c context.Context, id int) error {
	err := s.repository.Delete(c, id)

	if err != nil {
		switch err {
		case RepoErrNoRows:
			return ErrNotFound
		default:
			return ErrUnexpected
		}
	}
	return nil
}
