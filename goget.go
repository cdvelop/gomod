package gomod

import (
	"fmt"
	"os"

	"github.com/cdvelop/strings"
)

// si encuentra comando lo ejecutamos
func CommandFound() bool {

	for _, v := range os.Args {

		if strings.Contains(v, "goget") != 0 {

			fmt.Println("EJECUTANDO Goget")
			return true
		}

	}

	return false
}

func GogetAll() {

}
