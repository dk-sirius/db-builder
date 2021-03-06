package table

import (
	"reflect"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/token"
)

// FieldDef table field define
type FieldDef struct {
	// 字段类型
	DefType string
	// 字段名称
	DefName string
	// 字段描述
	DefDesc map[string]string
}

// AstFieldDef example `db:"name,size=10,default='0'"`
type AstFieldDef string

func (s AstFieldDef) String() string {
	tmp := string(s)
	tmp = strings.ReplaceAll(tmp, "`", "")
	tag := reflect.StructTag(tmp)
	if v, ok := tag.Lookup(token.Db.String()); ok {
		return v
	}
	return ""
}

func (s AstFieldDef) Field() *FieldDef {
	tmp := s.String()
	if tmp != "" {
		field := &FieldDef{}
		field.DefName = tmp
		// exist desc
		ins := strings.Split(tmp, ",")
		if ins != nil && len(ins) >= 1 {
			field.DefName = ins[0]
			field.DefDesc = make(map[string]string)
			for i := 1; i < len(ins); i++ {
				kv := strings.Split(ins[i], "=")
				if kv != nil && len(kv) == 2 {
					field.DefDesc[kv[0]] = kv[1]
				} else {
					field.DefDesc[ins[i]] = ins[i]
				}
			}
		}
		return field
	}
	return nil
}
