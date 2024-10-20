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
func (g *Grammar) String() string {
	var sb strings.Builder
	sb.WriteString("\n")

	for head, bodies := range g.productions {
		sb.WriteString(head.String())
		sb.WriteString(" -> ")
		for index, body := range bodies {
			for _, symbol := range body {
				sb.WriteString(symbol.String())
			}
			if index < len(body)-1 {
				sb.WriteString("|")
			}
		}
		sb.WriteString("\n")
	}

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
