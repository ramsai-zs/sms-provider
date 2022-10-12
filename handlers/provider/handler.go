package provider

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"sms-provider/stores"
)

type handler struct {
	store stores.Provider
}

func (h handler) Create(c *gofr.Context) (interface{}, error) {

}

func (h handler) GetByID(c *gofr.Context) (interface{}, error) {

}

func (h handler) Update(c *gofr.Context) (interface{}, error) {

}

func (h handler) Delete(c *gofr.Context) (interface{}, error) {

}

// New is a factory function to inject service in handler.
// nolint:revive // handler has to be unexported
func New(s stores.Provider) handler {
	return handler{store: s}
}
