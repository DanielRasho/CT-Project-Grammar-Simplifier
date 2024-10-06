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
	"C": {"ABA", Epsilon + "BA", "AB" + Epsilon, "A" + Epsilon + "A", Epsilon + "B" + Epsilon, Epsilon + Epsilon + "A", "A" + Epsilon + Epsilon, Epsilon + Epsilon + Epsilon},
}

// Grammar y  resultado esperado par TestRemoveEpsilons
var grammarTestRemoveEpsilons = Grammar{
	"A": {"a", Epsilon},
	"B": {"b", Epsilon},
	"C": {"ABA", Epsilon + "BA", "AB" + Epsilon, "A" + Epsilon + "A", Epsilon + "B" + Epsilon, Epsilon + Epsilon + "A", "A" + Epsilon + Epsilon, Epsilon + Epsilon + Epsilon},
}

var expectedGrammarTestRemoveEpsilons = Grammar{
	"A": {"a"},
	"B": {"b"},
	"C": {"ABA", "BA", "AB", "AA", "B", "A"},
}

// Grammar y  resultado esperado par TestRemoveDuplicates
var grammarTestRemoveDuplicates = Grammar{
	"A": {"a"},
	"B": {"b"},
	"C": {"ABA", "BA", "AB", "AA", "B", "A", "A"},
}

var expectedGrammarTestRemoveDuplicates = Grammar{
	"A": {"a"},
	"B": {"b"},
	"C": {"ABA", "BA", "AB", "AA", "B", "A"},
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

// Test para verificar que elimina epsilon de producciones
func TestRemoveEpsilons(t *testing.T) {
	// Ejecutar la función RemoveEpsilons
	result := RemoveEpsilons(&grammarTestRemoveEpsilons)

	// Comparar el resultado con la gramática esperada
	if !reflect.DeepEqual(result, &expectedGrammarTestRemoveEpsilons) {
		t.Errorf("Test failed. Expected: %v, Got: %v", expectedGrammarTestRemoveEpsilons, result)
	}
}

// Test para verificar que eliminar producciones duplicadas
func TestRemoveDuplicates(t *testing.T) {
	// Crear una copia de la gramática para aplicar la función removeDuplicates
	resultGrammar := make(Grammar)
	for head, productions := range grammarTestRemoveDuplicates {
		resultGrammar[head] = RemoveDuplicates(productions)
	}

	// Comparar el resultado con la gramática esperada
	if !reflect.DeepEqual(resultGrammar, expectedGrammarTestRemoveDuplicates) {
		t.Errorf("Test failed. Expected: %v, Got: %v", expectedGrammarTestRemoveDuplicates, resultGrammar)
	}
}
