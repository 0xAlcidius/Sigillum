package support

import (
	"github.com/0xAlcidius/Sigillum/crypto/sealing"

	"github.com/0xAlcidius/Sigillum/export"
)

/* ADD NEWLY SUPPORTED ALGORITHM HERE */
var SupportedSeals = map[string]func([]byte, []byte) ([]byte, error){
	"RC4": sealing.RC4CreateSeal,
	"XOR": sealing.XORCreateSeal,
	"AES": sealing.AESCreateSeal,
}

/* ADD NEWLY SUPPORTED PROGRAMMING LANGUAGES HERE */
var SupportedLanguages = map[string]func(export.ExportOptions) error{
	"C": export.ExportC,
}
