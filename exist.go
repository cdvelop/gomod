package gomod

import (
	"github.com/cdvelop/fileserver"
)

// dir ej: "test_project" default ""
func Exist(dir string) (content, err string) {

	return fileserver.FindFirstFileWithExtension(dir, ".mod")
}
