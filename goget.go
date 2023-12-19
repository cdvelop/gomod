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
			err := GogetAll()
			fmt.Println(err)

			return true
		}

	}

	return false
}

func GogetAll() (err string) {
	fmt.Println("EJECUTANDO GogetAll...")

	pkg_list, err := GetUsedPackageNames("./")
	if err != "" {
		return err
	}

	fmt.Println("LISTADO PAQUETES USADOS", pkg_list)

	return

}
