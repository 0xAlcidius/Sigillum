package utils

import (
	"os"
)

func ParseKey(key string) ([]byte, error) {
	return []byte(key), nil
}

func ParseFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, info.Size())

	file.Read(bytes)

	return bytes, nil
}

func ParseText(text string) ([]byte, error) {
	return []byte(text), nil

}
