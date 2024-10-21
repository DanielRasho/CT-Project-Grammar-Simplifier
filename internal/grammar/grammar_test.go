package grammar

import (
	"fmt"
	"testing"
)

func TestAddProductionFromString(t *testing.T) {
	// Create a new Grammar instance
	g := Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}

	// Add some productions
	g.AddProductionFromString("A -> a|{B}C|{B}C")
	g.AddProductionFromString("B -> b|{C}D")

	// Check the productions after adding
	expectedGrammar := `NonTerminals: [{A_0},{B_0},{C_0}]
Terminals: [a,C,b,D]

{A_0} -> a|{B_0}C
{B_0} -> b|{C_0}D
`
	if g.String(true) != expectedGrammar {
		t.Errorf("Expected %q,\n but got %q", expectedGrammar, g.String(true))
	}

	// fmt.Println(g.String(true))
}

func TestAddDuplicateProductionFromString(t *testing.T) {

	// Create a new Grammar instance
	g := Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}

	// Add some productions
	g.AddProductionFromString("A -> a|{B}C|{B}C")
	g.AddProductionFromString("B -> b|{C}D")
	g.AddProductionFromString("B -> {C}D")

	expectedGrammar := `NonTerminals: [{A_0},{B_0},{C_0}]
Terminals: [a,C,b,D]

{A_0} -> a|{B_0}C
{B_0} -> b|{C_0}D
`
	if g.String(true) != expectedGrammar {
		t.Errorf("Expected %q\n, but got %q", expectedGrammar, g.String(true))
	}

	// fmt.Println(g.String(true))

}
func TestAddNonTerminalOnBodyFromString(t *testing.T) {

	// Create a new Grammar instance
	g := Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}

	// Add some productions
	g.AddProductionFromString("A -> a|{B}C|{B}C")
	g.AddProductionFromString("B -> b|{C}D")
	g.AddProductionFromString("B -> {M}")

	expectedGrammar := `NonTerminals: [{A_0},{B_0},{C_0},{M_0}]
Terminals: [a,C,b,D]

{A_0} -> a|{B_0}C
{B_0} -> b|{C_0}D|{M_0}
`
	if g.String(true) != expectedGrammar {
		t.Errorf("Expected %q\n, but got %q", expectedGrammar, g.String(true))
	}

	// fmt.Println(g.String(true))

}

func TestAddProduction(t *testing.T) {

	// Create a new Grammar instance
	g := Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}
	g.AddProductionFromString("A -> a|b|{B}C")
	g.AddProductionFromString("B -> b|{C}D")

	// Test 1: Add a production for head "A" with terminal body
	bodies1 := [][]Symbol{
		{{IsTerminal: true, Value: "a", Id: 0}},
	}
	head1 := g.AddProduction("A", bodies1)

	fmt.Println(g.String(true))
	fmt.Println(head1.String())
}

func TestAddProductionSymbol(t *testing.T) {

	// Create a new Grammar instance
	g := Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}
	g.AddProductionFromString("A -> a|b|{B}C")
	g.AddProductionFromString("B -> b|{C}D")

	// Test 1: Add a production for head "A" with terminal body
	bodies1 := [][]Symbol{
		{{IsTerminal: true, Value: "j", Id: 0}},
		{{IsTerminal: true, Value: "b", Id: 0}},
	}
	head1 := Symbol{Value: "A", IsTerminal: false}

	g.AddProductionBodies(head1, bodies1)

	fmt.Println(g.String(true))
	fmt.Println(head1.String())
}

func TestRemoveDuplicates(t *testing.T) {
	slice := []Symbol{
		{IsTerminal: true, Value: "a", Id: 0},
		{IsTerminal: true, Value: "a", Id: 0},
		{IsTerminal: false, Value: "B", Id: 0},
	}

	uniqueSlice := removeDuplicatesSymbols(slice)

	if len(uniqueSlice) != 2 {
		t.Errorf("Expected 2 unique symbols, but got %d", len(uniqueSlice))
	}

	expected := []Symbol{
		{IsTerminal: true, Value: "a", Id: 0},
		{IsTerminal: false, Value: "B", Id: 0},
	}

	for i, sym := range expected {
		if uniqueSlice[i] != sym {
			t.Errorf("Expected symbol %v at index %d, but got %v", sym, i, uniqueSlice[i])
		}
	}
}
