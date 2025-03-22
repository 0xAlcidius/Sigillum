package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sigillum/constants"
	"strings"
)

func GetPath(sealtype string, language string) (string, error) {
	identifier := strings.ToLower(sealtype + "_" + language)

	switch identifier {
	case "rc4_c":
		filePath, err := getProjectRoot()
		if err != nil {
			return "", err
		}
		return filepath.Join(filePath, constants.RC4DECRYPTIONPATH), nil
	default:
		fmt.Println("Usage GetPath(sealtype, language) (e.g., GetPath(RC4, C)")
		return "", nil
	}
}

func getProjectRoot() (string, error) {
	ex, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}
