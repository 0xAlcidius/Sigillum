package tests

import (
	"crypto/aes"
	"crypto/cipher"
	"path/filepath"
	"sigillum/crypto/sealing"
	"sigillum/export"
	"testing"
)

const (
	PAYLOAD     = "Sigillum is cool!"
	KEY         = "Keep it secret!"
	SEAL        = "AES"
	LANGUAGE    = "C"
	OUTPUT_FILE = "aes.c"
	EXPORT_NAME = "payload.txt"
)

func TestAesText(t *testing.T) {
	tempDir := "output"
	filePath := filepath.Join(tempDir, OUTPUT_FILE)

	ciphertext, err := sealing.AESCreateSeal([]byte(KEY), []byte(PAYLOAD))

	if err != nil {
		t.Error("Error creating ciphertext in AES test. Err:", err)
	}

	options := export.CreateExportOptions([]byte(KEY), ciphertext, SEAL, LANGUAGE, filePath, EXPORT_NAME)
	export.ExportC(options)

	paddingLength := 32 - (len(KEY) % 32)

	paddedKey := []byte(KEY)
	for i := 0; i < paddingLength; i++ {
		paddedKey = append(paddedKey, 0xFF)
	}

	cBlock, err := aes.NewCipher(paddedKey)
	if err != nil {
		t.Error("Could not create new cipher for decryption")
	}

	ciphertext = ciphertext[:len(ciphertext)-1]
	iv := ciphertext[:16]
	ciphertext = ciphertext[16:]

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(cBlock, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	if string(plaintext[:len(PAYLOAD)]) != PAYLOAD {
		t.Error("plaintext was not the same as provided payload")
	}
}
