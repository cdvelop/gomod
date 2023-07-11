package gomod

import "os"

func Exist() (content string, err error) {

	cont, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	return string(cont), nil
}
