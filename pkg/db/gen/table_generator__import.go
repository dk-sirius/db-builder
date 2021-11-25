package gen

import (
	"fmt"
	"strings"
)

// Imports  /** 生成文件的import
func (t *TableGoFileGenerator) Imports() string {
	s := make([]string, 0)
	imports := t.IValues()
	for i, _ := range imports {
		s = append(s, fmt.Sprintf("import \"%s\"", imports[i]))
	}
	if len(s) > 0 {
		return strings.Join(s, "\r\n")
	}
	return ""
}
