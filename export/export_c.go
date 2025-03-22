package export

import (
	"bufio"
	"fmt"
	"os"
	"sigillum/constants"
	"sigillum/utils"
)

func ExportC(key []byte, cipertext []byte, filepath string) bool {
	_, err := os.Create(filepath)
	if err != nil {
		return false
	}

	return parseC(key, cipertext, filepath)
}

func PrintC(key []byte, cipertext []byte) bool {
	return parseC(key, cipertext, "")
}

func parseC(key []byte, cipertext []byte, filepath string) bool {
	flag := filepath == ""

	filePath, err := utils.GetPath("rc4", "c")
	if err != nil {
		return false
	}

	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case constants.CIPER:
			output("unsigned char cipertext[] = {", flag, filepath)
			outputAnomaly(cipertext, flag, filepath)
		case constants.KEY:
			output("unsigned char key[] = {", flag, filepath)
			outputAnomaly(key, flag, filepath)
		default:
			output(line+"\n", flag, filepath)
		}
	}
	return true
}

func outputAnomaly(anomaly []byte, flag bool, filepath string) {
	for i, ciperchar := range anomaly {
		if i%4 == 0 {
			output("\n\t", flag, filepath)
		}
		if i != len(anomaly)-1 {
			out := fmt.Sprintf("0x%x, ", ciperchar)
			output(out, flag, filepath)
		} else {
			out := fmt.Sprintf("0x%x\n};\n", ciperchar)
			output(out, flag, filepath)
		}
	}
}

func output(line string, flag bool, filename string) {
	if flag {
		fmt.Print(line)
	} else {
		os.WriteFile(filename, []byte(line), 0774)
	}
}
