package table

import (
	"fmt"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/token"
)

type SqlIndexDef IndexDef

type IndexCursor uint8

const (
	IndexCursorClass  IndexCursor = iota // 索引类别
	IndexCursorName                      // 索引字段
	IndexCursorFields                    // 索引映射字段
)

func (in SqlIndexDef) String() string {
	index := [...]string{
		IndexCursorClass:  in.conv(),
		IndexCursorName:   in.DefName,
		IndexCursorFields: strings.Join(in.DefRelativeName, ","),
	}
	return strings.Join(index[0:], "$")
}

func (in SqlIndexDef) conv() string {
	def := ""
	switch token.Lookup(in.DefClass) {
	case token.Unique:
		def = "UNIQUE INDEX"
	case token.Index:
		def = "INDEX"
	}
	if def == "" {
		fmt.Println(in)
	}
	return def
}

func (in SqlIndexDef) Index() [3]string {
	return [...]string{
		IndexCursorClass:  in.conv(),
		IndexCursorName:   in.DefName,
		IndexCursorFields: strings.Join(in.DefRelativeName, ","),
	}
}

func (in SqlIndexDef) IsUnique(class string) bool {
	if class != "" {
		return class == token.Unique.String()
	}
	return false
}
