package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func caesar(input string, shift int) string {

	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	L := len(alphabet)

	output := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		if input[i] != 0x20 {
			pos := (strings.IndexByte(alphabet, input[i]) + shift) % L
			output[i] = alphabet[pos]
		} else {
			output[i] = 0x20
		}
	}
	return string(output)
}

func main() {
	raw, _ := ioutil.ReadFile("cipher.text")
	cipher := strings.Trim(string(raw), "\n")
	chars := []byte(cipher)
	fmt.Println("Original: ", string(chars))
	//cipher = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for shift := 0; shift <= 26; shift++ {
		fmt.Printf("Shift: %d\tResult: %s\n", shift, caesar(cipher, shift))
	}
}
