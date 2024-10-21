package grammar

import (
	"fmt"
	"strings"
)

const Epsilon = "ε"

type Symbol struct {
	isTerminal bool
	value      string
	id         int
}

func (s *Symbol) String() string {
	if s.isTerminal {
		return s.value
	}
	return fmt.Sprintf("{%s_%d}", s.value, s.id)
}

type Grammar struct {
	terminals    []Symbol              // List of all cached terminals in the grammar.
	nonTerminals []Symbol              // List of all cached NON terminals in the grammar.
	productions  map[Symbol][][]Symbol // The actual productions.
}

// returns: a readable representation of the grammar.
func (g *Grammar) String(verbose bool) string {
	var sb strings.Builder
	if verbose {
		sb.WriteString(fmt.Sprintf("NonTerminals: %v\n", getSymbolSliceString(&g.nonTerminals)))
		sb.WriteString(fmt.Sprintf("Terminals: %v\n\n", getSymbolSliceString(&g.terminals)))
	}

	for head, bodies := range g.productions {
		sb.WriteString(head.String())
		sb.WriteString(" -> ")
		for index, body := range bodies {
			for _, symbol := range body {
				sb.WriteString(symbol.String())
			}
			if index != len(bodies)-1 {
				sb.WriteString("|")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func getSymbolSliceString(slice *[]Symbol) string {
	var sb strings.Builder
	sb.WriteString("[")
	for index, symbol := range *slice {
		sb.WriteString(symbol.String())
		if index != len(*slice)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

// Adds a production to a grammar, removing repeated body values.
func (g *Grammar) AddProduction(production string) {
	// Since a production has the shape A -> a|{B}C
	// There are 2 divisions between the Head, Arrow, And Body.
	division1 := strings.Index(production, " ")                               // Find first space index
	division2 := division1 + 1 + strings.Index(production[division1+1:], " ") // Find second space index

	head := Symbol{value: production[:division1], isTerminal: false, id: 0}
	body := production[division2+1:]
	bodyItems := strings.Split(body, "|")

	// If production is not registered create it
	if _, exist := g.productions[head]; !exist {
		// Add new NON terminal
		g.nonTerminals = append(g.nonTerminals, head)
		bodySymbols := make([][]Symbol, 0)
		// Add bodies
		for _, v := range bodyItems {
			body, nonTerminal, terminal := splitStringIntoSymbols(v)
			g.nonTerminals = append(g.nonTerminals, nonTerminal...)
			g.terminals = append(g.terminals, terminal...)
			bodySymbols = append(bodySymbols, body)
		}
		// Remove duplicate bodies.
		bodySymbols = removeDuplicatesSlices(bodySymbols)
		g.productions[head] = bodySymbols
	} else {
		// Else append the body new items with the old ones
		existentBodyItems := g.productions[head]
		for _, v := range bodyItems {
			body, nonTerminal, terminal := splitStringIntoSymbols(v)
			g.nonTerminals = append(g.nonTerminals, nonTerminal...)
			g.terminals = append(g.terminals, terminal...)
			existentBodyItems = append(existentBodyItems, body)
		}
		// Remove duplicate bodies.
		existentBodyItems = removeDuplicatesSlices(existentBodyItems)
		g.productions[head] = existentBodyItems
	}
	g.nonTerminals = removeDuplicatesSymbols(g.nonTerminals)
	g.terminals = removeDuplicatesSymbols(g.terminals)
}

// Revoves duplicates on a slice.
func removeDuplicatesSymbols(slice []Symbol) []Symbol {
	uniqueMap := make(map[Symbol]bool)
	var result []Symbol

	for _, item := range slice {
		if _, exists := uniqueMap[item]; !exists {
			uniqueMap[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Helper function to check if two slices of Symbol are equal
func areSymbolSlicesEqual(a, b []Symbol) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] { // Compare individual Symbol structs
			return false
		}
	}
	return true
}

// Remove the duplicate slices within a slice of slices.
func removeDuplicatesSlices(slice [][]Symbol) [][]Symbol {
	unique := [][]Symbol{}

	for _, current := range slice {
		found := false
		for _, uniq := range unique {
			if areSymbolSlicesEqual(current, uniq) {
				found = true
				break
			}
		}
		if !found {
			unique = append(unique, current)
		}
	}

	return unique
}

// Split the string into Symbols and return body, nonTerminals, and terminals
func splitStringIntoSymbols(input string) (body []Symbol, nonTerminals []Symbol, terminals []Symbol) {
	var current strings.Builder
	inBraces := false

	for _, char := range input {
		switch char {
		case '{':
			inBraces = true // We are inside curly braces

		case '}':
			if inBraces {
				// Add non-terminal symbol to body and nonTerminals slice
				nonTerminalSymbol := Symbol{value: current.String(), isTerminal: false, id: 0}
				body = append(body, nonTerminalSymbol)
				nonTerminals = append(nonTerminals, nonTerminalSymbol)
				current.Reset()
				inBraces = false // Exiting curly braces
			}

		default:
			current.WriteRune(char) // Build the current symbol
			if !inBraces {
				terminalSymbol := Symbol{value: current.String(), isTerminal: true, id: 0}
				body = append(body, terminalSymbol)
				terminals = append(terminals, terminalSymbol)
				current.Reset()
			}
		}
	}

	// Add any remaining non-braced symbol as terminal
	if current.Len() > 0 {
		terminalSymbol := Symbol{value: current.String(), isTerminal: true, id: 0}
		body = append(body, terminalSymbol)
		terminals = append(terminals, terminalSymbol)
	}

	return body, nonTerminals, terminals
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

// Checks if a string exists in a slice
func containsSymbol(slice []Symbol, item Symbol) bool {
	for _, s := range slice {
		if s == item { // Compare each Symbol in the slice
			return true
		}
	}
	return false
}

// Verifica si una producción ya existe en la lista de producciones
func containsProduction(productions [][]Symbol, production []Symbol) bool {
	for _, p := range productions {
		if len(p) == len(production) {
			match := true
			for i := range p {
				if p[i] != production[i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

// Helper function que verifica si una producción está compuesta solo de símbolos generadores.
func isComposedOfSymbol(generatingSymbols []Symbol, production []Symbol) bool {
	generatingSet := make(map[Symbol]struct{})
	for _, sym := range generatingSymbols {
		generatingSet[sym] = struct{}{}
	}

	for _, sym := range production {
		if _, exists := generatingSet[sym]; !exists {
			return false
		}
	}
	return true
}

// Helper function para obtener las claves de un mapa.
func getKeys(m map[Symbol]struct{}) []Symbol {
	keys := make([]Symbol, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Función auxiliar para comparar dos gramáticas
func compareGrammars(g1, g2 *Grammar) bool {
	// Comparar terminales
	if !compareSymbolSlices(g1.terminals, g2.terminals) {
		return false
	}

	// Comparar no terminales
	if !compareSymbolSlices(g1.nonTerminals, g2.nonTerminals) {
		return false
	}

	// Comparar producciones
	if len(g1.productions) != len(g2.productions) {
		return false
	}

	for head, g1Productions := range g1.productions {
		g2Productions, exists := g2.productions[head]
		if !exists {
			return false
		}

		if !compareProductionSlices(g1Productions, g2Productions) {
			return false
		}
	}

	return true
}

// Función auxiliar para comparar dos slices de símbolos
func compareSymbolSlices(s1, s2 []Symbol) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// Función auxiliar para comparar dos slices de producciones
func compareProductionSlices(p1, p2 [][]Symbol) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i := range p1 {
		if !compareSymbolSlices(p1[i], p2[i]) {
			return false
		}
	}

	return true
}
