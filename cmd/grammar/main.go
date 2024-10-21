package main

import (
	"fmt"
	"strings"
	"time"

	io "github.com/DanielRasho/Computation-Theory/internal/IO"
	ast "github.com/DanielRasho/Computation-Theory/internal/abstract_syntax_tree"
	"github.com/DanielRasho/Computation-Theory/internal/grammar"
	nfaAutomata "github.com/DanielRasho/Computation-Theory/internal/nfa"
	runner "github.com/DanielRasho/Computation-Theory/internal/runner_simulation"
	shuttingyard "github.com/DanielRasho/Computation-Theory/internal/shuntingyard"
)

// REGEX DEFINITIONS
const LETTERS string = "[abcdefghijklmnopqrstuvwxyz ]"
const CAPITAL_LETTERS string = "[ABCDEFGHIJKLMNOPQRSTUVWXYZ]"
const OPERATORS string = "[\\+\\*\\(\\)]"
const DIGITS string = "[0123456789]"
const NON_TERMINALS string = "\\{([ABCDEFGHIJKLMNOPQRSTUVWXYZ])+\\}"

// PRODUCTIONS_REGEX for matching grammar productions
var PRODUCTIONS_REGEX = fmt.Sprintf("(%s)+ -> ((%s|%s|%s|%s|%s|ε)+\\|)*(%s|%s|%s|%s|%s|ε)+",
	CAPITAL_LETTERS,
	OPERATORS, NON_TERMINALS, LETTERS, CAPITAL_LETTERS, DIGITS,
	OPERATORS, NON_TERMINALS, LETTERS, CAPITAL_LETTERS, DIGITS)

func main() {
	filepath := "./input_data/grammars.txt"

	fileReader, err := io.ReadFile(filepath)

	if err != nil {
		fmt.Printf("No se puedo abrir el archivo: %s.\n %v", filepath, err)
		return
	}

	defer fileReader.Close()

	// Creamos un NFA para validar las producciones
	nfa := NFA_initializer()
	currentGrammar := grammar.Grammar{Productions: make(map[grammar.Symbol][][]grammar.Symbol)}
	grammarCounter := 1

	fmt.Println("\n=================================")
	fmt.Printf("📝 Procesando gramática %d:\n", grammarCounter)
	fmt.Println("=================================")

	var line string
	for fileReader.NextLine(&line) {

		// Detecta el delimitador que separa las gramáticas
		if line == "---" {
			grammar.SimplifyGrammar(&currentGrammar, true)

			grammarCounter++
			fmt.Println("\n=================================")
			fmt.Printf("📝 Procesando gramática %d:\n", grammarCounter)
			fmt.Println("=================================")

			// Preparar para la siguiente gramática
			currentGrammar = grammar.Grammar{Productions: make(map[grammar.Symbol][][]grammar.Symbol)}
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
			currentGrammar.AddProductionFromString(line)
		} else {
			fmt.Printf(" is ❌\n ERROR: incorrect grammar\n")
			return
		}
	}

	startSymbol := currentGrammar.NonTerminals[0]
	// Simplify Grammar

	// Capturar el tiempo de inicio
	start := time.Now()
	newGrammar := grammar.SimplifyGrammar(&currentGrammar, true)
	// Capturar el tiempo después de la simplificación
	elapsed := time.Since(start)
	fmt.Println(newGrammar.Productions[startSymbol])
	// Imprimir el tiempo que tomó la simplificación
	fmt.Printf("Tiempo de simplificación: %s\n", elapsed)

	// Get User Input
	var input string
	fmt.Print("🔰Ingresar valor para verificar: ")
	fmt.Scanln(&input)

	accepted := grammar.CYKParse(newGrammar, input, startSymbol)
	if accepted {
		fmt.Println("La cadena es aceptada por la gramática.")
	} else {
		fmt.Println("La cadena NO es aceptada por la gramática.")
	}
}

// Creates the NFA for checking if a production is valid
func NFA_initializer() *nfaAutomata.NFA {
	_, postfix, _ := shuttingyard.RegexToPostfix(PRODUCTIONS_REGEX, false)
	root := ast.BuildAST(postfix)
	nfa := nfaAutomata.BuildNFA(root)
	return nfa
}
