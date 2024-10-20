package grammar

/*
import (
	"reflect"
	"testing"
)

// Datos de prueba
var grammarTest1 = Grammar{
	"1":   {"2 3"},
	"2":   {"4 5", "a"},
	"3":   {"b"},
	"4":   {"c"},
	"5":   {"6", "d"},
	"6":   {"7"},
	"7":   {"e"},
	"100": {"f"},
}

var grammarTest2 = Grammar{
	"!": {"@ #"},
	"@": {"a"},
	"#": {"b", "&"},
	"&": {"c"},
	"%": {"d"},
}

var grammarTest3 = Grammar{
	"Start":      {"Noun Verb Object"},
	"Noun":       {"Dog", "Cat"},
	"Verb":       {"chases", "eats"},
	"Object":     {"Ball", "Food"},
	"UnusedRule": {"Tree"},
}

var grammarTest4 = Grammar{
	"S": {"A B"},
	"A": {"a"},
	"B": {"b"},
	"C": {"c"},
}

// Test para buscar simbolos generadores
func TestFindGeneratingSymbols(t *testing.T) {
	tests := []struct {
		name           string
		grammar        *Grammar
		expectedGen    []string
		expectedNonGen []string
	}{
		{
			name:           "GRAMATICA1",
			grammar:        &grammarTest1,
			expectedGen:    []string{"1", "2", "3", "4", "5", "6", "7", "100"},
			expectedNonGen: []string{},
		},
		{
			name:           "GRAMATICA2",
			grammar:        &grammarTest2,
			expectedGen:    []string{"@", "#", "&", "%", "!"},
			expectedNonGen: []string{},
		},
		{
			name:           "GRAMATICA3",
			grammar:        &grammarTest3,
			expectedGen:    []string{"Start", "Noun", "Verb", "Object", "UnusedRule"},
			expectedNonGen: []string{},
		},
		{
			name:           "GRAMATICA4",
			grammar:        &grammarTest4,
			expectedGen:    []string{"S", "A", "B", "C"},
			expectedNonGen: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generators, nonGenerators := findGeneratingSymbols(tt.grammar)

			// Usamos la nueva función sameElements para comparar sin importar el orden.
			if !sameElements(getKeys(generators), tt.expectedGen) {
				t.Errorf("Símbolos generadores incorrectos para %s: got %v, want %v", tt.name, getKeys(generators), tt.expectedGen)
			}
			if !sameElements(getKeys(nonGenerators), tt.expectedNonGen) {
				t.Errorf("Símbolos no generadores incorrectos para %s: got %v, want %v", tt.name, getKeys(nonGenerators), tt.expectedNonGen)
			}
		})
	}
}

// Test para buscar simbolos alcanzables
func TestFindReachableSymbols(t *testing.T) {
	tests := []struct {
		name             string
		grammar          *Grammar
		startSymbol      string
		expectedReach    []string
		expectedNonReach []string
	}{
		{
			name:             "GRAMATICA1",
			grammar:          &grammarTest1,
			startSymbol:      "1",
			expectedReach:    []string{"1", "2", "3", "4", "5", "6", "7"},
			expectedNonReach: []string{"100"},
		},
		{
			name:             "GRAMATICA2",
			grammar:          &grammarTest2,
			startSymbol:      "!",
			expectedReach:    []string{"!", "@", "#", "&"},
			expectedNonReach: []string{"%"},
		},
		{
			name:             "GRAMATICA3",
			grammar:          &grammarTest3,
			startSymbol:      "Start",
			expectedReach:    []string{"Start", "Noun", "Verb", "Object"},
			expectedNonReach: []string{"UnusedRule"},
		},
		{
			name:             "GRAMATICA4",
			grammar:          &grammarTest4,
			startSymbol:      "S",
			expectedReach:    []string{"S", "A", "B"},
			expectedNonReach: []string{"C"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reachable, nonReachable := findReachableSymbols(tt.grammar, tt.startSymbol)

			// Usamos la función sameElements para comparar sin importar el orden.
			if !sameElements(getKeys(reachable), tt.expectedReach) {
				t.Errorf("Símbolos alcanzables incorrectos para %s: got %v, want %v", tt.name, getKeys(reachable), tt.expectedReach)
			}
			if !sameElements(getKeys(nonReachable), tt.expectedNonReach) {
				t.Errorf("Símbolos no alcanzables incorrectos para %s: got %v, want %v", tt.name, getKeys(nonReachable), tt.expectedNonReach)
			}
		})
	}
}

// Función auxiliar para comparar dos slices ignorando el orden.
func sameElements(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// Convertimos ambos slices en mapas para comparar sus contenidos sin importar el orden.
	mapA := make(map[string]struct{}, len(a))
	mapB := make(map[string]struct{}, len(b))

	for _, v := range a {
		mapA[v] = struct{}{}
	}
	for _, v := range b {
		mapB[v] = struct{}{}
	}

	return reflect.DeepEqual(mapA, mapB)
}
*/
