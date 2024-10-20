package grammar

import (
	"fmt"
	"strings"
)

const Epsilon = "ε"

var EpsilonSymbol Symbol = Symbol{
	isTerminal: true,
	value:      "ε",
	id:         0,
}

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
func (g *Grammar) AddProductionFromString(production string) {
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

func (g *Grammar) AddProduction(head string, bodies [][]Symbol) *Symbol {
	// Find the highest ID for the given head value in nonTerminals
	newID := 0
	for _, nonTerminal := range g.nonTerminals {
		if nonTerminal.value == head {
			newID = nonTerminal.id + 1 // Increment the ID to be unique
		}
	}

	// Create a new Symbol for the head with the next available ID
	newHead := Symbol{
		isTerminal: false,
		value:      head,
		id:         newID,
	}

	// Add the new production to the productions map
	g.productions[newHead] = bodies

	// UPDATE nonTerminals an Terminals list:
	// Add the new head to the nonTerminals list if it's not already there
	g.nonTerminals = append(g.nonTerminals, newHead)

	// Add new symbols from the body to Terminals list.
	for _, body := range bodies {
		for _, symbol := range body {
			if symbol.isTerminal {
				g.terminals = append(g.terminals, symbol)
			}
		}
	}
	g.terminals = removeDuplicatesSymbols(g.terminals)

	// Return a reference to the new head Symbol
	return &newHead
}

func (g *Grammar) AddProductionBodies(head Symbol, bodies [][]Symbol) *Symbol {

	if _, exist := g.productions[head]; !exist {
		g.nonTerminals = append(g.nonTerminals, head)
		g.productions[head] = bodies
	} else {
		g.productions[head] = append(g.productions[head], bodies...)
	}

	// Remove duplicate bodies
	g.productions[head] = removeDuplicatesSlices(g.productions[head])

	// UPDATE nonTerminals lists
	// Add new symbols from the body to Terminals list.
	for _, body := range bodies {
		for _, symbol := range body {
			if symbol.isTerminal {
				g.terminals = append(g.terminals, symbol)
			}
		}
	}
	g.terminals = removeDuplicatesSymbols(g.terminals)

	return &head
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

// Removes all occurrences of symbolToRemove from items and returns the updated slice.
func removeSymbols(items *[]Symbol, symbolToRemove *Symbol) *[]Symbol {
	// Create a new slice to hold the result
	result := []Symbol{}

	// Iterate through the items slice
	for _, symbol := range *items {
		// Add to result only if the current symbol is not equal to symbolToRemove
		if symbol != *symbolToRemove {
			result = append(result, symbol)
		}
	}

	return &result
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
// Returns: true if item is made only by items of slice
func isComposedOfSymbols(validSymbols []Symbol, body []Symbol) bool {
	// Create a set for quick lookup of valid symbols
	set := make(map[Symbol]struct{}, len(validSymbols))
	for _, symbol := range validSymbols {
		set[symbol] = struct{}{}
	}

	// Check if every symbol in the body is in the set of valid symbols
	for _, symbol := range body {
		if _, exists := set[symbol]; !exists {
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

// Checks if a slice of symbols (item) exists in a slice of slices (items)
func containsSymbolSlice(items [][]Symbol, item []Symbol) bool {
	for _, slice := range items {
		if areSymbolSlicesEqual(slice, item) { // Use the helper function to compare slices
			return true
		}
	}
	return false
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
