package sealing

import "fmt"

func XORCreateSeal(key []byte, shellcode []byte) ([]byte, error) {
	fmt.Println("XOR")
	for i, j := 0, 0; i < len(shellcode); i, j = i+1, j+1 {
		if j >= len(key) {
			j = 0
		}
		shellcode[i] = shellcode[i] ^ key[j]
	}

	return shellcode, nil
}
