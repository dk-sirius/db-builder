package table

import (
	"fmt"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/token"
)

// TableFieldDef table field define
type TableFieldDef struct {
	// 字段类型
	DefType string
	// 字段名称
	DefName string
	// 字段描述
	DefDesc map[string]string
}

// AstFieldDef example `db:"name,size=10,default='0'"`
type AstFieldDef string

// just `db` and remove `db`、`"`
func (s AstFieldDef) String() string {
	tmp := string(s)
	if tmp != "" && strings.Contains(tmp, token.Db.String()) {
		tmp = strings.ReplaceAll(tmp, "`", "")
		tag := strings.Split(tmp, ":")
		if len(tag) != 2 {
			panic(fmt.Sprintf("%s is a illegal field define", tmp))
		}
		return strings.ReplaceAll(tag[1], "\"", "")
	}
	return ""
}

func (s AstFieldDef) Field() *TableFieldDef {
	tmp := s.String()
	if tmp != "" {
		field := &TableFieldDef{}
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
