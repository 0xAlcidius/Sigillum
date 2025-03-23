package constants

import "sigillum/crypto/sealing"

const (
	CIPER    = "// [[CIPERTEXT]]"
	KEY      = "// [[KEY]]"
	FILENAME = "// [[FILENAME]]"

	DESEALINGPATH = "Sigillum/crypto/desealing/"

	RC4DECRYPTIONPATH = "Sigillum/crypto/desealing/rc4.c"
)

/** ADD NEWLY CREATED ALGORITHM HERE **/
var SupportedSeals = map[string]func([]byte, []byte) ([]byte, error){
	"RC4": sealing.RC4CreateSeal,
	"XOR": sealing.XORCreateSeal,
}

var SupportedLanguages = []string{
	"C",
}
