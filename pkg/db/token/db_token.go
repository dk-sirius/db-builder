package token

import (
	"strconv"
)

type DbToken int

const (
	ddlKeyBeg DbToken = iota
	Indent
	Db
	Def
	Index
	Unique
	ddlKeyEnd

	ddlCheckBeg
	Primary
	Size
	ValueDefault
	AutoIncr
	ddlCheckEnd
)

// sequence
var dTokens = [...]string{
	Indent:       "Indent",
	Db:           "db",
	Def:          "@def",
	Index:        "index",
	Unique:       "unique_index",
	Primary:      "primary",
	Size:         "size",
	ValueDefault: "default",
	AutoIncr:     "autoincrement",
}

func (tok DbToken) String() string {
	s := ""
	if 0 <= tok && tok < DbToken(len(dTokens)) {
		s = dTokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

var keywords map[string]DbToken

func init() {
	keywords = make(map[string]DbToken)
	for i := ddlKeyBeg + 1; i < ddlCheckEnd; i++ {
		keywords[dTokens[i]] = i
	}
}

func Lookup(ident string) DbToken {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return Indent
}

func (tok DbToken) IsKeyWord() bool {
	return ddlKeyBeg < tok && tok < ddlKeyEnd
}

func (tok DbToken) IsFieldCheck() bool {
	return ddlCheckBeg < tok && tok < ddlCheckEnd
}
