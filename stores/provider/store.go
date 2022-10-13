package provider

import (
	"database/sql"
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

func (s store) Create(ctx *gofr.Context, provider models.Provider) (models.Provider, error) {
	provider.ID = uuid.New()

	// Exec executes a query without returning any rows and inserts the row.
	_, err := ctx.DB().ExecContext(ctx, create, provider.ID, provider.URL, provider.ChannelRefID, provider.Name)
	if err != nil {
		return models.Provider{}, errors.DB{Err: err}
	}

	return provider, nil
}

func (s store) GetByID(ctx *gofr.Context, id string) (models.Provider, error) {
	var p models.Provider

	// QueryRowContext to execute a query with context. QueryRowContext executes a query that return row.
	err := ctx.DB().QueryRowContext(ctx, getByID, id).Scan(&p.ID, &p.URL, &p.ChannelRefID, &p.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Provider{}, errors.EntityNotFound{Entity: "provider", ID: id}
		}

		return models.Provider{}, errors.DB{Err: err}
	}

	return p, nil
}

func (s store) Patch(ctx *gofr.Context, provider models.Provider) error {
	// Exec executes a query without returning any rows and inserts the row.
	_, err := ctx.DB().ExecContext(ctx, update, provider.ID, provider.URL, provider.ChannelRefID, provider.Name)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func (s store) Delete(ctx *gofr.Context, id string) error {
	// Exec executes a query without returning any rows and updates the row.
	_, err := ctx.DB().ExecContext(ctx, deleteProvider, id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
