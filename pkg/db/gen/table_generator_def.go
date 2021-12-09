package gen

import (
	"fmt"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/table"
)

type TableGoFileGenerator struct {
	TableGenDef
	pName               string
	desc                string
	tName               string
	primaryFieldKey     []string
	fieldKey            []string
	uniqueIndexFieldKey []map[string]string
	indexFieldKey       []map[string]string
	autoFieldKey        map[string]bool
}

func TableGenerator(pkg, desc, obj string, tableInfo func() *table.Analysis) *TableGoFileGenerator {
	info := tableInfo()
	return &TableGoFileGenerator{
		pName:               pkg,
		desc:                desc,
		tName:               obj,
		primaryFieldKey:     info.PrimaryFieldKey,
		fieldKey:            info.FieldKey,
		uniqueIndexFieldKey: info.UniqueIndexFieldKey,
		indexFieldKey:       info.IndexFieldKey,
		autoFieldKey:        info.AutoFieldKey,
	}
}

func (gen *TableGoFileGenerator) VarName() string {
	return fmt.Sprintf("_%s", strings.ToLower(gen.tName))
}

func (gen *TableGoFileGenerator) TName() string {
	return fmt.Sprintf("t_%s", strings.ToLower(gen.tName))
}

func (gen *TableGoFileGenerator) OName() string {
	return gen.tName
}

func (gen *TableGoFileGenerator) PName() string {
	return gen.pName
}
func (gen *TableGoFileGenerator) IValues() []string {
	return []string{
		"database/sql",
		"github.com/jmoiron/sqlx",
	}
}
func (t *TableGoFileGenerator) Generate() []byte {
	s := make([]string, 0)
	s = append(s, t.Package())
	s = append(s, t.Imports())
	s = append(s, t.VarDecl())
	s = append(s, t.Method()...)
	return []byte(strings.Join(s, "\r\n"))
}
