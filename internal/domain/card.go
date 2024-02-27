package domain

type Card struct {
	CardID                int    `json:"card_id"`
	CardNumber            int    `json:"card_number"`
	CardType              string `json:"card_type"`
	ExpirationDate        string `json:"expiration_date"`
	CardState             string `json:"card_state"`
	TimestampCreation     string `json:"timestamp_creation"`
	TimestampModification string `json:"timestamp_modification"`
}
