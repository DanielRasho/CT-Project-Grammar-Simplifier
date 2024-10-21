package grammar

import (
	"fmt"
	"sort"
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
	NonTerminals []Symbol              // List of all cached NON terminals in the grammar.
	Productions  map[Symbol][][]Symbol // The actual productions.
}

// returns: a readable representation of the grammar.
func (g *Grammar) String(verbose bool) string {
	var sb strings.Builder
	if verbose {
		sb.WriteString(fmt.Sprintf("NonTerminals: %v\n", getSymbolSliceString(&g.NonTerminals)))
		sb.WriteString(fmt.Sprintf("Terminals: %v\n\n", getSymbolSliceString(&g.terminals)))
	}

	// Recorrer las producciones según el orden de los no terminales
	for _, head := range g.NonTerminals {
		if bodies, exists := g.Productions[head]; exists {
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
	if _, exist := g.Productions[head]; !exist {
		// Add new NON terminal
		g.NonTerminals = append(g.NonTerminals, head)
		bodySymbols := make([][]Symbol, 0)
		// Add bodies
		for _, v := range bodyItems {
			body, nonTerminal, terminal := splitStringIntoSymbols(v)
			g.NonTerminals = append(g.NonTerminals, nonTerminal...)
			g.terminals = append(g.terminals, terminal...)
			bodySymbols = append(bodySymbols, body)
		}
		// Remove duplicate bodies.
		bodySymbols = removeDuplicatesSlices(bodySymbols)
		g.Productions[head] = bodySymbols
	} else {
		// Else append the body new items with the old ones
		existentBodyItems := g.Productions[head]
		for _, v := range bodyItems {
			body, nonTerminal, terminal := splitStringIntoSymbols(v)
			g.NonTerminals = append(g.NonTerminals, nonTerminal...)
			g.terminals = append(g.terminals, terminal...)
			existentBodyItems = append(existentBodyItems, body)
		}
		// Remove duplicate bodies.
		existentBodyItems = removeDuplicatesSlices(existentBodyItems)
		g.Productions[head] = existentBodyItems
	}
	g.NonTerminals = removeDuplicatesSymbols(g.NonTerminals)
	g.terminals = removeDuplicatesSymbols(g.terminals)
}

func (g *Grammar) AddProduction(head string, bodies [][]Symbol) *Symbol {
	// Find the highest ID for the given head value in nonTerminals
	newID := 0
	for _, nonTerminal := range g.NonTerminals {
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
	g.Productions[newHead] = bodies

	// UPDATE nonTerminals an Terminals list:
	// Add the new head to the nonTerminals list if it's not already there
	g.NonTerminals = append(g.NonTerminals, newHead)

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

	if _, exist := g.Productions[head]; !exist {
		g.NonTerminals = append(g.NonTerminals, head)
		g.Productions[head] = bodies
	} else {
		g.Productions[head] = append(g.Productions[head], bodies...)
	}

	// Remove duplicate bodies
	g.Productions[head] = removeDuplicatesSlices(g.Productions[head])

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

// Set te bodies of a production (true = succes, false = error)
// but DO NOT update the terminals list. For that use RecalculateTerminals()
func (g *Grammar) SetProductionBodies(head Symbol, bodies [][]Symbol) bool {

	if _, exist := g.Productions[head]; !exist {
		return false
	}

	g.Productions[head] = bodies

	return true
}

func (g *Grammar) RecalculateTerminals() {
	g.terminals = make([]Symbol, 0)
	for _, bodies := range g.Productions {
		for _, body := range bodies {
			for _, symbol := range body {
				if symbol.isTerminal {
					g.terminals = append(g.terminals, symbol)
				}
			}
		}
	}
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

// Función para eliminar duplicados en un slice de strings
func removeDuplicatesString(slice []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, v := range slice {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
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

// Checks if a slice of symbols (item) exists in a slice of slices (items)
func containsSymbolSlice(items [][]Symbol, item []Symbol) bool {
	for _, slice := range items {
		if areSymbolSlicesEqual(slice, item) { // Use the helper function to compare slices
			return true
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

// Función auxiliar para comparar dos gramáticas
func compareGrammars(g1, g2 *Grammar) bool {
	// Comparar terminales
	if !compareSymbolSlices(g1.terminals, g2.terminals) {
		return false
	}

	// Comparar no terminales
	if !compareSymbolSlices(g1.NonTerminals, g2.NonTerminals) {
		return false
	}

	// Comparar producciones
	if len(g1.Productions) != len(g2.Productions) {
		return false
	}

	for head, g1Productions := range g1.Productions {
		g2Productions, exists := g2.Productions[head]
		if !exists {
			return false
		}

		if !compareProductionSlices(g1Productions, g2Productions) {
			return false
		}
	}

	return true
}

// Función auxiliar para comparar dos slices de símbolos (sin importar el orden)
func compareSymbolSlices(s1, s2 []Symbol) bool {
	if len(s1) != len(s2) {
		return false
	}

	// Ordenar ambos slices antes de compararlos
	sort.Slice(s1, func(i, j int) bool {
		return s1[i].String() < s1[j].String()
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i].String() < s2[j].String()
	})

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// Función auxiliar para comparar dos slices de producciones (sin importar el orden)
func compareProductionSlices(p1, p2 [][]Symbol) bool {
	if len(p1) != len(p2) {
		return false
	}

	// Ordenar ambas listas de producciones por su representación de cadena
	sort.Slice(p1, func(i, j int) bool {
		return productionString(p1[i]) < productionString(p1[j])
	})
	sort.Slice(p2, func(i, j int) bool {
		return productionString(p2[i]) < productionString(p2[j])
	})

	for i := range p1 {
		if !compareSymbolSlices(p1[i], p2[i]) {
			return false
		}
	}

	return true
}

// Función auxiliar para convertir una producción en cadena para facilitar la comparación
func productionString(p []Symbol) string {
	result := ""
	for _, symbol := range p {
		result += symbol.String()
	}
	return result
}

// Función para ordenar las producciones de la gramática según el orden de los no terminales
func OrderProductionsByNonTerminals(originalGrammar *Grammar) *Grammar {
	// Crear una nueva gramática para almacenar las producciones ordenadas
	orderedGrammar := &Grammar{
		terminals:    originalGrammar.terminals,
		NonTerminals: originalGrammar.NonTerminals,
		Productions:  make(map[Symbol][][]Symbol),
	}

	// Recorrer la lista de no terminales en el orden dado
	for _, nonTerminal := range originalGrammar.NonTerminals {
		// Si existen producciones para este no terminal en la gramática original
		if productions, exists := originalGrammar.Productions[nonTerminal]; exists {
			// Añadir las producciones a la gramática ordenada en el mismo orden de los no terminales
			orderedGrammar.Productions[nonTerminal] = productions
		}
	}

	return orderedGrammar
}

// Función que busca terminales en las producciones y devuelve los heads
func FindHeadsProducingTerminal(grammar *Grammar, terminalValue string) []string {
	heads := []string{} // Lista para almacenar los heads que producen directamente el terminal

	// Recorrer los no terminales en el orden de la lista de nonTerminals
	for _, nonTerminal := range grammar.NonTerminals {
		// Buscar si existen producciones para ese no terminal
		if bodies, exists := grammar.Productions[nonTerminal]; exists {
			for _, body := range bodies {
				// Verificar si la producción tiene exactamente un símbolo y es un terminal
				if len(body) == 1 && body[0].isTerminal && body[0].value == terminalValue {
					// Añadir el value del head (nonTerminal) a la lista si produce directamente el terminal
					heads = append(heads, nonTerminal.value)
				}
			}
		}
	}

	return heads
}

// Función que busca producciones con un par de no terminales y devuelve los heads
func FindHeadsProducingNonTerminals(grammar *Grammar, nonTerminal1, nonTerminal2 Symbol) []string {
	heads := []string{} // Lista para almacenar los heads que producen el par de no terminales

	// Recorrer los no terminales en el orden de la lista de nonTerminals
	for _, nonTerminal := range grammar.NonTerminals {
		// Buscar si existen producciones para ese no terminal
		if bodies, exists := grammar.Productions[nonTerminal]; exists {
			for _, body := range bodies {
				// Verificar si la producción tiene exactamente dos no terminales
				if len(body) == 2 && !body[0].isTerminal && !body[1].isTerminal {
					// Comparar los no terminales de la producción con los proporcionados
					if body[0] == nonTerminal1 && body[1] == nonTerminal2 {
						// Añadir el value del head (nonTerminal) a la lista si produce el par de no terminales
						heads = append(heads, nonTerminal.value)
					}
				}
			}
		}
	}

	return heads
}
