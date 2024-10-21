package grammar

import "testing"

// Definir los símbolos
var SCYK = Symbol{IsTerminal: false, Value: "S", Id: 0}
var ACYK = Symbol{IsTerminal: false, Value: "A", Id: 0}
var BCYK = Symbol{IsTerminal: false, Value: "B", Id: 0}
var CCYK = Symbol{IsTerminal: false, Value: "C", Id: 0}
var aCYK = Symbol{IsTerminal: true, Value: "a", Id: -1}
var bCYK = Symbol{IsTerminal: true, Value: "b", Id: -1}

// Definir la gramática
var testGrammar = &Grammar{
	terminals:    []Symbol{aCYK, bCYK},
	NonTerminals: []Symbol{SCYK, ACYK, BCYK, CCYK},
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
