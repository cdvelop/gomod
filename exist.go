package gomod

import (
	"fmt"
	"os"
)

// dir ej: "test_project" default ""
func Exist(dir ...string) (content string, err error) {

	var gomod_dir = "go.mod"

	for _, v := range dir {
		if v != "" {
			gomod_dir = v
		}
	}

	cont, err := os.ReadFile(gomod_dir)
	if err != nil {
		return "", fmt.Errorf("error archivo %v no encontrado", gomod_dir)
	}

	return string(cont), nil
}
