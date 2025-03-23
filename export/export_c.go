package export

import (
	"bufio"
	"fmt"
	"os"
	"sigillum/constants"
	"sigillum/utils"
)

func ExportC(options ExportOptions) error {
	var file *os.File = nil
	var err error = nil

	if options.outputFile != "" {
		file, err = os.Create(options.outputFile)
		if err != nil {
			return err
		}
	}

	return parseC(options, file)
}

func parseC(options ExportOptions, file *os.File) error {
	flag := file == nil

	filePath, err := utils.GetPath(options.seal, "C")
	if err != nil {
		return err
	}

	codeTemplate, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer codeTemplate.Close()

	scanner := bufio.NewScanner(codeTemplate)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case constants.CIPER:
			output("unsigned char cipertext[] = {", flag, file)
			outputAnomaly(options.cipertext, flag, file)
		case constants.KEY:
			output("unsigned char key[] = {", flag, file)
			outputAnomaly(options.key, flag, file)
		case constants.FILENAME:
			output("LPWSTR filename = L\""+options.exportName+"\";\n", flag, file)
		default:
			output(line+"\n", flag, file)
		}
	}
	return nil
}

func outputAnomaly(anomaly []byte, flag bool, file *os.File) {
	for i, ciperchar := range anomaly {
		if i%4 == 0 {
			output("\n\t", flag, file)
		}
		if i != len(anomaly)-1 {
			out := fmt.Sprintf("0x%x, ", ciperchar)
			output(out, flag, file)
		} else {
			out := fmt.Sprintf("0x%x\n};\n", ciperchar)
			output(out, flag, file)
		}
	}
}

func output(line string, flag bool, file *os.File) {
	if flag {
		fmt.Print(line)
	} else {
		file.Write([]byte(line))
	}
}
