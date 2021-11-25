package gen

import (
	"fmt"
)

func (t *TableGoFileGenerator) VarDecl() string {
	return fmt.Sprintf("var %s = `%s`", t.VarName(), t.desc)
}
