package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
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

func matches(a []rune, b []rune) int {
	count := 0
	for pos, letter := range a {
		if letter == b[pos] {
			count++
		}
	}
	return count
}

func superimpose(input string) {
	s := []rune(input)
	res := make(map[int]int, len(s)-1)
	for i := 1; i < len(s); i++ {
		displaced := make([]rune, len(s))
		pos := copy(displaced, s[i:])
		copy(displaced[pos:], s[0:i])
		m := matches(s, displaced)
		res[i] = m
	}

	keys := []int{}
	for k := range res {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("Shift: %d, Matches: %d\n", k, res[k])
	}

}

func main() {
	raw, _ := ioutil.ReadFile("cipher.text")
	cipher := strings.Trim(string(raw), "\n")
	superimpose(cipher)
}
