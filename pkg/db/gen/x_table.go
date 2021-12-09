package gen

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type XTable interface {
	TableName() string
	Schema() string
	Migrate(db *sqlx.DB) sql.Result
	PrimaryKeys() []string
}
