package tests

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/0xAlcidius/Sigillum/crypto/sealing"
	"github.com/0xAlcidius/Sigillum/export"
)

func TestRC4SealingText(t *testing.T) {
	var language string = "C"
	filePath := filepath.Join(TEMPDIR, "rc4.c")

	ciphertext, err := sealing.RC4CreateSeal([]byte(KEY), []byte(PAYLOAD))

	if err != nil {
		t.Error("Error creating ciphertext in RC4 test. Err:", err)
	}

	options := export.CreateExportOptions([]byte(KEY), ciphertext, "RC4", language, filePath, EXPORT_NAME)
	export.ExportC(options)

	plaintext, err := sealing.RC4CreateSeal([]byte(KEY), ciphertext)
	if err != nil {
		t.Error("Error creating plaintext in RC4 test. Err:", err)
	}

	if string(plaintext[:len(PAYLOAD)]) != PAYLOAD {
		t.Error("plaintext was not the same as provided payload")
	}
}

/* To run this test successfully, please make sure gcc is installed on the system this test will be performed on. */
func TestRC4CompilationC(t *testing.T) {
	compile := exec.Command("gcc", TEMPDIR+"\\rc4.c", "-o", TEMPDIR+"\\rc4.exe", "-lbcrypt", "-mconsole")
	output, err := compile.CombinedOutput()
	if err != nil {
		t.Fatalf("Compilation of RC4 in C failed: %s\nOutput: %s", err, string(output))
	}

	runRc4 := exec.Command(".\\" + TEMPDIR + "\\rc4.exe")
	output, err = runRc4.CombinedOutput()
	if err != nil {
		t.Fatalf("Compilation of RC4 in C failed: %s\nOutput: %s", err, string(output))
	}

	if !strings.Contains(string(output), PAYLOAD) {
		fmt.Printf("Output was not similar to input\nExpected: %s\nOutput: %s\n", PAYLOAD, string(output))
		t.Fatalf("Output was not similar to input\nExpected: %s\nOutput: %s\n", PAYLOAD, string(output))
	}
	t.Cleanup(cleanup)
}
