package gomod_test

import (
	"log"
	"testing"

	"github.com/cdvelop/gomod"
)

func TestGetPackage(t *testing.T) {
	var testData = map[string]struct {
		line   string
		expect string
	}{
		"se espera cutkey":                               {"require github.com/cdvelop/cutkey v0.6.0", "cutkey"},
		"se espera module_medical_history":               {"github.com/cdvelop/module_medical_history v0.0.0-00010101000000-000000000000", "module_medical_history"},
		"se espera model":                                {"github.com/cdvelop/model v0.0.32", "model"},
		"se espera file":                                 {"require github.com/cdvelop/file v0.6.0 // indirect", "file"},
		"contiene replace se espera model ":              {"replace github.com/cdvelop/model => ../model", ""},
		"no se espera nada no contiene version":          {"module github.com/cdvelop/module_client", ""},
		"no se espera nada no cumple con repositorio":    {"go 1.20", ""},
		"no se espera nada sin contenido ni repositorio": {"", ""},
	}

	for testName, testData := range testData {
		t.Run(testName, func(t *testing.T) {
			result := gomod.GetPackage(testData.line)
			if result != testData.expect {
				t.Fatalf("Para entrada '%s', se esperaba '%s' pero se obtuvo '%s'", testData.line, testData.expect, result)
				log.Fatal()
			}
		})
	}
}
