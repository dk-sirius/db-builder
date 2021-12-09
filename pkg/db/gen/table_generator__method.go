package gen

import (
	"fmt"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/utils"
)

func (t *TableGoFileGenerator) TableName() string {
	k := NewExpr("TableName", t.OName(), t.TName())
	return k.Gen()
}

func (t *TableGoFileGenerator) Migrate() string {
	return fmt.Sprintf("func (%s)Migrate(db *sqlx.DB) sql.Result{\n return db.MustExec(%s)}", t.OName(), t.VarName())
}

func (t *TableGoFileGenerator) Schema() string {
	k := NewExpr("Schema", t.OName(), t.VarName()).ToKeep(true)
	return k.Gen()

}

func (t *TableGoFileGenerator) PrimaryKeys() string {
	k := NewExpr("PrimaryKeys", t.OName(), t.primaryFieldKey)
	return k.Gen()
}

func (t *TableGoFileGenerator) FieldKey() string {
	s := make([]string, 0)
	for i, _ := range t.fieldKey {
		// remove f_
		key := strings.ReplaceAll(t.fieldKey[i], "f_", "")
		subs := strings.Split(key, "_")
		keys := make([]string, 0)
		for j, _ := range subs {
			keys = append(keys, utils.Capitalize(subs[j]))
		}
		name := fmt.Sprintf("Field%sKey", strings.Join(keys, ""))
		k := NewExpr(name, t.OName(), t.fieldKey[i])
		s = append(s, k.Gen())
	}
	return strings.Join(s, "\n")
}

func indexName(s string) string {
	// remove i_
	key := strings.ReplaceAll(s, "i_", "")
	subs := strings.Split(key, "_")
	keys := make([]string, 0)
	for j, _ := range subs {
		keys = append(keys, utils.Capitalize(subs[j]))
	}
	return strings.Join(keys, "")
}

func (t *TableGoFileGenerator) IndexKey() string {
	s := make([]string, 0)
	if t.indexFieldKey != nil {
		for i, _ := range t.indexFieldKey {
			for k, _ := range t.indexFieldKey[i] {
				key := indexName(k)
				name := fmt.Sprintf("Index%sKey", key)
				f := NewExpr(name, t.OName(), k)
				s = append(s, f.Gen())
			}
		}
	}
	return strings.Join(s, "\n")
}

func (t *TableGoFileGenerator) FieldKeys() string {
	f := NewExpr("FieldKeys", t.OName(), t.fieldKey)
	return f.Gen()
}

func (t *TableGoFileGenerator) IndexKeyValue() string {
	v := make([]string, 0)
	if t.indexFieldKey != nil {
		for i, _ := range t.indexFieldKey {
			for k, va := range t.indexFieldKey[i] {
				key := indexName(k)
				name := fmt.Sprintf("Index%sValue", key)
				f := NewExpr(name, t.OName(), va)
				v = append(v, f.Gen())
			}
		}
	}
	return strings.Join(v, "\n")
}

func (t *TableGoFileGenerator) UniqueIndexKey() string {
	s := make([]string, 0)
	if t.uniqueIndexFieldKey != nil {
		for i, _ := range t.uniqueIndexFieldKey {
			for k, _ := range t.uniqueIndexFieldKey[i] {
				// remove i_
				key := indexName(k)
				name := fmt.Sprintf("UniqueIndex%sKey", key)
				f := NewExpr(name, t.OName(), k)
				s = append(s, f.Gen())
			}
		}
	}
	return strings.Join(s, "\n")
}

func (t *TableGoFileGenerator) UniqueIndexValue() string {
	v := make([]string, 0)
	if t.uniqueIndexFieldKey != nil {
		for i, _ := range t.uniqueIndexFieldKey {
			for k, va := range t.uniqueIndexFieldKey[i] {
				key := indexName(k)
				name := fmt.Sprintf("UniqueIndex%sValue", key)
				f := NewExpr(name, t.OName(), va)
				v = append(v, f.Gen())
			}
		}
	}
	return strings.Join(v, "\n")
}

func (t *TableGoFileGenerator) AutoFieldKeys() string {
	return NewExpr("AutoFieldKeys", t.OName(), t.autoFieldKey).Gen()
}

func (t *TableGoFileGenerator) Method() []string {
	s := make([]string, 0)
	s = append(s, t.TableName())
	s = append(s, t.Migrate())
	s = append(s, t.Schema())
	s = append(s, t.PrimaryKeys())
	s = append(s, t.FieldKey())
	s = append(s, t.IndexKey())
	s = append(s, t.IndexKeyValue())
	s = append(s, t.UniqueIndexKey())
	s = append(s, t.UniqueIndexValue())
	s = append(s, t.FieldKeys())
	s = append(s, t.AutoFieldKeys())
	return s
}
