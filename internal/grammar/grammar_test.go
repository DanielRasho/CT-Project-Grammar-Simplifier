package grammar

import (
	"testing"
)

var grammarFinalTest = Grammar{
	"S":   {"NP VP"},
	"VP":  {"VP PP", "V NP", "cooks", "drinks", "eats", "cuts"},
	"PP":  {"P NP"},
	"NP":  {"Det N", "he", "she"},
	"V":   {"cooks", "drinks", "eats", "cuts"},
	"P":   {"in", "with"},
	"N":   {"cat", "dog", "beer", "cake", "juice", "meat", "soup", "fork", "knife", "oven", "spoon"},
	"Det": {"a", "the"},
}

// Grammar y resultado esperado para probar la obtención de terminales y no terminales
var grammarTerminalesTest = Grammar{
	"A": {"ab", "bb", "Ba"},
	"B": {"ab", "bb", "Ba"},
	"C": {"ab", "bb", Epsilon},
}

var expectedNonTerminals = map[string]struct{}{
	"A": {},
	"B": {},
	"C": {},
}

var expectedTerminals = map[string]struct{}{
	"a": {},
	"b": {},
	"ε": {},
}

// Test para la función getNonTerminals usando TestGrammar1.
func TestGetNonTerminals(t *testing.T) {

	nonTerminals := getNonTerminals(&grammarTerminalesTest)

	if len(nonTerminals) != len(expectedNonTerminals) {
		t.Errorf("Se esperaba %v, pero se obtuvo %v", expectedNonTerminals, nonTerminals)
	}

	for nt := range expectedNonTerminals {
		if _, exists := nonTerminals[nt]; !exists {
			t.Errorf("El no terminal %s no se encontró en el resultado", nt)
		}
	}
}

// Test para la función getTerminals usando TestGrammar1.
func TestGetTerminals(t *testing.T) {
	nonTerminals := getNonTerminals(&grammarTerminalesTest)

	terminals := getTerminals(&grammarTerminalesTest, nonTerminals)

	if len(terminals) != len(expectedTerminals) {
		t.Errorf("Se esperaba %v, pero se obtuvo %v", expectedTerminals, terminals)
	}

	for st := range expectedTerminals {
		if _, exists := terminals[st]; !exists {
			t.Errorf("El terminal %s no se encontró en el resultado", st)
		}
	}
}
