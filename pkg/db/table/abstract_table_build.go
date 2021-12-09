package table

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/utils"
)

/**
*\ *ast.File
	**\ []ast.Decl
		***\ *ast.GenDecl
			****\ *ast.CommentGroup
			****\ token.Token  (import,type,const,var)
			****\ []ast.Spec
				*****\ *ast.ImportSpec
				*****\ *ast.ValueSpec
				*****\ *ast.TypeSpec
					******\ *ast.Ident
					******\ *ast.Expr (*ast.ArrayType,*ast.StructType,*ast.InterfaceType,*ast.MapType...)
						*******\ *ast.StructType
							********\ *ast.FieldList
								*********\ *[]ast.Field
									**********\ []*ast.Ident (Names)
									**********\ ast.Expr     (SqlType)
									**********\ *ast.BasicLit (Tag)
									**********\ ...
				*****\ *ast.FuncType
			****\ ...
		***\ *ast.FuncDecl
	**\ *ast.Scope
	**\ *ast.CommentGroup
	**\ *[]*ast.ImportSpec
	**\ ...
*/

// Def /** parser table description file with return *DBTableDef
func Def(path string) *DBTableDef {
	result := &DBTableDef{}
	tableFile, _ := OpenAstFile(path)
	ast.Inspect(tableFile, func(node ast.Node) bool {
		if indexDef, ok := node.(*ast.GenDecl); ok {
			// pick index comments from genDecl doc
			result.IndexDef = append(result.IndexDef, PickIndexDef(indexDef.Doc)...)
			result.ConstraintDef = append(result.ConstraintDef, PickConstraintDef(indexDef.Doc)...)
		}
		return node != nil
	})
	result.FieldDef = append(result.FieldDef, PickFieldDef(TraversalFile(path, "")...)...)
	if tableFile != nil {
		// pick file name
		result.Name = tableFile.Name.Name
	}
	return result
}

// OpenAstFile /** Open ast.File and its fileName
// path target file full path
func OpenAstFile(path string) (*ast.File, string) {
	lf, err := os.Open(path)
	defer func() {
		err := lf.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	c, err := ioutil.ReadAll(lf)
	if err != nil {
		panic(err)
	}
	f, err := parser.ParseFile(token.NewFileSet(), lf.Name(), c, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return f, filepath.Base(lf.Name())
}

// TraversalAstDir /** Traversal dir to search struct type which is not has traversal file,
// typeName struct type Name;
// path target search dir;
// traversalFile has traversal file;
// any error will panic
func TraversalAstDir(typeName, path, traversalFile string) string {
	pkgs, err := parser.ParseDir(token.NewFileSet(), path, func(info fs.FileInfo) bool {
		if info.IsDir() || info.Name() == traversalFile {
			return false
		}
		return true
	}, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	findFile := ""
	for _, p := range pkgs {
		for f, _ := range p.Files {
			ast.Inspect(p.Files[f], func(node ast.Node) bool {
				if tp, ok := node.(*ast.TypeSpec); ok {
					if isSameObject(tp, typeName) {
						findFile = f
						return false
					}
				}
				return true
			})
		}
	}
	return findFile
}

// compare same object
func isSameObject(node *ast.TypeSpec, targetName string) bool {
	if targetName == "" {
		return true
	}
	return node.Name.Name == targetName
}

// SwitchImportFilePath /** Base $GOPATH to switch import path with file path
func SwitchImportFilePath(f *ast.File, expr *ast.SelectorExpr) string {
	if name, ok := expr.X.(*ast.Ident); ok {
		for i, _ := range f.Imports {
			base := strings.ReplaceAll(filepath.Base(f.Imports[i].Path.Value), "\"", "")
			if (f.Imports[i].Name != nil && f.Imports[i].Name.Name == name.Name) || (base == name.Name) {
				path, err := utils.SwitchImportPathToPath(f.Imports[i].Path.Value)
				if err != nil {
					panic(err)
				}
				return path
			}
		}
	}
	return ""
}

// TraversalFile  /** Traversal file to collect all table define,
// path target search file;
// objName target struct name which can be empty;
// any error will panic
func TraversalFile(path, objName string) []*ast.Field {
	ff, fileName := OpenAstFile(path)
	if ff != nil {
		parentPath := filepath.Dir(path)
		fields := make([]*ast.Field, 0)
		ast.Inspect(ff, func(node ast.Node) bool {
			if tp, ok := node.(*ast.TypeSpec); ok {
				ast.Inspect(tp, func(node ast.Node) bool {
					if field, ok := node.(*ast.Field); ok && isSameObject(tp, objName) {
						switch ft := field.Type.(type) {
						case *ast.Ident:
							if field.Names != nil && field.Tag != nil && field.Type != nil {
								// follow rules
								// pick fields from struct type field
								fields = append(fields, field)
							} else {
								// same package but not same file
								// need find target file to traversal
								if ff.Scope.Lookup(ft.Name) == nil {
									tf := TraversalAstDir(ft.Name, parentPath, fileName)
									if tf != "" {
										fields = append(fields, TraversalFile(tf, ft.Name)...)
									}
								}
							}
						case *ast.SelectorExpr:
							// different package different file
							importPath := SwitchImportFilePath(ff, ft)
							importPath = strings.ReplaceAll(importPath, "\"", "")
							if importPath != "" {
								tf := TraversalAstDir(ft.Sel.Name, importPath, "")
								if tf != "" {
									fields = append(fields, TraversalFile(tf, ft.Sel.Name)...)
								}
							}
						}
					}
					return true
				})
			}
			return true

		})
		return fields
	}
	return nil
}

// PickConstraintDef table constraint from doc
func PickConstraintDef(doc *ast.CommentGroup) []*DocPlaceHolderDef {
	if doc != nil && len(doc.List) > 0 {
		cs := make([]*DocPlaceHolderDef, 0)
		for i, _ := range doc.List {
			def := AstHolderPlaceDef(doc.List[i].Text).Constraint()
			if def != nil {
				cs = append(cs, def)
			}
		}
		return cs
	}
	return nil
}

// PickIndexDef table index
func PickIndexDef(doc *ast.CommentGroup) []*IndexDef {
	if doc != nil && len(doc.List) > 0 {
		cs := make([]*IndexDef, 0)
		for i, _ := range doc.List {
			def := AstIndexDef(doc.List[i].Text).Index()
			if def != nil {
				cs = append(cs, def)
			}
		}
		return cs
	}
	return nil
}

// PickFieldDef table filed
func PickFieldDef(fields ...*ast.Field) []*FieldDef {
	if fields != nil {
		tf := make([]*FieldDef, 0)
		for _, field := range fields {
			if ftp, ok := field.Type.(*ast.Ident); ok {
				if field.Tag != nil {
					ff := AstFieldDef(field.Tag.Value).Field()
					ff.DefType = ftp.Name
					tf = append(tf, ff)
				}
			}
		}
		return tf
	}
	return nil
}
