package grammar

import (
	"testing"
)

// Simbolos para la gramática
var A = Symbol{IsTerminal: false, Value: "A", Id: 1}
var B = Symbol{IsTerminal: false, Value: "B", Id: 1}
var C = Symbol{IsTerminal: false, Value: "C", Id: 1}
var D = Symbol{IsTerminal: false, Value: "D", Id: 1}
var E = Symbol{IsTerminal: false, Value: "E", Id: 1}
var abc = Symbol{IsTerminal: true, Value: "abc", Id: -1}
var fgh = Symbol{IsTerminal: true, Value: "fgh", Id: -1}

// Grammar Test
var grammarTestUnaryElimination = &Grammar{
	terminals:    []Symbol{abc, fgh},
	NonTerminals: []Symbol{A, B, C, D, E},
	Productions: map[Symbol][][]Symbol{
		A: {{A}, {B}, {D}},
		B: {{C}},
		C: {{abc}, {D}},
		D: {{C}, {E}},
		E: {{fgh}},
	},
}

// Resultado esperado en de Grammar sin producciones unarias
var expectedRemoveUnaryProductions = &Grammar{
	terminals:    []Symbol{abc, fgh},
	NonTerminals: []Symbol{A, B, C, D, E},
	Productions: map[Symbol][][]Symbol{
		A: {{abc}, {fgh}},
		B: {{abc}, {fgh}},
		C: {{abc}, {fgh}},
		D: {{abc}, {fgh}},
		E: {{fgh}},
	},
}

// No terminales de Grammar
var nonTerminalsTest = []Symbol{A, B, C, D, E}

// Resultado esperado para encontrar Pares Unarios Base
var expectedinitializeUnaryPairs = map[Symbol][]Symbol{
	A: {A, B, D},
	B: {B, C},
	C: {C, D},
	D: {D, C, E},
	E: {E},
}

// Resultado esperado para encontrar Pares Unarios
var expectedFindUnaryPairs = map[Symbol][]Symbol{
	A: {A, B, D, C, E},
	B: {B, C, D, E},
	C: {C, D, E},
	D: {D, C, E},
	E: {E},
}

// TestInitializeUnaryPairs verifica la correcta inicialización de pares unarios.
func TestInitializeUnaryPairs(t *testing.T) {
	// Se inicializan las parejas unarias usando la gramática de prueba
	unaryPairs := InitializeUnaryPairs(grammarTestUnaryElimination)

	// Verificar que los pares unarios coinciden con los esperados
	for nonTerminal, expectedValues := range expectedinitializeUnaryPairs {
		resultValues, exists := unaryPairs[nonTerminal]
		if !exists {
			t.Errorf("Error: No se encontraron pares unarios para %v", nonTerminal)
			continue
		}

		// Verificar que el número de valores sea correcto
		if len(resultValues) != len(expectedValues) {
			t.Errorf("Error: El número de pares unarios para %v no es correcto. Se esperaba %v, pero se obtuvo %v", nonTerminal, len(expectedValues), len(resultValues))
		}

		// Verificar que cada valor esperado esté presente en los resultados
		for _, expectedValue := range expectedValues {
			if !containsSymbol(resultValues, expectedValue) {
				t.Errorf("Error: El valor %v no se encontró en los pares unarios de %v. Pares actuales: %v", expectedValue, nonTerminal, resultValues)
			}
		}
	}
}

// TestFindUnaryPairs verifica la correcta expansión de las parejas unarias.
func TestFindUnaryPairs(t *testing.T) {
	// Se inicializan los pares unarios
	unaryBase := InitializeUnaryPairs(grammarTestUnaryElimination)

	// Se expanden las parejas unarias utilizando findUnaryPairs
	unaryPairs := FindUnaryPairs(unaryBase)

	// Verificar que los pares unarios coinciden con los esperados
	for nonTerminal, expectedValues := range expectedFindUnaryPairs {
		resultValues, exists := unaryPairs[nonTerminal]
		if !exists {
			t.Errorf("Error: No se encontraron pares unarios para %v", nonTerminal)
			continue
		}

		// Verificar que el número de valores sea correcto
		if len(resultValues) != len(expectedValues) {
			t.Errorf("Error: El número de pares unarios para %v no es correcto. Se esperaba %v, pero se obtuvo %v", nonTerminal, len(expectedValues), len(resultValues))
		}

		// Verificar que cada valor esperado esté presente en los resultados
		for _, expectedValue := range expectedValues {
			if !containsSymbol(resultValues, expectedValue) {
				t.Errorf("Error: El valor %v no se encontró en los pares unarios de %v. Pares actuales: %v", expectedValue, nonTerminal, resultValues)
			}
		}
	}
}

// Test for removeUnaryProductions
func TestRemoveUnaryProductions(t *testing.T) {
	// Execute the removeUnaryProductions function
	resultGrammar := RemoveUnaryProductions(grammarTestUnaryElimination, nonTerminalsTest)

	// Compare the result with the expected output
	for key, expectedProductions := range expectedRemoveUnaryProductions.Productions {
		resultProductions, exists := resultGrammar.Productions[key]
		if !exists {
			t.Errorf("Error: No producciones encontradas para %v", key)
		}

		if len(expectedProductions) != len(resultProductions) {
			t.Errorf("Error: Longitud de producciones incorrecta para %v. Esperado %v, pero se obtuvo %v", key, expectedProductions, resultProductions)
		}

		for _, production := range expectedProductions {
			if !containsProduction(resultProductions, production) {
				t.Errorf("Error: Se esperaba %v en %v, pero no se encontró", production, key)
			}
		}
	}
}
