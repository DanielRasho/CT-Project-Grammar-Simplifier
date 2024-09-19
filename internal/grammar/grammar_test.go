package grammar

import (
	"reflect"
	"testing"
)

func areSlicesEqual(t *testing.T, response []string, expect []string) {
	value := ""
	for _, v := range response {
		value += v + " "
	}

	if len(response) < len(expect) {
		t.Fatalf("Response has less characters than expected. Has %d, %d given. %s", len(response), len(expect), value)
	} else if len(response) > len(expect) {
		t.Fatalf("Response has more characters than expected. Has %d, %d given. %s", len(response), len(expect), value)
	}
	for i, expected := range expect {
		if response[i] != expected {
			t.Fatalf("Characters not match, Given string %s", value)
		}
	}
}

func TestAddProductionToEmptyGrammar(t *testing.T) {
	grammar := make(Grammar)
	grammar.AddProduction("A -> a|a|bc|C")
	expectedBodyItems := []string{"a", "bc", "C"}

	if _, exist := grammar["A"]; !exist {
		t.Fatalf("A production was not appended to the grammar\n")
	}
	areSlicesEqual(t, grammar["A"], expectedBodyItems)
}
func TestAddProductionToNotEmptyGrammar(t *testing.T) {
	grammar := make(Grammar)
	grammar.AddProduction("A -> a|a|bc|C")
	grammar.AddProduction("A -> a|B|J")
	expectedBodyItems := []string{"a", "bc", "C", "B", "J"}

	if _, exist := grammar["A"]; !exist {
		t.Fatalf("A production was not appended to the grammar\n")
	}
	areSlicesEqual(t, grammar["A"], expectedBodyItems)
}

func TestIdentifyDirectNullables(t *testing.T) {
	grammar := make(Grammar)
	grammar.AddProduction("A -> ε|a|bc|C")
	grammar.AddProduction("B -> a|B|ε")
	grammar.AddProduction("C -> m")
	expectedNullables := []string{"A", "B"}
	response := identifyDirectNullables(&grammar)

	areSlicesEqual(t, *response, expectedNullables)
}

func TestIdentifyIndirectNullables(t *testing.T) {
	grammar := make(Grammar)
	grammar.AddProduction("A -> ε")
	grammar.AddProduction("B -> AA|ab")
	grammar.AddProduction("C -> Ab")
	directNullabes := identifyDirectNullables(&grammar)
	response := identifyIndirectNullables(&grammar, *directNullabes)
	expectedDirectNullables := []string{"A", "B"}

	areSlicesEqual(t, *response, expectedDirectNullables)
}

func TestIdentifyManyIndirectNullables(t *testing.T) {
	grammar := make(Grammar)
	grammar.AddProduction("A -> ε")
	grammar.AddProduction("D -> C")
	grammar.AddProduction("C -> B")
	grammar.AddProduction("B -> A")
	grammar.AddProduction("J -> d")
	directNullabes := identifyDirectNullables(&grammar)
	response := identifyIndirectNullables(&grammar, *directNullabes)
	expectedDirectNullables := []string{"A", "B", "C", "D"}

	areSlicesEqual(t, *response, expectedDirectNullables)
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

// Grammar, nullables y resultado esperado par TestReplaceNullables
var grammarTestReplace1 = Grammar{
	"A": {"BB", "a"},
	"B": {"b", Epsilon},
}
var expectedGrammarTestReplace1 = Grammar{
	"A": {"BB", Epsilon + "B", "B" + Epsilon, "a", Epsilon + Epsilon},
	"B": {"b", Epsilon},
}

var nullablesGrammarTest1 = []string{"B"}

// Grammar, nullables y resultado esperado par TestReplaceNullables
var grammarTestReplace2 = Grammar{
	"A": {"a", Epsilon},
	"B": {"b", Epsilon},
	"C": {"ABA"},
}

var nullablesGrammarTest2 = []string{"A", "B"}

var expectedGrammarTestReplace2 = Grammar{
	"A": {"a", Epsilon},
	"B": {"b", Epsilon},
	"C": {"ABA", Epsilon + "BA", "AB" + Epsilon, "AA" + Epsilon + "A", Epsilon + "B" + Epsilon, Epsilon + Epsilon + "A", "A" + Epsilon + Epsilon, Epsilon + Epsilon + Epsilon},
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

// Función de prueba para ReplaceNullables
func TestReplaceNullables(t *testing.T) {

	// Llamar a la función ReplaceNullables
	resultGrammar := ReplaceNullables(&grammarTestReplace1, nullablesGrammarTest1)

	// Comparar el resultado con el esperado
	if !reflect.DeepEqual(*resultGrammar, expectedGrammarTestReplace1) {
		t.Errorf("Resultado incorrecto.\nEsperado: %v\nObtenido: %v", expectedGrammarTestReplace1, *resultGrammar)
	}
}

// Función de prueba para ReplaceNullables
func TestReplaceNullables2(t *testing.T) {

	// Llamar a la función ReplaceNullables
	resultGrammar := ReplaceNullables(&grammarTestReplace2, nullablesGrammarTest2)

	// Comparar el resultado con el esperado
	if !reflect.DeepEqual(*resultGrammar, expectedGrammarTestReplace2) {
		t.Errorf("Resultado incorrecto.\nEsperado: %v\nObtenido: %v", expectedGrammarTestReplace2, *resultGrammar)
	}
}
