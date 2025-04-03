package utils

import (
	"log"
	"os"
)

func ParseKey(key string) ([]byte, error) {
	return []byte(key), nil
}

func ParsePayload(payload string) ([]byte, bool, error) {
	file, err := os.Open(payload)
	if err != nil {
		log.Print("No\n")
		return []byte(payload), false, err
	}

	info, err := file.Stat()
	if err != nil {
		return nil, true, err
	}

	bytes := make([]byte, info.Size())

	file.Read(bytes)

	return bytes, true, nil
}

func ParseText(text string) ([]byte, error) {
	return []byte(text), nil

}
