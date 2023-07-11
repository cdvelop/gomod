package gomod_test

import (
	"log"
	"testing"

	"github.com/cdvelop/gomod"
)

func TestGetUsedPackageNames(t *testing.T) {
	const data_gomod_file = `module github.com/cdvelop/module_client

go 1.20

require (
	github.com/cdvelop/model v0.0.32
)

require github.com/cdvelop/cutkey v0.6.0 // indirect

require (
	github.com/cdvelop/api v0.0.1
	github.com/cdvelop/js v0.0.0-20230710200247-0d9f5c8314da
	github.com/cdvelop/platform v0.0.2

)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/godev => ../godev`

	var expected = []string{"model", "cutkey", "api", "js", "platform"}

	owner := "github.com/cdvelop"
	packages := gomod.GetUsedPackageNames(owner, data_gomod_file)

	if len(packages) != len(expected) {
		t.Errorf("Expected %d packages, but got %d\n%v", len(expected), len(packages), packages)
		log.Fatal()
	}

	// Comparar elemento por elemento
	for i, valor := range packages {
		if valor != expected[i] {
			t.Errorf("Expected %v packages, but got %v", expected[i], valor)
			log.Fatal()

		}
	}

}
