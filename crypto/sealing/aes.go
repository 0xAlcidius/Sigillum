package sealing

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type AESSeal struct{}

const MAXKEYSIZE = 32

func (AESSeal) ExecuteSeal(key []byte, payload []byte) ([]byte, error) {
	paddingLength := 0
	if len(key) < MAXKEYSIZE {
		key, paddingLength = addPadding(key, MAXKEYSIZE)
	} else if len(key) > MAXKEYSIZE {
		return nil, errors.New("key was too big. Can be a maximum of 32 characters with AES")
	}

	cBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(payload)%(cBlock.BlockSize()) != 0 {
		payload, paddingLength = addPadding(payload, cBlock.BlockSize())
	}

	ciphertext := make([]byte, len(payload))
	iv := make([]byte, cBlock.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(cBlock, iv)
	mode.CryptBlocks(ciphertext, payload)

	ciphertext = append(iv, ciphertext...)
	return append(ciphertext, byte(paddingLength)), nil
}

func addPadding(item []byte, modulus int) ([]byte, int) {
	paddingLength := modulus - (len(item) % modulus)

	for i := 0; i < paddingLength; i++ {
		item = append(item, 0xFF)
	}
	return item, paddingLength
}
