package grammar

import (
	"strings"
)

type leftFactorPrefix struct {
	prefix []Symbol
	count  int
}

func factorizeGrammar(grammar *Grammar) *Grammar {

	for _, nonTerminal := range grammar.NonTerminals {
		if productions, exist := grammar.Productions[nonTerminal]; exist {
			result := leftFactor(grammar, &nonTerminal, productions)
			grammar.SetProductionBodies(nonTerminal, result)
		}
	}

	grammar.RecalculateTerminals()

	return grammar
}

func leftFactor(grammar *Grammar, head *Symbol, bodies [][]Symbol) [][]Symbol {
	// Step 1: Find the longest common prefix
	prefix, prefixBodies, notPrefixBodies := findLongestCommonPrefix(bodies)

	// If no prefix, return the original bodies (no factoring possible)
	if prefix == nil {
		return bodies
	}

	newBodies := make([][]Symbol, 0)

	// Step 2: Create a new non-terminal for the factored suffixes
	newNonTerminal := grammar.AddProduction(head.Value, [][]Symbol{})
	newNonTerminalBodies := leftFactor(grammar, newNonTerminal, prefixBodies)
	grammar.SetProductionBodies(*newNonTerminal, newNonTerminalBodies)

	// Step 3
	factoredBody := append(prefix, *newNonTerminal)
	newBodies = append(newBodies, factoredBody)

	// Step 3
	if len(notPrefixBodies) == 0 {
		return newBodies
	}

	temp := leftFactor(grammar, head, notPrefixBodies)
	newBodies = append(newBodies, temp...)

	return newBodies
}

// Finds the most frequent and longest prefix within all bodies.
//
// Returns:
//
//	[]Symbol : common prefix
//	[][]Symbol : List of all the bodies that had the prefixed, with the prefix removed
//	[][]Symbol : List of all the bodies that do not had the prefix.
func findLongestCommonPrefix(bodies [][]Symbol) ([]Symbol, [][]Symbol, [][]Symbol) {
	if len(bodies) == 0 {
		return nil, nil, nil
	}

	// This will map each potential prefix to its occurrence count
	prefixCountMap := make(map[string]leftFactorPrefix)

	// Check each body against all others to find common prefixes
	for i := 0; i < len(bodies); i++ {
		body := bodies[i]
		for j := i + 1; j < len(bodies); j++ {
			otherBody := bodies[j]

			// Find the common prefix between two bodies
			prefix := commonPrefix(body, otherBody)

			// fmt.Printf("COMMON PREFIX FOUND: %v\n", prefix)

			if len(prefix) > 0 {
				// Convert the prefix to a string to use as the key
				prefixKey := symbolsToString(prefix)

				// if the prefix exist increase the counter
				if _, exist := prefixCountMap[prefixKey]; exist {
					entry := prefixCountMap[prefixKey]
					entry.count++
					prefixCountMap[prefixKey] = entry
					// else, add the entry
				} else {
					prefixCountMap[prefixKey] = leftFactorPrefix{
						prefix: prefix,
						count:  1,
					}
				}
			}
		}
	}

	// Now, find the most frequent and longest common prefix
	var longestPrefix []Symbol
	var maxOccurrences int

	for _, value := range prefixCountMap {
		if value.count > maxOccurrences {
			longestPrefix = value.prefix // Convert the string back to symbols
			maxOccurrences = value.count
		}
	}

	// Filter out bodies that had the longest common prefix
	var prefixBodies [][]Symbol
	var notPrefixBodies [][]Symbol
	bodyLen := len(longestPrefix)
	for _, body := range bodies {
		if startsWith(body, longestPrefix) {
			newBody := body[bodyLen:]
			if len(newBody) == 0 {
				newBody = append(newBody, EpsilonSymbol)
			}
			prefixBodies = append(prefixBodies, newBody)
			// fmt.Println(body)
			// fmt.Println(newBody)
			// fmt.Println("---------")
			continue
		}
		notPrefixBodies = append(notPrefixBodies, body)
	}

	return longestPrefix, prefixBodies, notPrefixBodies
}

// Helper function to find the common prefix between two symbol slices
func commonPrefix(a, b []Symbol) []Symbol {
	minLen := min(len(a), len(b))
	var prefix []Symbol

	for i := 0; i < minLen; i++ {
		if a[i] == b[i] {
			prefix = append(prefix, a[i])
		} else {
			break
		}
	}

	return prefix
}

// Check if a body starts with a given prefix
func startsWith(body, prefix []Symbol) bool {
	if len(body) < len(prefix) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if body[i] != prefix[i] {
			return false
		}
	}

	return true
}

// Convert a slice of Symbols to a string (for use as a map key)
func symbolsToString(symbols []Symbol) string {
	var sb strings.Builder
	for _, sym := range symbols {
		sb.WriteString(sym.String())
	}
	return sb.String()
}
