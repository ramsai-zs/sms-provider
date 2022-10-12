package migrations

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/log"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

// nolint:dupl //Cannot use same testCase for migrateUP and migrateDOWN
func TestK20221012150349_Up(t *testing.T) {
	mock, db := initializeTests(t)
	k := K20221012142722{}

	// register mock calls for success case
	mock.ExpectExec(createTableProvider).WillReturnResult(sqlmock.NewResult(1, 0))
	// register mock calls for failure case
	mock.ExpectExec(createTableProvider).WillReturnError(errors.DB{Err: errors.Error("exec error")})

	testCases := []struct {
		desc string
		err  error
	}{
		{"success", nil},
		{"failure", errors.DB{Err: errors.Error("exec error")}},
	}

	for i, tc := range testCases {
		err := k.Up(&db, log.NewMockLogger(io.Discard))

		assert.IsType(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

// nolint:dupl //Cannot use same testCase for migrateUP and migrateDOWN
func TestK20221012150349_Down(t *testing.T) {
	mock, db := initializeTests(t)
	k := K20221012142722{}

	// register mock calls for success case
	mock.ExpectExec(dropTableProvider).WillReturnResult(sqlmock.NewResult(1, 0))
	// register mock calls for failure case
	mock.ExpectExec(dropTableProvider).WillReturnError(errors.DB{Err: errors.Error("exec error")})

	testCases := []struct {
		desc string
		err  error
	}{
		{"success", nil},
		{"failure", errors.DB{Err: errors.Error("exec error")}},
	}

	for i, tc := range testCases {
		err := k.Down(&db, log.NewMockLogger(io.Discard))

		assert.IsType(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}
