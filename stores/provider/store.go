package provider

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
	"sms-provider/models"
	"sms-provider/stores"
)

type store struct{}

// New factory function to initialize store.
func New() stores.Provider {
	return store{}
}

func (s store) Create(ctx *gofr.Context, message models.Provider) (models.Provider, error) {
	message.ChannelRefID = uuid.New()
	res, err := ctx.DB().Exec(INSERT, message.ID, message.Name, message.ChannelRefID)

	if err != nil {
		return models.Provider{}, errors.DB{Err: err}
	}

	row, err := res.RowsAffected()
	if err != nil {
		return models.Provider{}, errors.DB{Err: err}
	}

	if row == 0 {
		return models.Provider{}, errors.EntityAlreadyExists{}
	}

	return message, nil
}

func (s store) GetByID(ctx *gofr.Context, id string) (models.Provider, error) {
	var resp models.Provider
	err := ctx.DB().QueryRowContext(ctx, GETBYID, id).Scan(&resp.ID, &resp.URL)

	if err != nil {
		return models.Provider{}, errors.DB{Err: err}
	}

	return resp, nil
}

func (s store) PUT(ctx *gofr.Context, id string, message models.Provider) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return errors.InvalidParam{Param: []string{id}}
	}

	message.ID = ID
	_, err = ctx.DB().ExecContext(ctx, UPDATE, message.ID, message.URL, message.ChannelRefID)

	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func (s store) Delete(ctx *gofr.Context, id string) error {
	//TODO implement me
	//panic("implement me")
	_, err := ctx.DB().ExecContext(ctx, DELETE, id)

	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
