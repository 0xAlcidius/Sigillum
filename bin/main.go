package main

import (
	"os"
	"sigillum/export"
	"sigillum/support"
	"sigillum/utils"
	"strings"

	"github.com/alecthomas/kingpin/v2"
)

var (
	app = kingpin.New("Sigillum", "A cryptor for your payload.")

	payload_command  = app.Flag("payload", "Provide the payload by text.").Short('p').Required().String()
	seal_command     = app.Flag("seal", "The type of encryption / obfuscation (e.g., xor, rc4, etc.)").Short('s').Default("RC4").String()
	key_command      = app.Flag("key", "The key to decrypt the payload.").Short('k').Required().String()
	language_command = app.Flag("language", "The outputted programming language.").Short('l').Default("C").String()
	output_command   = app.Flag("output", "Path to save the file.").Short('o').String()
	filename_command = app.Flag("filename", "Filename that should be given once the payload is decrypted and saved to a file.").Short('f').Default("payload.txt").String()
)

func main() {
	var payload []byte
	_, err := app.Parse(os.Args[1:])
	kingpin.FatalIfError(err, "Error parsing arguments")

	payload, isFile, err := utils.ParsePayload(*payload_command)
	if err != nil && isFile {
		kingpin.FatalIfError(err, "Could not find file")
	}

	key, err := utils.ParseKey(*key_command)
	kingpin.FatalIfError(err, "Could not parse key")

	encrypt, found := support.SupportedSeals[strings.ToUpper(*seal_command)]

	if !found {
		kingpin.Fatalf("Sealing algorithm not supported")
	}

	cipertext, err := encrypt(key, payload)
	kingpin.FatalIfError(err, "Failed to encrypt payload")

	exportCipertext, found := support.SupportedLanguages[strings.ToUpper(*language_command)]

	if !found {
		kingpin.Fatalf("Programming language not supported")
	}

	options := export.CreateExportOptions(key, cipertext, *seal_command, *language_command, *output_command, *filename_command)
	err = exportCipertext(options)
	kingpin.FatalIfError(err, "Failed to export ciphertext")
}
