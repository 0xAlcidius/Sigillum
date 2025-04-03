package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sigillum/constants"
	"strings"
)

func GetPath(sealtype string, language string) (string, error) {
	identifier := strings.ToLower(sealtype + "." + language)

	root, err := getProjectRoot()
	if err != nil {
		return "", err
	}

	if !strings.Contains(root, constants.PROJECTNAME) {
		err = os.Chdir(constants.PROJECTNAME)
		if err != nil {
			return "", errors.New(fmt.Sprint("Could not find this project: ", constants.PROJECTNAME, " in root: ", root, " and could not change into the directory of ", constants.PROJECTNAME))
		}
	}

	root, lastDir := filepath.Split(root)

	for lastDir != constants.PROJECTNAME {
		root, lastDir = filepath.Split(root[:len(root)-1])
	}

	path := filepath.Join(root, constants.DESEALINGPATH, identifier)
	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	return path, nil
}

func getProjectRoot() (string, error) {
	ex, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return ex, nil
}
