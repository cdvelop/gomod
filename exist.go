package gomod

import (
	"github.com/cdvelop/gotools"
)

// dir ej: "test_project" default ""
func Exist(dir string) (content string, err error) {

	return gotools.FindFirstFileWithExtension(dir, ".mod")
}
