package provider

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"sms-provider/models"
	"sms-provider/stores"
)

type store struct{}

// New factory function to initialize store.
func New() stores.Provider {
	return store{}
}

func (s store) Create(ctx *gofr.Context, message models.Provider) (models.Provider, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) GetByID(ctx *gofr.Context, id string) (models.Provider, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) PUT(ctx *gofr.Context, id string) (models.Provider, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) Delete(ctx *gofr.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
