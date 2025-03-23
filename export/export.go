package export

/*
Options for exporting:
  - key: 		The key provided by the user used for sealing;
  - cipertext:	The sealed payload derived from the unsealed payload;
  - seal:		The algorithm used for sealing;
  - outputFile:	The filename of the file and the file extension;
  - exportName:	The name the deseal function will use to save the output in.
*/
type ExportOptions struct {
	key        []byte
	cipertext  []byte
	seal       string
	language   string
	outputFile string
	exportName string
}

func CreateExportOptions(key []byte, cipertext []byte, seal string, language string, outputFile string, exportName string) ExportOptions {
	return ExportOptions{key: key, cipertext: cipertext, seal: seal, language: language, outputFile: outputFile, exportName: exportName}
}
