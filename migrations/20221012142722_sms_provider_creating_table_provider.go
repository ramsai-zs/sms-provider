package migrations

import (
	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/log"
)

type K20221012142722 struct {
}

func (k K20221012142722) Up(d *datastore.DataStore, logger log.Logger) error {
	return nil
}

func (k K20221012142722) Down(d *datastore.DataStore, logger log.Logger) error {
	return nil
}
