package migrations

import (
	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/log"
)

type K20221012150349 struct {
}

func (k K20221012150349) Up(d *datastore.DataStore, logger log.Logger) error {
	_, err := d.DB().Exec(createTableMessage)
	if err != nil {
		return err
	}

	return nil
}

func (k K20221012150349) Down(d *datastore.DataStore, logger log.Logger) error {
	_, err := d.DB().Exec(dropTableMessage)
	if err != nil {
		return err
	}

	return nil
}
