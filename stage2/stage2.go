package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	raw, _ := ioutil.ReadFile("cipher.text")
	cipher := strings.Trim(string(raw), "\n")
	fmt.Println(string(cipher))
}
