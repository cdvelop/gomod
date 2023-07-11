package gomod_test

import (
	"testing"

	"github.com/cdvelop/gomod"
)

func TestVersion(t *testing.T) {

	var testData = map[string]struct {
		line   string
		expect string
	}{
		"se espera v0.6.0":                     {"require github.com/cdvelop/cutkey v0.6.0", "v0.6.0"},
		"se espera v0.0.32":                    {"github.com/cdvelop/model v0.0.32", "v0.0.32"},
		"se espera v0.6.0-rc12":                {"github.com/cdvelop/new v0.6.0-rc12", "v0.6.0-rc12"},
		"se espera 0.6.0-rc3":                  {"github.com/cdvelop/old 0.6.0-rc3", "0.6.0-rc3"},
		"versionado go largo":                  {"github.com/lxn/win v0.0.0-20210218163916-a377121e959e", "v0.0.0-20210218163916-a377121e959e"},
		"versionado go largo incompatible":     {"github.com/tdewolff/minify v2.3.6+incompatible", "v2.3.6+incompatible"},
		"formato incorrecto no se espera nada": {"github.com/cdvelop/old 0.6", ""},
		"sin version no se espera resultado":   {"github.com/cdvelop/old", ""},
	}

	for testName, testData := range testData {
		t.Run(testName, func(t *testing.T) {
			result := gomod.GetSemanticVersion(testData.line)
			if result != testData.expect {
				t.Fatalf("Para entrada '%s', se esperaba '%s' pero se obtuvo '%s'", testData.line, testData.expect, result)
			}
		})
	}
}
