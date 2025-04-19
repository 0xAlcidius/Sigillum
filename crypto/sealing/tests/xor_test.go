package tests

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/0xAlcidius/Sigillum/support"

	"github.com/0xAlcidius/Sigillum/export"
)

func TestXORSealingText(t *testing.T) {
	var language string = "C"
	filePath := filepath.Join(TEMPDIR, "xor.c")

	seal, _ := support.SupportedSeals["XOR"]

	ciphertext, err := seal.ExecuteSeal([]byte(KEY), []byte(PAYLOAD))

	if err != nil {
		t.Error("Error creating ciphertext in XOR test. Err:", err)
	}

	options := export.CreateExportOptions([]byte(KEY), ciphertext, "XOR", language, filePath, EXPORT_NAME)
	export.ExportC(options)

	plaintext, err := seal.ExecuteSeal([]byte(KEY), ciphertext)

	if err != nil {
		t.Error("Error creating plaintext in XOR test. Err:", err)
	}

	if string(plaintext[:len(PAYLOAD)]) != PAYLOAD {
		t.Error("plaintext was not the same as provided payload")
	}
}

/* To run this test successfully, please make sure gcc is installed on the system this test will be performed on. */
func TestXORompilationC(t *testing.T) {
	compile := exec.Command("gcc", TEMPDIR+"\\xor.c", "-o", TEMPDIR+"\\xor.exe", "-lbcrypt", "-mconsole")
	output, err := compile.CombinedOutput()
	if err != nil {
		t.Fatalf("Compilation of XOR in C failed: %s\nOutput: %s", err, string(output))
	}

	runXor := exec.Command(".\\" + TEMPDIR + "\\xor.exe")
	output, err = runXor.CombinedOutput()
	if err != nil {
		t.Fatalf("Compilation of XOR in C failed: %s\nOutput: %s", err, string(output))
	}

	if !strings.Contains(string(output), PAYLOAD) {
		fmt.Printf("Output was not similar to input\nExpected: %s\nOutput: %s\n", PAYLOAD, string(output))
		t.Fatalf("Output was not similar to input\nExpected: %s\nOutput: %s\n", PAYLOAD, string(output))
	}
	t.Cleanup(cleanup)
}
