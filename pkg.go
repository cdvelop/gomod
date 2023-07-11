package gomod

import "strings"

func GetPackage(line string) string {
	// Buscar si la línea tiene una versión
	version := GetSemanticVersion(line)

	if version != "" {
		// Obtener el índice de la versión en la línea
		versionIndex := strings.Index(line, version)

		if versionIndex != -1 {
			// Quitar todo el contenido después de la versión
			line = line[:versionIndex]

			// Quitar espacios restantes
			line = strings.TrimSpace(line)

			// Obtener el último componente de la línea restante
			parts := strings.Split(line, "/")
			if len(parts) > 0 {
				return parts[len(parts)-1]
			}
		}
	}

	return ""
}
