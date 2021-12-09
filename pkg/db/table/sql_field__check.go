package table

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/token"
)

// TagAttachment Field Check 字段约束
type TagAttachment map[string]string

type FieldCheck struct {
	TagAttachment       // 字段约束描述
	CheckType           // 字段约束
	IsAuto        bool  // 是否自动增长
	Size          int64 // 包含的size 字段值
}

func NewFieldCheck(c TagAttachment) *FieldCheck {
	return &FieldCheck{
		TagAttachment: c,
	}
}

func (f *FieldCheck) Handle() *FieldCheck {
	checkGroup := make([]string, 0)
	checkGroup = append(checkGroup, "NOT NULL")
	for k, v := range f.TagAttachment {
		ck := token.Lookup(k)
		if ck.IsFieldCheck() {
			switch ck {
			case token.ValueDefault:
				checkGroup = append(checkGroup, fmt.Sprintf("DEFAULT %s", v))
			case token.Primary:
				// just support single primary key
				checkGroup = append(checkGroup, fmt.Sprintf("PRIMARY KEY"))
			case token.AutoIncr:
				f.IsAuto = true
			case token.Size:
				size, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					panic(err)
				}
				f.Size = size
			}
		}
	}
	f.CheckType = strings.Join(checkGroup, " ")
	return f
}
