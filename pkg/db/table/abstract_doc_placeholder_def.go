package table

import (
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/token"
)

type DocPlaceHolderDef struct {
	PlaceHolderKey    string
	PlaceHolderValues []string
}

// AstHolderPlaceDef example: // @def primary ID
type AstHolderPlaceDef string

func (astP AstHolderPlaceDef) String() string {
	cIndex := AstIndexDef(astP)
	if cIndex != "" {
		return cIndex.String()
	}
	return ""
}

func (astP AstHolderPlaceDef) Constraint() *DocPlaceHolderDef {
	tmp := astP.String()
	if tmp != "" {
		ins := strings.Split(tmp, " ")
		if len(ins) > 0 {
			if ins[0] != token.Index.String() && ins[0] != token.Unique.String() {
				return &DocPlaceHolderDef{
					PlaceHolderKey:    ins[0],
					PlaceHolderValues: ins[1:],
				}
			}
		}
	}
	return nil
}
