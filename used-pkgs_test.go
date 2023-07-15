package gomod_test

import (
	"log"
	"testing"

	"github.com/cdvelop/gomod"
)

func TestGetUsedPackageNames(t *testing.T) {

	var expected = []string{"model", "cutkey", "api", "js", "platform"}

	packages, err := gomod.GetUsedPackageNames("test")
	if err != nil {
		log.Fatalln(err)
	}

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
