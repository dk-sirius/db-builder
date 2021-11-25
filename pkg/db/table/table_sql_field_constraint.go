package table

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/token"
)

// Constraint FieldConstraint 字段约束
type Constraint map[string]string

func (f Constraint) String() string {
	checkGroup := make([]string, 0)
	checkGroup = append(checkGroup, "NOT NULL")
	for k, v := range f {
		ck := token.Lookup(k)
		if ck.IsFieldCheck() {
			switch ck {
			case token.ValueDefault:
				checkGroup = append(checkGroup, fmt.Sprintf("DEFAULT %s", v))
			case token.Primary:
				// just support single primary key
				checkGroup = append(checkGroup, fmt.Sprintf("PRIMARY KEY"))
			}
		}
	}
	return strings.Join(checkGroup, " ")
}

func valid(key string) (token.DbToken, bool) {
	ck := token.Lookup(key)
	return ck, ck.IsFieldCheck()
}

func (f Constraint) IncludeSequence() bool {
	for k, _ := range f {
		if token.Lookup(k).IsSequenceCheck() {
			return true
		}
	}
	return false
}

func (f Constraint) HasSizeConstraint() string {
	for k, v := range f {
		if ck, ok := valid(k); ok {
			if ck == token.Size {
				size, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					panic(err)
				}
				return ToVarchar(size)
			}
		}
	}
	return ""
}

func ToVarchar(size int64) string {
	return fmt.Sprintf("varchar(%d)", size)
}
