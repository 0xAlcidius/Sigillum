package main

import (
	"os"
	"strings"

	"github.com/0xAlcidius/Sigillum/sigillum"
	"github.com/0xAlcidius/Sigillum/utils"

	"github.com/0xAlcidius/Sigillum/export"

	"github.com/0xAlcidius/Sigillum/crypto/desealing"

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
	desealing.GetDesealPath("RC4", "C")
	var payload []byte
	_, err := app.Parse(os.Args[1:])
	kingpin.FatalIfError(err, "Error parsing arguments")

	payload, isFile, err := utils.ParsePayload(*payload_command)
	if err != nil && isFile {
		kingpin.FatalIfError(err, "Could not find file")
	}

	key, err := utils.ParseKey(*key_command)
	kingpin.FatalIfError(err, "Could not parse key")

	seal, found := sigillum.Seal[strings.ToUpper(*seal_command)]

	if !found {
		kingpin.Fatalf("Sealing algorithm not supported")
	}

	ciphertext, err := seal.ExecuteSeal(key, payload)
	kingpin.FatalIfError(err, "Failed to seal payload")

	exportciphertext, found := sigillum.Language[strings.ToUpper(*language_command)]

	if !found {
		kingpin.Fatalf("Programming language not supported")
	}

	options := export.CreateExportOptions(key, ciphertext, *seal_command, *language_command, *output_command, *filename_command)
	err = exportciphertext(options)
	kingpin.FatalIfError(err, "Failed to export ciphertext")
}
