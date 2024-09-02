package main

import (
	"fmt"
	"strings"

	io "github.com/DanielRasho/TC-1-ShuntingYard/internal/io"
	shuttingyard "github.com/DanielRasho/TC-1-ShuntingYard/internal/shuntingyard"
)

func main() {

	fileReader, err := io.ReadFile("input_data/shuntingYardInput.txt")
	if err != nil {
		panic("Input file do not exist")
	}
	defer fileReader.Close()

	var line string
	i := 0
	for fileReader.NextLine(&line) {
		fmt.Printf("\nLINE %d\n", i)
		postFix, _, _ := shuttingyard.RegexToPostfix(strings.Trim(line, "\n"), false)
		fmt.Printf("RESPONSE: %s\n", postFix)
		i++
	}
}
