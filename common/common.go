package common

import (
	"sigillum/crypto/encryption"
)

var SupportedSeals = map[string]func([]byte, []byte) ([]byte, error){
	"RC4": encryption.RC4CreateSeal,
	"XOR": encryption.XORCreateSeal,
}

var SupportedLanguages = map[string]func([]byte, *[]byte) ([]byte, error){}
