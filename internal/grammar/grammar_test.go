package grammar

import (
	"fmt"
	"testing"
)

func TestAddProduction(t *testing.T) {
	// Create a new Grammar instance
	g := Grammar{
		productions: make(map[Symbol][][]Symbol),
	}

	// Add some productions
	g.AddProduction("A -> a|{B}C")
	g.AddProduction("B -> b|{C}D")

	fmt.Println(g.String())

	// Check the productions after adding
	expectedGrammar := "{A_0} -> a|{B_0}C\n{B_0} -> b|{C_0}D\n"
	if g.String() != expectedGrammar {
		t.Errorf("Expected %q, but got %q", expectedGrammar, g.String())
	}

	// Add a duplicate production
	g.AddProduction("A -> a|{B}C")
	expectedGrammar = "{A_0} -> a|{B_0}C\n{B_0} -> b|{C_0}D\n"
	if g.String() != expectedGrammar {
		t.Errorf("Expected %q, but got %q after adding duplicate production", expectedGrammar, g.String())
	}

	// Add another production with new symbols
	g.AddProduction("C -> c")
	expectedGrammar = "{A_0} -> a|{B_0}C\n{B_0} -> b|{C_0}D\n{C_0} -> c\n"
	if g.String() != expectedGrammar {
		t.Errorf("Expected %q, but got %q after adding new production", expectedGrammar, g.String())
	}

	// Check terminals and non-terminals
	if len(g.terminals) != 2 { // "a" and "b" from the first two productions
		t.Errorf("Expected 2 terminals, but got %d", len(g.terminals))
	}
	if len(g.nonTerminals) != 3 { // A, B, C
		t.Errorf("Expected 3 non-terminals, but got %d", len(g.nonTerminals))
	}
}

func TestRemoveDuplicates(t *testing.T) {
	slice := []Symbol{
		{isTerminal: true, value: "a", id: 0},
		{isTerminal: true, value: "a", id: 0},
		{isTerminal: false, value: "B", id: 0},
	}

	uniqueSlice := removeDuplicatesSymbols(slice)

	if len(uniqueSlice) != 2 {
		t.Errorf("Expected 2 unique symbols, but got %d", len(uniqueSlice))
	}

	expected := []Symbol{
		{isTerminal: true, value: "a", id: 0},
		{isTerminal: false, value: "B", id: 0},
	}

	for i, sym := range expected {
		if uniqueSlice[i] != sym {
			t.Errorf("Expected symbol %v at index %d, but got %v", sym, i, uniqueSlice[i])
		}
	}
}
