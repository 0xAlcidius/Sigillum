package support

import (
	"sigillum/crypto/sealing"
	"sigillum/export"
)

/* ADD NEWLY CREATED ALGORITHM HERE */
var SupportedSeals = map[string]func([]byte, []byte) ([]byte, error){
	"RC4": sealing.RC4CreateSeal,
	"XOR": sealing.XORCreateSeal,
	"AES": sealing.AESCreateSeal,
}

/* ADD NEWLY SUPPORTED PROGRAMMING LANGUAGES HERE */
var SupportedLanguages = map[string]func(export.ExportOptions) error{
	"C": export.ExportC,
}
