package db

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/gen"
	"github.com/dk-sirius/db-builder/pkg/db/table"
	"github.com/dk-sirius/db-builder/pkg/utils"
	"golang.org/x/tools/imports"
)

func Generate(dbName, tableName, srcPath string) {
	if srcPath == "" {
		path, err := utils.GetGenerateGoFile()
		if err != nil {
			panic(err)
		}
		srcPath = path
	}

	def := table.NewSqlBuilder(dbName, tableName, srcPath).Build()
	if def != nil {
		sql := def.CreateTable()
		if sql != "" {
			content := gen.TableGenerator(def.Def.Name, sql, tableName, def.AnalysisTable)
			filePath := generatorFileName(srcPath)
			origin, err := imports.Process(filePath, content.Generate(), nil)
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(filePath, origin, 0666)
			if err != nil {
				panic(err)
			}
		}
	}
}

func generatorFileName(src string) string {
	sName := strings.ReplaceAll(filepath.Base(src), filepath.Ext(src), "")
	tPath := fmt.Sprintf("%s%c%s__generated%s", filepath.Dir(src), filepath.Separator, sName, filepath.Ext(src))
	return tPath
}
