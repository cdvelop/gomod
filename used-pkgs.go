package gomod

import (
	"strings"
)

func GetUsedPackageNames(owner, gomod_file_contend string) (pkg_list []string) {

	lines := strings.Split(gomod_file_contend, "\n")
	reg := map[string]struct{}{}
	for _, line := range lines {
		line = strings.TrimSpace(line)

		pkg_name := GetPackage(line)

		if _, exist := reg[pkg_name]; !exist && pkg_name != "" {

			// fmt.Println(i, " linea: ", line, " PAQUETE: ", pkg_name)

			pkg_list = append(pkg_list, pkg_name)
			reg[pkg_name] = struct{}{}
		}

	}

	return

}
