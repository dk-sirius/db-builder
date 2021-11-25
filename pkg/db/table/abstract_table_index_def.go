package table

import (
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/token"
)

// TableIndexDef table index define
type TableIndexDef struct {
	// 索引名称
	DefName string
	// 索引类别
	DefClass string
	// 索引关联字段名称
	DefRelativeName []string
}

// AstIndexDef example // @def index i_name name (列别/索引名称/映射字段域名称)
type AstIndexDef string

// just @def and remove @def
func (s AstIndexDef) String() string {
	tmp := string(s)
	if tmp != "" && strings.Contains(tmp, token.Def.String()) {
		tmp = strings.ReplaceAll(tmp, "//", "")
		tmp = strings.ReplaceAll(tmp, token.Def.String(), "")
		tmp = strings.Trim(tmp, " ")
		return tmp
	}
	return ""
}

// Index switch to TableIndexDef
func (s AstIndexDef) Index() *TableIndexDef {
	tmp := s.String()
	if tmp != "" {
		ins := strings.Split(tmp, " ")
		if len(ins) >= 3 {
			return &TableIndexDef{
				DefClass:        ins[0],
				DefName:         ins[1],
				DefRelativeName: ins[2:],
			}
		}
	}
	return nil
}
