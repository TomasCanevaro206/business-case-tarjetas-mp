package card

import (
	"github.com/TomasCanevaro206/business-case-tarjetas-mp.git/internal/domain"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	employees []domain.Card
	Mock      mock.Mock
}
