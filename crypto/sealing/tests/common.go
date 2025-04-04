package tests

import (
	"os"
)

const (
	PAYLOAD     = "Sigillum is cool!"
	KEY         = "Keep it secret!"
	EXPORT_NAME = "output/payload.txt"
	TEMPDIR     = "output"
)

func cleanup() {
	os.RemoveAll(TEMPDIR)
}
