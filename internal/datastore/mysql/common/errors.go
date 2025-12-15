package common

import (
	"errors"

	"github.com/go-sql-driver/mysql"

	dscommon "github.com/authzed/spicedb/internal/datastore/common"
)

const (
	// mysqlMissingTableErrorNumber is the MySQL error number for "table doesn't exist"
	mysqlMissingTableErrorNumber = 1146
)

// IsMissingTableError returns true if the error is a MySQL error indicating a missing table.
// This typically happens when migrations have not been run.
func IsMissingTableError(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == mysqlMissingTableErrorNumber
}

// WrapMissingTableError checks if the error is a missing table error and wraps it with
// a helpful message instructing the user to run migrations. If it's not a missing table error,
// it returns nil.
func WrapMissingTableError(err error) error {
	if IsMissingTableError(err) {
		return dscommon.NewSchemaNotInitializedError(err)
	}
	return nil
}
