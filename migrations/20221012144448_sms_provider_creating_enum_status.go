package migrations

import (
	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/log"
)

type K20221012144448 struct {
}

func (k K20221012144448) Up(d *datastore.DataStore, logger log.Logger) error {
	_, err := d.DB().Exec(createStatusEnum)
	if err != nil {
		return err
	}

	return nil
}

func (k K20221012144448) Down(d *datastore.DataStore, logger log.Logger) error {
	_, err := d.DB().Exec(dropStatusEnum)
	if err != nil {
		return err
	}

	return nil
}
