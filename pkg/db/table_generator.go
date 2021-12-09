package db

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/gen"
	"github.com/dk-sirius/db-builder/pkg/db/table"
	"github.com/dk-sirius/db-builder/pkg/utils"
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
			err := ioutil.WriteFile(filePath, content.Generate(), 0666)
			if err == nil {
				format(filePath)
			} else {
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

func format(path string) {
	cmd := fmt.Sprintf("go fmt %s", path)
	_, err := utils.ChickenRun(cmd)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s \nhas run go fmt ...", path)
	}
}
