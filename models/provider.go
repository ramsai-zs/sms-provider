package models

import (
	"github.com/google/uuid"
)

type Provider struct {
	ID           uuid.UUID `json:"id"`
	URL          string    `json:"URL"`
	ChannelRefID uuid.UUID `json:"channelRefID"`
	Name         string    `json:"name"`
}
