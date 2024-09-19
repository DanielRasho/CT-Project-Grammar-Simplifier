package grammar

import (
	"strings"
)

/*
{
	"A" : ["ab", "bb", "Ba"],
	"B" : ["ab", "bb", "Ba"],
	"C" : ["ab", "bb", "epsilon"],
}
*/
// Definition of Grammar. Where the key is the production head, and
// the value is the production's body.
type Grammar map[string][]string

// PrintMap prints a map[string][]string in a readable format.
func (g *Grammar) String() string {
	var sb strings.Builder

	for key, values := range *g {
		sb.WriteString("Key: ")
		sb.WriteString(key)
		sb.WriteString("\n")
		sb.WriteString("\tValues: [")
		sb.WriteString(strings.Join(values, ", "))
		sb.WriteString("]\n")
	}

	return sb.String()
}

// Obtener el conjunto de no terminales (las claves del mapa).
func getNonTerminals(grammar *Grammar) map[string]struct{} {
	nonTerminals := make(map[string]struct{})

	// Los no terminales son las claves del mapa.
	for head := range *grammar {
		nonTerminals[head] = struct{}{}
	}

	return nonTerminals
}

// Obtener el conjunto de terminales.
func getTerminals(grammar *Grammar, nonTerminals map[string]struct{}) map[string]struct{} {
	terminals := make(map[string]struct{})

	// Recorrer cada producción de la gramática.
	for _, productions := range *grammar {
		for _, production := range productions {
			// Recorrer cada símbolo en la producción.
			for _, symbol := range production {
				symbolStr := string(symbol)

				// Si el símbolo no es un no terminal y no está en el conjunto de no terminales.
				if _, isNonTerminal := nonTerminals[symbolStr]; !isNonTerminal {
					terminals[symbolStr] = struct{}{}
				}
			}
		}
	}

	return terminals
}
