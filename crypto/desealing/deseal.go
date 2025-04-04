package desealing

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"
)

func GetDesealPath(seal string, language string) (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("error getting path")
	}

	dir, _ := filepath.Split(filename)
	dir = dir + strings.ToLower(language) + "/" + strings.ToLower(seal) + "." + strings.ToLower(language)
	return dir, nil
}
