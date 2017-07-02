package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func prettyPrint(input interface{}) {
	// Display Frequency Table
	b, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}

func freqTable(input string) {
	table := map[string]float32{}
	for _, letter := range input {
		if letter != 0x20 {
			_, ok := table[string(letter)]
			if ok {
				table[string(letter)]++
			} else {
				table[string(letter)] = 1
			}
		}
	}
	// Divide by total length
	for k, v := range table {
		table[k] = 100 * (v / float32(len(input)))
	}

	prettyPrint(table)
}

func loadMapping() map[string]string {
	raw, _ := ioutil.ReadFile("mapping.json")

	var mapping map[string]string

	err := json.Unmarshal(raw, &mapping)
	if err != nil {
		fmt.Println("Error unmarshalling: ", err)
	}

	return mapping
}

func applyMapping(input string, mapping map[string]string) {
	output := make([]string, len(input))
	for pos, letter := range input {
		output[pos] = mapping[string(letter)]
		//fmt.Printf("%d: %s, %s\n", pos, letter, output[pos])
	}
	fmt.Printf("\nPlaintext:\n%s\n", strings.Join(output, ""))
}

func main() {
	raw, _ := ioutil.ReadFile("cipher.text")
	cipher := strings.Trim(string(raw), "\n")
	mapping := loadMapping()
	fmt.Printf("Cipher:\n%s\n", string(cipher))
	applyMapping(cipher, mapping)
}
