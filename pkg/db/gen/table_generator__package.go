package gen

import (
	"fmt"
)

func (t *TableGoFileGenerator) Package() string {
	return fmt.Sprintf("package %s", t.PName())
}
