package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ast "github.com/DanielRasho/Computation-Theory/internal/abstract_syntax_tree"
	grammar "github.com/DanielRasho/Computation-Theory/internal/grammar"
	nfaAutomata "github.com/DanielRasho/Computation-Theory/internal/nfa"
	runner "github.com/DanielRasho/Computation-Theory/internal/runner_simulation"
	shuttingyard "github.com/DanielRasho/Computation-Theory/internal/shuntingyard"
)

// REGEX DEFINITIONS
const LETTERS string = "[abcdefghijklmnopqrstuvwxyz]"
const CAPITAL_LETTERS string = "[ABCDEFGHIJKLMNOPQRSTUVWXYZ]"
const DIGITS string = "[0123456789]"

// PRODUCTIONS_REGEX for matching grammar productions
var PRODUCTIONS_REGEX = fmt.Sprintf("%s -> ((%s|%s|%s|ε)+\\|)*(%s|%s|%s|ε)+",
	CAPITAL_LETTERS,
	LETTERS, CAPITAL_LETTERS, DIGITS,
	LETTERS, CAPITAL_LETTERS, DIGITS)

func main() {
	filepath := "./input_data/grammars.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("No se pudo abrir el archivo: %s.\n%v", filepath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Creamos un NFA para validar las producciones
	nfa := NFA_initializer()

	var currentGrammar grammar.Grammar = make(grammar.Grammar)
	grammarCounter := 1
	fmt.Println("\n=================================")
	fmt.Printf("📝 Procesando gramática %d:\n", grammarCounter)
	fmt.Println("=================================")
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Detecta el delimitador que separa las gramáticas
		if line == "---" {
			grammar.SimplifyGrammar(&currentGrammar, true)

			// Preparar para la siguiente gramática
			grammarCounter++
			fmt.Println("\n=================================")
			fmt.Printf("📝 Procesando gramática %d:\n", grammarCounter)
			fmt.Println("=================================")
			currentGrammar = make(grammar.Grammar)
			continue
		}

		// Ignorar líneas vacías o comentarios
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		fmt.Printf("Checking %s :", line)
		conclusion := runner.RunnerNFA(nfa, line)
		if conclusion {
			fmt.Printf(" is ✅\n")
			currentGrammar.AddProduction(line)
		} else {
			fmt.Printf(" is ❌\n ERROR: incorrect grammar\n")
			return
		}
	}
	grammar.SimplifyGrammar(&currentGrammar, true)
}

// Creates the NFA for checking if a production is valid
func NFA_initializer() *nfaAutomata.NFA {
	_, postfix, _ := shuttingyard.RegexToPostfix(PRODUCTIONS_REGEX, false)
	root := ast.BuildAST(postfix)
	nfa := nfaAutomata.BuildNFA(root)
	return nfa
}
