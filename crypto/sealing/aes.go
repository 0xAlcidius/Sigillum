package sealing

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

const MAXKEYSIZE = 32

func AESCreateSeal(key []byte, payload []byte) ([]byte, error) {

	if len(key) < MAXKEYSIZE {
		key = addPadding(key, MAXKEYSIZE)
	} else if len(key) > MAXKEYSIZE {
		return nil, errors.New("key was too big. Can be a maximum of 32 characters with AES")
	}

	cBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(payload)%(cBlock.BlockSize()) != 0 {
		payload = addPadding(payload, cBlock.BlockSize())
	}

	ciphertext := make([]byte, len(payload))
	iv := make([]byte, cBlock.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(cBlock, iv)
	mode.CryptBlocks(ciphertext, payload)

	return append(iv, ciphertext...), nil
}

func addPadding(item []byte, modulus int) []byte {
	paddingLength := modulus - (len(item) % modulus)

	for i := 0; i < paddingLength; i++ {
		item = append(item, byte(0))
	}
	return item
}
