package sql

import (
	"context"
	"database/sql"
)

type Database interface {
	Get(ctx context.Context, dest interface{}, q string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, q string, args ...interface{}) error
	Exec(ctx context.Context, q string, args ...interface{}) (sql.Result, error)
}

type Record interface {
	TableName() string  // the table name which record belongs to.
	PrimaryKey() string // the primary key of the record.
}

type Sorted interface {
	SortBy() string // column names separated by comma.
}

type Scannable interface {
	Scan(*sql.Rows) error
}
