package tests

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/0xAlcidius/Sigillum/export"
	"github.com/0xAlcidius/Sigillum/sigillum"
)

func TestAESSealingText(t *testing.T) {
	var language string = "C"
	filePath := filepath.Join(TEMPDIR, "aes.c")

	seal := sigillum.Seal["AES"]

	ciphertext, err := seal.ExecuteSeal([]byte(KEY), []byte(PAYLOAD))

	if err != nil {
		t.Error("Error creating ciphertext in AES test. Err:", err)
	}

	options := export.CreateExportOptions([]byte(KEY), ciphertext, "AES", language, filePath, EXPORT_NAME)
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

/* To run this test successfully, please make sure gcc is installed on the system this test will be performed on. */
func TestAESCompilationC(t *testing.T) {
	compile := exec.Command("gcc", TEMPDIR+"\\aes.c", "-o", TEMPDIR+"\\aes.exe", "-lbcrypt", "-mconsole")
	output, err := compile.CombinedOutput()
	if err != nil {
		t.Fatalf("Compilation of AES in C failed: %s\nOutput: %s", err, string(output))
	}

	runAes := exec.Command(".\\" + TEMPDIR + "\\aes.exe")
	output, err = runAes.CombinedOutput()
	if err != nil {
		t.Fatalf("Compilation of AES in C failed: %s\nOutput: %s", err, string(output))
	}

	if !strings.Contains(string(output), PAYLOAD) {
		fmt.Printf("Output was not similar to input\nExpected: %s\nOutput: %s\n", PAYLOAD, string(output))
		t.Fatalf("Output was not similar to input\nExpected: %s\nOutput: %s\n", PAYLOAD, string(output))
	}
	t.Cleanup(cleanup)
}
