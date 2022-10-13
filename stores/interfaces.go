package stores

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"sms-provider/models"
)

type Message interface {
	Create(ctx *gofr.Context, message models.Message) (models.Message, error)
	Get(ctx *gofr.Context) ([]models.Message, error)
	GetByID(ctx *gofr.Context, id string) (models.Message, error)
	Patch(ctx *gofr.Context, id string) error
	Delete(ctx *gofr.Context, phNo string) error
}

type Provider interface {
	Create(ctx *gofr.Context, provider models.Provider) (models.Provider, error)
	GetByID(ctx *gofr.Context, id string) (models.Provider, error)
	Patch(ctx *gofr.Context, provider models.Provider) error
	Delete(ctx *gofr.Context, id string) error
}
