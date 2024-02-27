package handler

import (
	"net/http"
	"strconv"

	"github.com/TomasCanevaro206/business-case-tarjetas-mp.git/internal/card"
	"github.com/TomasCanevaro206/business-case-tarjetas-mp.git/internal/domain"
	"github.com/TomasCanevaro206/business-case-tarjetas-mp.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type Card struct {
	cardService card.Service
}

type CardJSONBody struct {
	CardID                int    `json:"card_id"`
	CardNumber            int    `json:"card_number"`
	CardType              string `json:"card_type"`
	ExpirationDate        string `json:"expiration_date"`
	CardState             string `json:"card_state"`
	TimestampCreation     string `json:"timestamp_creation"`
	TimestampModification string `json:"timestamp_modification"`
}

func NewCard(e card.Service) *Card {
	return &Card{
		cardService: e,
	}
}

// showCards godoc
//
//	@Summary		Show a card
//	@Description	Get card by ID
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Card ID"
//	@Success		200	{object}	web.Response{data=domain.Card}
//	@Failure		404	{object}	web.ErrorResponse{code=string,message=string}
//	@Failure		422	{object}	web.ErrorResponse{code=string,message=string}
//	@Router			/card/{id} [get]
func (e *Card) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.ResponseMessage(c, http.StatusBadRequest, map[string]interface{}{
				"message": "invalid id",
			})
			return
		}

		foundCard, err := e.cardService.Get(c, id)
		if err != nil {
			if err == card.ErrNotFound {
				web.ResponseMessage(c, http.StatusNotFound, map[string]interface{}{
					"message": "no card with matching id found",
				})
			} else {
				web.ResponseMessage(c, http.StatusInternalServerError, map[string]interface{}{
					"message": "unexpected server error",
				})
			}

			return
		}

		data := map[string]interface{}{
			"data": foundCard,
		}

		web.ResponseMessage(c, http.StatusOK, data)
	}
}

// listCards godoc
//
//	@Summary		Show all cards
//	@Description	Show all available cards
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.Response{data=[]domain.Card}
//	@Failure		503	{object}	web.ErrorResponse{code=string,message=string}
//	@Router			/card [get]
func (e *Card) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		cards, err := e.cardService.GetAll(c)

		if err != nil {
			if err == card.ErrNotFound {
				web.ResponseMessage(c, http.StatusNotFound, map[string]interface{}{
					"message": "no cards where found",
				})
			} else {
				web.ResponseMessage(c, http.StatusInternalServerError, map[string]interface{}{
					"message": "unexpected server error",
				})
			}
			return
		}

		data := map[string]interface{}{
			"data": cards,
		}
		web.ResponseMessage(c, http.StatusOK, data)
	}
}

// newCard godoc
//
//	@Summary		Create an card
//	@Description	Create a new card with unique ID
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Param			card	body		domain.Card	true	"Card Details"
//	@Success		201		{object}	web.Response{data=domain.Card}
//	@Failure		409		{object}	web.ErrorResponse{code=string,message=string}
//	@Failure		422		{object}	web.ErrorResponse{code=string,message=string}
//	@Router			/card [post]
func (e *Card) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request CardJSONBody
		err := c.ShouldBindJSON(&request)
		if err != nil {
			web.ResponseMessage(c, http.StatusBadRequest, map[string]interface{}{
				"message": "invalid request body parameters",
			})
			return
		}

		newCard := domain.Card{
			CardID:                request.CardID,
			CardNumber:            request.CardNumber,
			CardType:              request.CardType,
			ExpirationDate:        request.ExpirationDate,
			CardState:             request.CardState,
			TimestampCreation:     request.TimestampCreation,
			TimestampModification: request.TimestampModification,
		}

		id, err := e.cardService.Create(c, newCard)
		if err != nil {
			switch err {
			case card.ErrExists:
				web.ResponseMessage(c, http.StatusConflict, map[string]interface{}{
					"message": "card already exists",
				})
			case card.ErrInvalidBody:
				web.ResponseMessage(c, http.StatusUnprocessableEntity, map[string]interface{}{
					"message": "invalid request body parameters",
				})
			default:
				web.ResponseMessage(c, http.StatusInternalServerError, map[string]interface{}{
					"message": "unexpected server error",
				})

			}
			return
		}
		newCard.CardID = id

		data := map[string]interface{}{
			"data": newCard,
		}
		web.ResponseMessage(c, http.StatusCreated, data)
	}
}

// updateCard godoc
//
//	@Summary		Update an card
//	@Description	Update an card's fields by ID
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Card ID"
//	@Param			card	body		domain.Card	false	"Card Details"
//	@Success		200	{object}	web.Response{data=domain.Card}
//	@Failure		404	{object}	web.ErrorResponse{code=string,message=string}
//	@Failure		409	{object}	web.ErrorResponse{code=string,message=string}
//	@Failure		422	{object}	web.ErrorResponse{code=string,message=string}
//	@Failure		503	{object}	web.ErrorResponse{code=string,message=string}
//	@Router			/card/{id} [patch]
func (e *Card) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "invalid id")
			return
		}

		var request CardJSONBody
		err = c.ShouldBindJSON(&request)
		if err != nil {
			web.ResponseMessage(c, http.StatusBadRequest, map[string]interface{}{
				"message": "invalid request body parameters",
			})
			return
		}

		newCard := domain.Card{
			CardID:                request.CardID,
			CardNumber:            request.CardNumber,
			CardType:              request.CardType,
			ExpirationDate:        request.ExpirationDate,
			CardState:             request.CardState,
			TimestampCreation:     request.TimestampCreation,
			TimestampModification: request.TimestampModification,
		}

		updatedCard, err := e.cardService.Update(c, newCard, id)

		if err != nil {
			switch err {
			case card.ErrNotFound:
				web.Error(c, http.StatusNotFound, "card not found")
			case card.ErrInvalidBody:
				web.Error(c, http.StatusBadRequest, "invalid body parameters")
			default:
				web.Error(c, http.StatusInternalServerError, "unexpected server error")
			}
			return
		}

		data := map[string]interface{}{
			"data": updatedCard,
		}
		web.ResponseMessage(c, http.StatusOK, data)
	}
}

// deleteCard godoc
//
//	@Summary		Delete an card
//	@Description	Delete an card by ID
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Card ID"
//	@Success		204
//	@Failure		404	{object}	web.ErrorResponse{code=string,message=string}
//	@Failure		422	{object}	web.ErrorResponse{code=string,message=string}
//	@Failure		503	{object}	web.ErrorResponse{code=string,message=string}
//	@Router			/card/{id} [delete]
func (e *Card) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "invalid id")
			return
		}

		err = e.cardService.Delete(c, id)
		if err != nil {
			if err == card.ErrNotFound {
				web.ResponseMessage(c, http.StatusNotFound, map[string]interface{}{
					"message": "no card with matching id found",
				})
			} else {
				web.ResponseMessage(c, http.StatusInternalServerError, map[string]interface{}{
					"message": "unexpected server error",
				})
			}
			return
		}

		web.ResponseMessage(c, http.StatusNoContent, map[string]interface{}{
			"message": "card deleted successfully",
		})
	}
}
