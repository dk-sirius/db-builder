package example

import (
	"fmt"
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

func PrintAst(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fc, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fset := token.NewFileSet()
	astF, err := parser.ParseFile(fset, "x.go", fc, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, astF)
}

// AstPrint /*
func AstPrint() {
	p := "/Users/dunbar/workspace/go_workspace/src/gitee.com/ksirius/tools/pkg/db/example/account.go"
	fmt.Print("\r\n\n")
	PrintAst(p)
	fmt.Print("\r\n\n")
	//f := TraversalAstFile(p)
	//for i, _ := range f {
	//	fmt.Println(f[i].Names, f[i].Tag.Value)
	//}
}

func OpenAstFile(path string) (*ast.File, string) {
	lf, err := os.Open(path)
	defer lf.Close()
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

func TraversalAstFile(path string) []*ast.Field {
	parentPath := filepath.Dir(path)
	f, name := OpenAstFile(path)
	if f != nil {
		fs := make([]*ast.Field, 0)
		ast.Inspect(f, func(node ast.Node) bool {
			if field, ok := node.(*ast.Field); ok {
				switch t := field.Type.(type) {
				case *ast.Ident:
					// 相同包
					if field.Names != nil && field.Type != nil && field.Tag != nil {
						// 有效输出
						fs = append(fs, field)
					} else {
						// 相同目录中寻找类型
						if f.Scope.Lookup(t.Name) == nil {
							tf := TraversalAstDir(t.Name, parentPath, name)
							if tf != "" {
								fs = append(fs, TraversalAstFile(tf)...)
							}
						}
					}
				case *ast.SelectorExpr:
					// 不同包
					path := SwitchImportFilePath(f, t)
					path = strings.ReplaceAll(path, "\"", "")
					tf := TraversalAstDir(t.Sel.Name, path, "")
					if tf != "" {
						fs = append(fs, TraversalAstFile(tf)...)
					}
				}
			}
			return true
		})
		return fs
	}
	return nil
}

func TraversalAstDir(typeName, path, traversalFile string) string {
	pkgs, err := parser.ParseDir(token.NewFileSet(), path, func(info fs.FileInfo) bool {
		if info.IsDir() || info.Name() == traversalFile {
			return false
		}
		fmt.Println(info.Name())
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
					if tp.Name.Name == typeName {
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

func SwitchImportFilePath(f *ast.File, expr *ast.SelectorExpr) string {
	if name, ok := expr.X.(*ast.Ident); ok {
		for i, _ := range f.Imports {
			if f.Imports[i].Name.Name == name.Name {
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
