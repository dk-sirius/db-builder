package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ChickenRun(targetCmd string) ([]byte, error) {
	if targetCmd != "" {
		chicken := exec.Command("sh", "-c", targetCmd)
		return chicken.CombinedOutput()
	} else {
		panic("cmd is empty")
	}
}

func GetGenerateGoFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	n, err := ChickenRun("echo $GOFILE")
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(fmt.Sprintf("%s%c%s", dir, filepath.Separator, n), "\n", ""), nil
}

func GetWorkSpace() (string, error) {
	n, err := ChickenRun("echo $GOPATH")
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(string(n), "\n", ""), nil
}

func SwitchImportPathToPath(importPath string) (string, error) {
	work, err := GetWorkSpace()
	if err != nil {
		return "", err
	}
	pt := fmt.Sprintf("%s%c%s%c%s", work, filepath.Separator, "src", filepath.Separator, importPath)
	return pt, nil
}
