package grammar

import (
	"fmt"
	"strings"
)

const Epsilon = "ε"

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

// Adds a production to a grammar, removing repeated body values.
func (g *Grammar) AddProduction(production string) {
	// Since a production has the shape A -> a|bC
	// There are 2 divisions between the Head, Arrow, And Body.
	division1 := strings.Index(production, " ")                               // Find first space index
	division2 := division1 + 1 + strings.Index(production[division1+1:], " ") // Find second space index

	head := production[:division1]
	body := production[division2+1:]
	bodyItems := strings.Split(body, "|")

	// If production is not registered create it
	if _, exist := (*g)[head]; !exist {
		(*g)[head] = RemoveDuplicates(bodyItems)
	} else {
		// Else append the body new items with the old ones
		existentBodyItems := (*g)[head]
		(*g)[head] = RemoveDuplicates(append(existentBodyItems, bodyItems...))
	}
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

// slice: sliceof single character strings,
//
// item: string to check
//
// Returns: true if item is make only by items of slice
func isComposedOf(slice []string, item string) bool {
	// Create a set for quick lookup
	set := make(map[string]struct{}, len(slice))
	for _, char := range slice {
		set[char] = struct{}{}
	}

	// Check if every character in the item is in the set
	for _, char := range item {
		if _, exists := set[string(char)]; !exists {
			return false
		}
	}
	return true
}

// Checks if a string exists in a slice
func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

// Función que determina si un símbolo es terminal.
func isTerminal(symbol string, terminals map[string]struct{}, nonTerminals map[string]struct{}) bool {
	// Un símbolo es terminal si no está en el conjunto de no terminales.
	if _, exists := nonTerminals[symbol]; !exists {
		return true
	}
	return isComposedOf(getKeys(terminals), symbol) // Usar la función isComposedOf para verificar
}

// Función que determina si un símbolo es no terminal.
func isNonTerminal(symbol string, nonTerminals map[string]struct{}) bool {
	_, exists := nonTerminals[symbol]
	return exists
}

// Función que obtiene las claves de un mapa como un slice.
func getKeys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Función que divide una producción separada por espacios en sus símbolos.
func splitProduction(production string) []string {
	// Dividimos la producción usando los espacios en blanco como delimitador
	return strings.Fields(production)
}

// imprime el resultado de parejas unarias
func printResultUnarPairs(unaryPairs map[string][]string) {
	// Imprimir el contenido de unaryPairs
	for key, values := range unaryPairs {
		// Imprimir la clave
		fmt.Print(key + ":")
		// Imprimir los valores asociados a la clave
		for _, value := range values {
			fmt.Print(value + "-")
		}
		fmt.Println() // Nueva línea al final de cada clave
	}
}

// printGrammar imprime la gramática de manera legible
func printGrammar(g *Grammar) {
	for key, productions := range *g {
		fmt.Printf("%s -> ", key)
		for i, production := range productions {
			if i > 0 {
				fmt.Print(" | ") // separador entre producciones
			}
			fmt.Print(production)
		}
		fmt.Println() // nueva línea después de cada no terminal
	}
}

// Función para imprimir los no terminales que generan cadenas de terminales.
func printGeneratingSymbols(generatingSymbols map[string]struct{}) {
	if len(generatingSymbols) == 0 {
		fmt.Println("No symbols found.")
		return
	}

	for symbol := range generatingSymbols {
		fmt.Println(symbol)
	}
}
