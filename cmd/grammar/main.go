package main

import (
	"fmt"

	io "github.com/DanielRasho/Computation-Theory/internal/IO"
	ast "github.com/DanielRasho/Computation-Theory/internal/abstract_syntax_tree"
	nfaAutomata "github.com/DanielRasho/Computation-Theory/internal/nfa"
	runner "github.com/DanielRasho/Computation-Theory/internal/runner_simulation"
	shuttingyard "github.com/DanielRasho/Computation-Theory/internal/shuntingyard"
)

// REGEX DEFINITIONS
const LETTERS string = "[abcdefghijklmnopqrstuvwxyz]"
const CAPITAL_LETTERS string = "[ABCDEFGHIJKLMNOPQRSTUVWXYZ]"
const DIGITS string = "[0123456789]"

// A -> ab|ε
var PRODUCTIONS_REGEX = fmt.Sprintf("%s -> ((%s|%s|%s|ε)+\\#)*(%s|%s|%s|ε)+",
	CAPITAL_LETTERS,
	LETTERS, CAPITAL_LETTERS, DIGITS,
	LETTERS, CAPITAL_LETTERS, DIGITS)

// MAIN PROGRAM
func main() {

	nfa := NFA_initializer()

	filepath := ""

	fileReader, err := io.ReadFile("./input_data/test.txt")

	if err != nil {
		fmt.Printf("No se puedo abrir el archivo: %s.\n %v", filepath, err)
		return
	}
	defer fileReader.Close()

	var line string
	for fileReader.NextLine(&line) {
		fmt.Printf("Checking %s :", line)
		conclusion := runner.RunnerNFA(nfa, line)
		if conclusion {
			fmt.Printf(" is ✅\n")
		} else {
			fmt.Printf(" is ❌\n ERROR: incorrect grammar\n")
			return
		}
	}

	// TODO:
	// - Test above code
	// - Create a Production Type and parse the string to this type
	// - Simplify the grammar

}

// Creates the NFA for check if a production is valid
func NFA_initializer() *nfaAutomata.NFA {
	postfix, _, _ := shuttingyard.RegexToPostfix(PRODUCTIONS_REGEX, false)
	root := ast.BuildAST(postfix)
	nfa := nfaAutomata.BuildNFA(root)

	return nfa
}
