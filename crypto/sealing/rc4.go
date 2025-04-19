package sealing

import (
	"crypto/rc4"
)

type RC4Seal struct{}

func (RC4Seal) ExecuteSeal(key []byte, shellcode []byte) ([]byte, error) {

	ciphertext, err := createCiphertext(key, shellcode)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func createCiphertext(key []byte, shellcode []byte) ([]byte, error) {
	ciphertext := make([]byte, len(shellcode))

	c, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	c.XORKeyStream(ciphertext, shellcode)

	return ciphertext, nil
}
