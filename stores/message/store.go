package message

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"sms-provider/models"
	"sms-provider/stores"
)

type store struct{}

// New factory function to initialize store.
func New() stores.Message {
	return store{}
}

func (s store) Create(ctx *gofr.Context, message models.Message) (models.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) Get(ctx *gofr.Context) ([]models.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) GetByID(ctx *gofr.Context, id string) (models.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) Patch(ctx *gofr.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s store) Delete(ctx *gofr.Context, phNo string) error {
	//TODO implement me
	panic("implement me")
}
