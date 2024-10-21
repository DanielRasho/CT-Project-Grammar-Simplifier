package grammar

import "testing"

// Simbolos para la gramática
var S2 = Symbol{isTerminal: false, value: "S", id: 1}
var A2 = Symbol{isTerminal: false, value: "A", id: 1}
var B2 = Symbol{isTerminal: false, value: "B", id: 1}
var a2 = Symbol{isTerminal: true, value: "a", id: -1}
var b2 = Symbol{isTerminal: true, value: "b", id: -1}

// Nuevos Simbolos creados
var a_20 = Symbol{isTerminal: false, value: "a", id: 0}
var b_20 = Symbol{isTerminal: false, value: "b", id: 0}
var a_21 = Symbol{isTerminal: false, value: "A_a", id: 0}
var b_21 = Symbol{isTerminal: false, value: "B_b", id: 0}

// Grammar Test
var grammarTestCNFConversion = &Grammar{
	terminals:    []Symbol{a2, b2},
	NonTerminals: []Symbol{S2, A2, B2},
	Productions: map[Symbol][][]Symbol{
		S2: {{a2, A2, a2}, {b2, B2, b2}},
		A2: {{a2}},
		B2: {{b2}},
	},
}

// Resultado de reemplazar terminales por no terminales en producciones largas
var expectedTestCNFCNFTerminalSubstitution = &Grammar{
	terminals:    []Symbol{a2, b2},
	NonTerminals: []Symbol{S2, A2, B2, a_20, b_20},
	Productions: map[Symbol][][]Symbol{
		S2:   {{a_20, A2, a_20}, {b_20, B2, b_20}},
		A2:   {{a2}},
		B2:   {{b2}},
		a_20: {{a2}},
		b_20: {{b2}},
	},
}

// Resultado de reemplazar terminales por no terminales en producciones largas
var expectedTestCNFSplitLargeProductions = &Grammar{
	terminals:    []Symbol{a2, b2},
	NonTerminals: []Symbol{S2, A2, B2, a_20, b_20, a_21, b_21},
	Productions: map[Symbol][][]Symbol{
		S2:   {{a_20, a_21}, {b_20, b_21}},
		A2:   {{a2}},
		B2:   {{b2}},
		a_20: {{a2}},
		b_20: {{b2}},
		a_21: {{A2, a_20}},
		b_21: {{B2, b_20}},
	},
}

// Test para la función ReplaceBodyLarge
func TestReplaceBodyLarge(t *testing.T) {
	// Ejecutar la función ReplaceBodyLarge sobre la gramática de prueba
	result := CNFTerminalSubstitution(grammarTestCNFConversion)

	// Comparar la gramática resultante con la esperada
	if !compareGrammars(result, expectedTestCNFCNFTerminalSubstitution) {
		t.Errorf("Error: La gramática resultante de ReplaceBodyLarge no coincide con la esperada.\nEsperado: %v\nObtenido: %v", expectedTestCNFCNFTerminalSubstitution.String(true), result.String(true))
	}
}

// Test para la función CNFSplitLargeProductions
func TestCNFSplitLargeProductions(t *testing.T) {
	// Ejecutar la función CNFSplitLargeProductions sobre la gramática de prueba
	result := CNFSplitLargeProductions(expectedTestCNFSplitLargeProductions)

	// Comparar la gramática resultante con la esperada
	if !compareGrammars(result, expectedTestCNFSplitLargeProductions) {
		t.Errorf("Error: La gramática resultante de CNFSplitLargeProductions no coincide con la esperada.\nEsperado: %v\nObtenido: %v", expectedTestCNFSplitLargeProductions.String(true), result.String(true))
	}
}
