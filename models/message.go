package models

import (
	"github.com/google/uuid"

	"sms-provider/types"
)

type Message struct {
	ID            uuid.UUID    `json:"id"`
	Message       string       `json:"message"`
	Number        string       `json:"number"`
	Transactional string       `json:"transactional"`
	ProviderRefID uuid.UUID    `json:"ProviderRefID"`
	Status        types.Status `json:"status"`
	DeliveredTime int64        `json:"deliveredTime"`
}
