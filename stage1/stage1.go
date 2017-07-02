package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

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
		table[k] = v / float32(len(input))
	}

	// Display Frequency Table
	b, err := json.MarshalIndent(table, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}

func main() {
	raw, _ := ioutil.ReadFile("cipher.text")
	cipher := strings.Trim(string(raw), "\n")
	fmt.Println(string(cipher))
	freqTable(cipher)
}
