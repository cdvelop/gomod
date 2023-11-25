package gomod

import (
	"strings"
)

func GetUsedPackageNames(path string) (pkg_list []string, err string) {

	gomod_content, err := Exist(path)
	if err != "" {
		return nil, err
	}

	owner := GetRepositoryOwner(gomod_content)
	if owner == "" {
		return nil, "error no se logro obtener el dueño del repositorio gomod path: " + path
	}

	lines := strings.Split(gomod_content, "\n")
	reg := map[string]struct{}{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		pkg_name := GetPackage(line)

		if _, exist := reg[pkg_name]; !exist && pkg_name != "" && strings.Contains(line, owner) {

			// fmt.Println(" linea: ", line, " PAQUETE: ", pkg_name)

			pkg_list = append(pkg_list, pkg_name)
			reg[pkg_name] = struct{}{}
		}

	}

	return pkg_list, ""

}

func GetSeparateUsedPackageNames(path string) (modules, components []string, err string) {

	packages_names, err := GetUsedPackageNames(path)
	if err != "" {
		return nil, nil, err
	}

	modules, components = SeparatePackageNames(packages_names)

	return
}

func SeparatePackageNames(packages_names []string) (modules, components []string) {

	for _, name := range packages_names {

		if strings.Contains(name, "module") {
			modules = append(modules, name)
		} else {
			components = append(components, name)
		}
	}

	return
}
