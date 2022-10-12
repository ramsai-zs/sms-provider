package models

import (
	"github.com/google/uuid"

	"sms-provider/types"
)

type Provider struct {
	ID           uuid.UUID `json:"id"`
	URL          string    `json:"URL"`
	ChannelRefID string    `json:"channelRefID"`
	Name         string    `json:"name"`
}

type Message struct {
	ID            uuid.UUID    `json:"id"`
	Message       string       `json:"message"`
	Number        string       `json:"number"`
	Transactional string       `json:"transactional"`
	ProviderRefID uuid.UUID    `json:"ProviderRefID"`
	Status        types.Status `json:"status"`
	DeliveredTime string       `json:"deliveredTime"`
}