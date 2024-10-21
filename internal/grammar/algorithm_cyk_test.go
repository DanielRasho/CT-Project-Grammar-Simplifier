package grammar

import "testing"

// Definir los símbolos
var SCYK = Symbol{isTerminal: false, value: "S", id: 0}
var ACYK = Symbol{isTerminal: false, value: "A", id: 0}
var BCYK = Symbol{isTerminal: false, value: "B", id: 0}
var CCYK = Symbol{isTerminal: false, value: "C", id: 0}
var aCYK = Symbol{isTerminal: true, value: "a", id: -1}
var bCYK = Symbol{isTerminal: true, value: "b", id: -1}

// Definir la gramática
var testGrammar = &Grammar{
	terminals:    []Symbol{aCYK, bCYK},
	nonTerminals: []Symbol{SCYK, ACYK, BCYK, CCYK},
	Productions: map[Symbol][][]Symbol{
		SCYK: {{ACYK, BCYK}, {BCYK, CCYK}},
		ACYK: {{BCYK, ACYK}, {aCYK}},
		BCYK: {{CCYK, CCYK}, {bCYK}},
		CCYK: {{ACYK, BCYK}, {aCYK}},
	},
}

func TestCYKParse(t *testing.T) {
	// Cadena a probar
	cadena := "baaba"

	// Llamar a la función CYKParse con la gramática y la cadena
	resultado := CYKParse(testGrammar, cadena, SCYK)

	// Verificar si el resultado es true, ya que la cadena debería ser aceptada por la gramática
	if !resultado {
		t.Errorf("Error: La cadena '%s' debería ser aceptada por la gramática, pero fue rechazada.", cadena)
	} else {
		t.Logf("La cadena '%s' fue aceptada correctamente por la gramática.", cadena)
	}
}
