package grammar

import (
	"fmt"
	"testing"
)

// Helper function to create a Symbol easily
func createSymbol(value string, isTerminal bool, id int) Symbol {
	return Symbol{
		isTerminal: isTerminal,
		value:      value,
		id:         id,
	}
}

func TestFindAllBodyVariants(t *testing.T) {
	// Step 1: Define the symbols
	A := createSymbol("A", false, 0)

	// Step 2: Define the grammar
	grammar := &Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}
	grammar.AddProductionFromString("A -> {B}m")
	grammar.AddProductionFromString("B -> {C}|{D}o")
	grammar.AddProductionFromString("C -> {B}k")
	grammar.AddProductionFromString("D -> q")

	// fmt.Println(grammar.String(false))

	bodyVariants := findAllBodyVariants(&A, grammar)
	for _, variant := range bodyVariants {
		fmt.Println(variant)
	}
}

func TestFindAllBodyVariants2(t *testing.T) {
	// Step 1: Define the symbols
	A := createSymbol("A", false, 0)

	// Step 2: Define the grammar
	grammar := &Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}
	grammar.AddProductionFromString("A -> {B}{A}|{C}m|aa")
	grammar.AddProductionFromString("B -> {C}")
	grammar.AddProductionFromString("C -> {A}")

	// fmt.Println(grammar.String(false))

	bodyVariants := findAllBodyVariants(&A, grammar)
	for _, variant := range bodyVariants {
		fmt.Println(variant)
	}
}

func TestRemoveLeftRecursivity(t *testing.T) {
	// Step 2: Define the grammar
	grammar := &Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}
	grammar.AddProductionFromString("A -> {A}t{B}|{B}")
	grammar.AddProductionFromString("B -> int|l{A}l")

	newGrammar := removeLeftRecursivity(grammar)
	fmt.Println(newGrammar.String(false))
}
