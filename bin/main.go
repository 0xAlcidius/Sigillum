package main

import (
	"log"
	"os"
	"sigillum/constants"
	"sigillum/export"
	"sigillum/utils"
	"slices"

	"github.com/alecthomas/kingpin/v2"
)

var (
	app = kingpin.New("Sigillum", "A cryptor for your payload.")

	payload_command  = app.Flag("payload", "Provide the payload by text.").Short('p').Required().String()
	seal_command     = app.Flag("seal", "The type of encryption / obfuscation (e.g., aes, rc4, etc.)").Short('s').Default("RC4").String()
	key_command      = app.Flag("key", "The key to decrypt the payload.").Short('k').Required().String()
	language_command = app.Flag("language", "The outputted programming language.").Short('l').Default("C").String()
	output_command   = app.Flag("output", "Path to save the file.").Short('o').String()
	filename_command = app.Flag("filename", "Filename that should be given once the payload is decrypted and saved to a file.").Short('f').Default("payload.txt").String()
)

func main() {
	var payload []byte
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Printf("Error parsing arguments: %s", err)
		os.Exit(1)
	}

	payload, isFile, err := utils.ParsePayload(*payload_command)
	if err != nil && isFile {
		log.Printf("Could not find file: %s, error: %s", *payload_command, err)
		os.Exit(1)
	}

	key, err := utils.ParseKey(*key_command)
	if err != nil {
		log.Printf("Could not parse key: %s, error: %s", key, err)
		os.Exit(1)
	}

	if !slices.Contains(constants.SupportedLanguages, *language_command) {
		log.Printf("Programming language: %s not supported", *language_command)
		os.Exit(1)
	}

	if encrypt, found := constants.SupportedSeals[*seal_command]; found {
		cipertext, err := encrypt(key, payload)
		if err != nil {
			log.Println("Failed to encrypt payload. Error: ", err)
			os.Exit(1)
		}

		switch *language_command {
		case "C":
			options := export.CreateExportCOptions(key, cipertext, *seal_command, *output_command, *filename_command)
			if *output_command != "" {
				export.ExportC(options)
			} else {
				export.PrintC(options)
			}
		}
	}

	os.Exit(0)
}
