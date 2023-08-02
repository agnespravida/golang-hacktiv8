package main

import (
	"fmt"
	"strings"
)

func main() {
	var sentence string = "Lalala"
	var alphabets = map[string]int{}

	for i := 0; i < len(sentence); i++ {
		var alphabet string = string(sentence[i])
		fmt.Println(alphabet)
		var _, isExist = alphabets[strings.ToLower(alphabet)]

		if isExist {
			alphabets[strings.ToLower(alphabet)]++
		} else {
			alphabets[strings.ToLower(alphabet)] = 1
		}
	}

	fmt.Println(alphabets)
}
