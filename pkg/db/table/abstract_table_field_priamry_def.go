package table

import (
	"fmt"
	"strings"
)

type PrimaryDef []string

func (p PrimaryDef) String() string {
	if p != nil {
		return strings.Join(p, ",")
	}
	return ""
}

func (p PrimaryDef) Primary() string {
	tmp := p.String()
	if tmp != "" {
		return fmt.Sprintf("PRIMARY KEY(%s)", tmp)
	}
	return ""
}
