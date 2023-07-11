package gomod

import "regexp"

func GetRepositoryOwner(gomod_file_contend string) string {
	// Expresión regular para buscar el patrón 'module' seguido de un espacio y capturar el dueño del repositorio.
	re := regexp.MustCompile(`^module\s+([^\s/]+/[^\s/]+)`)
	match := re.FindStringSubmatch(gomod_file_contend)

	if len(match) > 1 {
		return match[1]
	}

	return ""
}
