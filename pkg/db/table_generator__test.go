package db_test

import (
	"testing"

	"github.com/dk-sirius/db-builder/pkg/db"
)

func TestGenerate(t *testing.T) {
	path := "/Users/dunbar/workspace/go_workspace/src/github.com/dk-sirius/db-builder/pkg/db/example/account.go"
	dbName := "test"
	tableName := "Account"
	db.Generate(dbName, tableName, path)
}
