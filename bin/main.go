package main

import (
	"log"
	"os"
	"sigillum/common"
	"sigillum/export"
	"sigillum/utils"

	"github.com/alecthomas/kingpin/v2"
)

var (
	app = kingpin.New("Sigillum", "A cryptor for your shellcode.")

	payload_command  = app.Flag("payload", "Provide the shellcode by text.").Short('t').Required().String()
	seal_command     = app.Flag("seal", "The type of encryption / obfuscation (e.g., aes, rc4, etc.)").Short('s').Default("RC4").String()
	key_command      = app.Flag("key", "The key to decrypt the shellcode.").Short('k').Required().String()
	language_command = app.Flag("language", "The outputted programming language.").Short('l').Default("C").String()
	output_command   = app.Flag("output", "Path to save the file.").Short('o').String()
)

func main() {
	var shellcode []byte
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Print("Error parsing arguments: ", err)
		os.Exit(1)
	}

	shellcode, isFile, err := utils.ParsePayload(*payload_command)
	if err != nil && isFile {
		log.Print("Could not find file: , error: ", *payload_command, err)
		os.Exit(1)
	}

	key, err := utils.ParseKey(*key_command)
	if err != nil {
		log.Print("Could not parse key: , error: ", key, err)
		os.Exit(1)
	}

	if encrypt, found := common.SupportedSeals[*seal_command]; found {
		cipertext, err := encrypt(key, shellcode)
		if err != nil {
			log.Println("Failed to encrypt payload. Error: ", err)
			os.Exit(1)
		}

		if *output_command != "" {

			export.ExportC(key, shellcode, *output_command)
		} else {
			export.PrintC(key, cipertext)
		}
	}

	os.Exit(0)
}
