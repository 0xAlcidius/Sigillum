package sealing

import (
	"crypto/rc4"
)

func RC4CreateSeal(key []byte, shellcode []byte) ([]byte, error) {

	cipertext, err := createCipertext(key, shellcode)
	if err != nil {
		return nil, err
	}
	return cipertext, nil
}

func createCipertext(key []byte, shellcode []byte) ([]byte, error) {
	cipertext := make([]byte, len(shellcode))

	c, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	c.XORKeyStream(cipertext, shellcode)

	return cipertext, nil
}
