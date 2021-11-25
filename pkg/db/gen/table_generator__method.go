package gen

import (
	"fmt"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/utils"
)

func (t *TableGoFileGenerator) TableName() string {
	return fmt.Sprintf("func (%s)TableName()string{\n return \"%s\"}", t.OName(), t.TName())
}

func (t *TableGoFileGenerator) Migrate() string {
	return fmt.Sprintf("func (%s)Migrate(db *sqlx.DB) sql.Result{\n return db.MustExec(%s)}", t.OName(), t.VarName())
}

func (t *TableGoFileGenerator) Schema() string {
	return fmt.Sprintf("func (%s)Schema()string{\n return %s}", t.OName(), t.VarName())
}

func (t *TableGoFileGenerator) PrimaryKeys() string {
	v := make([]string, 0)
	for i, _ := range t.primaryFieldKey {
		v = append(v, fmt.Sprintf("\"%s\"", t.primaryFieldKey[i]))
	}
	kys := strings.Join(v, ",")
	return fmt.Sprintf("func (%s)PrimaryKeys()[]string{\n return []string {%v}}", t.OName(), kys)
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
		s = append(s, fmt.Sprintf("func (%s)Field%sKey()string{\n return \"%s\"}", t.OName(), strings.Join(keys, ""), t.fieldKey[i]))
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
				s = append(s, fmt.Sprintf("func (%s)Index%sKey()string{\n return \"%s\"}", t.OName(), key, k))
			}
		}
	}
	return strings.Join(s, "\n")
}

func (t *TableGoFileGenerator) IndexKeyValue() string {
	v := make([]string, 0)
	if t.indexFieldKey != nil {
		for i, _ := range t.indexFieldKey {
			for k, va := range t.indexFieldKey[i] {
				key := indexName(k)
				v = append(v, fmt.Sprintf("func (t %s)Index%sValue()[]string{\n  return t.ToIndexSlice(\"%s\")}", t.OName(), key, va))
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
				s = append(s, fmt.Sprintf("func (%s)UniqueIndex%sKey()string{\n return \"%s\"}", t.OName(), key, k))
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
				v = append(v, fmt.Sprintf("func (t %s)UniqueIndex%sValue()[]string{\n return t.ToIndexSlice(\"%s\")}", t.OName(), key, va))
			}
		}
	}
	return strings.Join(v, "\n")
}

func (t *TableGoFileGenerator) ToIndexSlice() string {
	return fmt.Sprintf("func (%s)ToIndexSlice(s string)[]string{\n return strings.Split(s, \",\")}", t.OName())
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
	s = append(s, t.ToIndexSlice())
	return s
}
