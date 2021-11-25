package table

import (
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/token"
)

type ConstraintDef struct {
	ConstraintKey    string
	ConstraintValues []string
}

// AstConstraintDef example: // @def primary ID
type AstConstraintDef string

func (astP AstConstraintDef) String() string {
	cIndex := AstIndexDef(astP)
	if cIndex != "" {
		return cIndex.String()
	}
	return ""
}

func (astP AstConstraintDef) Constraint() *ConstraintDef {
	tmp := astP.String()
	if tmp != "" {
		ins := strings.Split(tmp, " ")
		if len(ins) > 0 {
			if ins[0] != token.Index.String() && ins[0] != token.Unique.String() {
				return &ConstraintDef{
					ConstraintKey:    ins[0],
					ConstraintValues: ins[1:],
				}
			}
		}
	}
	return nil
}
