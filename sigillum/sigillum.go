package sigillum

import (
	"github.com/0xAlcidius/Sigillum/crypto/sealing"

	"github.com/0xAlcidius/Sigillum/export"
)

/* ADD NEWLY SUPPORTED ALGORITHM HERE */
var Seals = map[string]sealing.ExecuteSealInterface{
	"RC4": sealing.RC4Seal{},
	"XOR": sealing.XORSeal{},
	"AES": sealing.AESSeal{},
}

/* ADD NEWLY SUPPORTED PROGRAMMING LANGUAGES HERE */
var Languages = map[string]func(export.ExportOptions) error{
	"C": export.ExportC,
}
