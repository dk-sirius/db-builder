package gen

type TableGenDef interface {
	// Package 文件包
	Package() string
	// Imports 文件引用
	Imports() string
	// VarDecl TableDefine 表定义
	VarDecl() string
	// Method   表对象的方法
	Method() []string
}
