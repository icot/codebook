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

func superimpose(input string) int {
	s := []rune(input)
	res := make(map[int]int, len(s)-1)
	maxd := 0
	for i := 1; i < len(s); i++ {
		displaced := make([]rune, len(s))
		pos := copy(displaced, s[i:])
		copy(displaced[pos:], s[0:i])
		m := matches(s, displaced)
		res[i] = m
		if m >= maxd {
			maxd = m
		}
	}

	// Reverse map
	rmap := make(map[int][]int, maxd)

	for k, v := range res {
		if len(rmap[v]) >= 1 {
			rmap[v] = append(rmap[v], k)
		} else {
			rmap[v] = []int{k}
		}
	}

	keys := []int{}
	for k := range rmap {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("Matches: %d, Displacements: %+v\n", k, rmap[k])
	}

	displacement := len(input)
	for _, v := range rmap[keys[len(keys)-1]] {
		if v <= displacement {
			displacement = v
		}
	}

	return displacement + 1

}

func main() {
	raw, _ := ioutil.ReadFile("cipher.text")
	cipher := strings.Trim(string(raw), "\n")
	fmt.Println("Applying superimposition")
	keylength := superimpose(cipher)
	fmt.Printf("Tempative keylength found: %d\n", keylength)

	ciphers := make([][]byte, keylength)

	fmt.Printf("Total characters: %d\n", len(cipher))

	pos := 0
	for i := 0; i <= len(cipher)/keylength; i++ {
		for shift := 0; shift < keylength; shift++ {
			if len(ciphers[shift]) > 0 {
				ciphers[shift] = append(ciphers[shift], cipher[pos])
			} else {
				ciphers[shift] = []byte{cipher[pos]}
			}
			pos++
			if pos == len(cipher) {
				break
			}
		}
	}
	fmt.Printf("Total processed characters: %d\n", pos-1)
	fmt.Println("Sub-ciphers")
	for k, v := range ciphers {
		fmt.Printf("\nShift: %d, Lenght: %d, Cipher: \n%v\n\n", k, len(v), v)
	}
}
