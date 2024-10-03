package grammar

import (
	"testing"
)

// Datos de prueba
var grammarTest = Grammar{
	"A": {"A", "B", "D"},
	"B": {"C"},
	"C": {"abc", "D"},
	"D": {"C", "E"},
	"E": {"fgh"},
}

var nonTerminalsTest = map[string]struct{}{
	"A": {},
	"B": {},
	"C": {},
	"D": {},
	"E": {},
}

var expectedUnaryPairs = map[string][]string{
	"A": {"A", "B", "D"},
	"B": {"B", "C"},
	"C": {"C", "D"},
	"D": {"D", "C", "E"},
	"E": {"E"},
}

var expectedFindUnaryPairs = map[string][]string{
	"A": {"A", "B", "D", "C", "E"},
	"B": {"B", "C", "D", "E"},
	"C": {"C", "D", "E"},
	"D": {"D", "C", "E"},
	"E": {"E"},
}

var expectedRemoveUnaryProductions = Grammar{
	"A": {"abc", "fgh"},
	"B": {"abc", "fgh"},
	"C": {"abc", "fgh"},
	"D": {"abc", "fgh"},
	"E": {"fgh"},
}

func TestUnaryPairs(t *testing.T) {
	result := initializeUnaryPairs(&grammarTest)

	// Comprobar si los resultados son iguales
	for key, expectedValues := range expectedUnaryPairs {
		resultValues, exists := result[key]
		if !exists {
			t.Errorf("Se esperaba que %s existiera en los resultados", key)
			continue
		}
		if !equalSlices(resultValues, expectedValues) {
			t.Errorf("Para %s, se esperaba %v, pero se obtuvo %v", key, expectedValues, resultValues)
		}
	}
}

func TestFindUnaryPairs(t *testing.T) {
	result := findUnaryPairs(expectedUnaryPairs)

	// Comprobar si los resultados son iguales
	for key, expectedValues := range expectedFindUnaryPairs {
		resultValues, exists := result[key]
		if !exists {
			t.Errorf("Se esperaba que %s existiera en los resultados", key)
			continue
		}
		if !equalSlices(resultValues, expectedValues) {
			t.Errorf("Para %s, se esperaba %v, pero se obtuvo %v", key, expectedValues, resultValues)
		}
	}
}

func TestRemoveUnaryProductions(t *testing.T) {
	result := removeUnaryProductions(&grammarTest, expectedFindUnaryPairs, nonTerminalsTest)

	// Comprobar si los resultados son iguales
	for key, expectedValues := range expectedRemoveUnaryProductions {
		resultValues, exists := (*result)[key]
		if !exists {
			t.Errorf("Se esperaba que %s existiera en los resultados", key)
			continue
		}
		if !equalSlices(resultValues, expectedValues) {
			t.Errorf("Para %s, se esperaba %v, pero se obtuvo %v", key, expectedValues, resultValues)
		}
	}
}

// Función para comparar dos slices de strings
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
