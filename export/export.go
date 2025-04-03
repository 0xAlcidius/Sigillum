package export

/*
Options for exporting:
  - key: 		The key provided by the user used for sealing;
  - ciphertext:	The sealed payload derived from the unsealed payload;
  - seal:		The algorithm used for sealing;
  - outputFile:	The filename of the file and the file extension;
  - exportName:	The name the deseal function will use to save the output in.
*/
type ExportOptions struct {
	key        []byte
	ciphertext []byte
	seal       string
	language   string
	outputFile string
	exportName string
}

func CreateExportOptions(key []byte, ciphertext []byte, seal string, language string, outputFile string, exportName string) ExportOptions {
	return ExportOptions{key: key, ciphertext: ciphertext, seal: seal, language: language, outputFile: outputFile, exportName: exportName}
}
