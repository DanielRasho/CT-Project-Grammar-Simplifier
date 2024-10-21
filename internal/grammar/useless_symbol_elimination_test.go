package grammar

import "testing"

// Simbolos para la gramática
var S1 = Symbol{isTerminal: false, value: "S", id: 1}
var A1 = Symbol{isTerminal: false, value: "A", id: 1}
var B1 = Symbol{isTerminal: false, value: "B", id: 1}
var a1 = Symbol{isTerminal: true, value: "a", id: -1}
var b1 = Symbol{isTerminal: true, value: "b", id: -1}

// Grammar Test
var grammarTestUselessSymbolElimination = &Grammar{
	terminals:    []Symbol{a1, b1},
	nonTerminals: []Symbol{S1, A1, B1},
	Productions: map[Symbol][][]Symbol{
		S1: {{A1, B1}, {a1}},
		A1: {{b1}},
	},
}

// Resultado de eliminar simbolos no alcanzables
var expectedRemoveReachableSymbols = &Grammar{
	terminals:    []Symbol{a1, b1},
	nonTerminals: []Symbol{S1, A1, B1},
	Productions: map[Symbol][][]Symbol{
		S1: {{A1, B1}, {a1}},
		A1: {{b1}},
	},
}

// Resultado de eliminar simbolos que no generan
var expectedRemoveGeneratingSymbols = &Grammar{
	terminals:    []Symbol{a1, b1},
	nonTerminals: []Symbol{S1, A1},
	Productions: map[Symbol][][]Symbol{
		S1: {{a1}},
		A1: {{b1}},
	},
}

// Resultado esperado en de Grammar sin simbolos inutiles
var expectedTestUselessSymbolElimination = &Grammar{
	terminals:    []Symbol{a1},
	nonTerminals: []Symbol{S1},
	Productions: map[Symbol][][]Symbol{
		S1: {{a1}},
	},
}

var expectedGeneratingSymbols1 = []Symbol{S1, A1}    // generan
var expectedGeneratingSymbols2 = []Symbol{B1}        // no generan
var expectedReachableSymbols1 = []Symbol{S1, A1, B1} // alcanzables
var expectedReachableSymbols2 = []Symbol{}           // no alcanzables

// Test para la función findGeneratingSymbols
func TestFindGeneratingSymbols(t *testing.T) {
	// Ejecutar la función para encontrar los símbolos que generan cadenas de terminales
	generatingSymbols, nonGeneratingSymbols := findGeneratingSymbols(grammarTestUselessSymbolElimination)

	// Verificar que los símbolos generadores coincidan con los esperados
	for _, expectedSymbol := range expectedGeneratingSymbols1 {
		if !containsSymbol(generatingSymbols, expectedSymbol) {
			t.Errorf("Error: El símbolo generador esperado %v no se encontró en los símbolos generadores: %v", expectedSymbol.String(), generatingSymbols)
		}
	}

	// Verificar que los símbolos no generadores coincidan con los esperados
	for _, expectedSymbol := range expectedGeneratingSymbols2 {
		if !containsSymbol(nonGeneratingSymbols, expectedSymbol) {
			t.Errorf("Error: El símbolo no generador esperado %v no se encontró en los símbolos no generadores: %v", expectedSymbol.String(), nonGeneratingSymbols)
		}
	}
}

// Test para la función findReachableSymbols
func TestFindReachableSymbols(t *testing.T) {
	// Ejecutar la función para encontrar los símbolos alcanzables a partir del símbolo inicial S1
	reachableSymbols, unreachableSymbols := findReachableSymbols(grammarTestUselessSymbolElimination, S1)

	// Verificar que los símbolos alcanzables coincidan con los esperados
	for _, expectedSymbol := range expectedReachableSymbols1 {
		if !containsSymbol(reachableSymbols, expectedSymbol) {
			t.Errorf("Error: El símbolo alcanzable esperado %v no se encontró en los símbolos alcanzables: %v", expectedSymbol.String(), reachableSymbols)
		}
	}

	// Verificar que los símbolos no alcanzables coincidan con los esperados
	for _, expectedSymbol := range expectedReachableSymbols2 {
		if !containsSymbol(unreachableSymbols, expectedSymbol) {
			t.Errorf("Error: El símbolo no alcanzable esperado %v no se encontró en los símbolos no alcanzables: %v", expectedSymbol.String(), unreachableSymbols)
		}
	}
}

// Test para la función RemoveNonGeneratingSymbols
func TestRemoveNonGeneratingSymbols(t *testing.T) {
	// Ejecutar la función para eliminar los símbolos no generadores
	result := RemoveNonGeneratingSymbols(grammarTestUselessSymbolElimination)

	// Comparar la gramática resultante con la esperada
	if !compareGrammars(result, expectedRemoveGeneratingSymbols) {
		t.Errorf("Error: La gramática resultante de RemoveNonGeneratingSymbols no coincide con la esperada.\nEsperado: %v\nObtenido: %v", expectedRemoveGeneratingSymbols.String(true), result.String(true))
	}
}

// Test para la función RemoveNonReachableSymbols
func TestRemoveNonReachableSymbols(t *testing.T) {
	// Ejecutar la función para eliminar los símbolos no alcanzables
	result := RemoveNonReachableSymbols(grammarTestUselessSymbolElimination, S1)

	// Comparar la gramática resultante con la esperada
	if !compareGrammars(result, expectedRemoveReachableSymbols) {
		t.Errorf("Error: La gramática resultante de RemoveNonReachableSymbols no coincide con la esperada.\nEsperado: %v\nObtenido: %v", expectedRemoveReachableSymbols.String(true), result.String(true))
	}
}

// Test para la función RemoveUselessSymbols
func TestRemoveUselessSymbols(t *testing.T) {
	// Ejecutar la función para eliminar los símbolos inútiles (no generadores y no alcanzables)
	result := RemoveUselessSymbols(grammarTestUselessSymbolElimination, S1)

	// Comparar la gramática resultante con la esperada
	if !compareGrammars(result, expectedTestUselessSymbolElimination) {
		t.Errorf("Error: La gramática resultante de RemoveUselessSymbols no coincide con la esperada.\nEsperado: %v\nObtenido: %v", expectedTestUselessSymbolElimination.String(true), result.String(true))
	}
}
