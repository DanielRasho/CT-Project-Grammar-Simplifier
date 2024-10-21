package grammar

import (
	"fmt"
	"sync"
)

func removeLeftRecursivity(originalGrammar *Grammar) *Grammar {
	// Create a new grammar
	newGrammar := Grammar{
		Productions:  make(map[Symbol][][]Symbol),
		nonTerminals: make([]Symbol, len(originalGrammar.nonTerminals)),
		terminals:    make([]Symbol, len(originalGrammar.terminals)),
	}
	// Deep copy nonTerminals and Terminals
	copy(newGrammar.nonTerminals, originalGrammar.nonTerminals)
	copy(newGrammar.terminals, originalGrammar.terminals)

	for head := range originalGrammar.Productions {

		AllBodyVariants := findAllBodyVariants(&head, originalGrammar)

		// Create the two new pro
		recursiveBodies := make([][]Symbol, 0)    // List of bodies that has left recursivity
		nonRecursiveBodies := make([][]Symbol, 0) // List of bodies that DO Not have recursivityh

		for _, bodyVariant := range AllBodyVariants {
			if bodyVariant[0] == head {
				recursiveBodies = append(recursiveBodies, bodyVariant)
			} else {
				nonRecursiveBodies = append(nonRecursiveBodies, bodyVariant)
			}
		}

		// If recursive bodies were not found,
		if len(recursiveBodies) == 0 {
			newGrammar.AddProduction(head.value, AllBodyVariants) // A
			continue
		}
		// If nonRecursiveBodies is empty, add epsilon
		if len(nonRecursiveBodies) == 0 {
			nonRecursiveBodies = append(nonRecursiveBodies, []Symbol{EpsilonSymbol})
		}

		fmt.Println()

		// If not, solve the recursion

		// FIX RECURSION
		// Create the productions
		production1 := newGrammar.AddProduction(head.value, make([][]Symbol, 0)) // A
		production2 := newGrammar.AddProduction(head.value, make([][]Symbol, 0)) // A'

		// Modify bodies to remove recursion
		var wg sync.WaitGroup
		wg.Add(2)
		go fixRecursiveBodies(&wg, &recursiveBodies, production2)
		go fixNonRecursiveBodies(&wg, &nonRecursiveBodies, production2)
		wg.Wait()
		// Add epsilon at the end of recursive bodies, to ensure the form A' -> αA'|...|ε
		recursiveBodies = append(recursiveBodies, []Symbol{EpsilonSymbol})
		// Eliminite Duplicates.
		recursiveBodies = removeDuplicatesSlices(recursiveBodies)
		nonRecursiveBodies = removeDuplicatesSlices(nonRecursiveBodies)
		// Add fixed bodies to grammar.
		newGrammar.AddProductionBodies(*production1, nonRecursiveBodies)
		newGrammar.AddProductionBodies(*production2, recursiveBodies)

	}

	return &newGrammar
}

func fixRecursiveBodies(wg *sync.WaitGroup, bodies *[][]Symbol, secondProduction *Symbol) {
	defer wg.Done()
	for i := range *bodies {
		(*bodies)[i] = (*bodies)[i][1:]                        // Remove the first element that causes the recursivity, to left with the rest: α
		(*bodies)[i] = append((*bodies)[i], *secondProduction) // Add A' at the end of the recursive body, so that we end up with a body with form: αA'
	}
}

func fixNonRecursiveBodies(wg *sync.WaitGroup, bodies *[][]Symbol, secondProduction *Symbol) {
	defer wg.Done()
	for i := range *bodies {
		(*bodies)[i] = append((*bodies)[i], *secondProduction) // Add A' at the end of the recursive body, so that we end up with a body with form: βA'
	}

}

// findAllBodyVariants finds all body variants for a given head symbol,
// including direct and indirect productions.
func findAllBodyVariants(head *Symbol, grammar *Grammar) [][]Symbol {
	var result [][]Symbol
	visited := make(map[Symbol]bool)
	visited[*head] = true

	var explore func(searchSymbol *Symbol) [][]Symbol

	// Recursive function to expand bodies
	explore = func(searchSymbol *Symbol) [][]Symbol {
		var variants [][]Symbol

		if searchSymbol.isTerminal {
			return nil
		}
		if _, isVisited := visited[*searchSymbol]; isVisited {
			return nil
		}

		visited[*searchSymbol] = true

		for _, production := range grammar.Productions[*searchSymbol] {
			variants = append(variants, production)
			temp := explore(&production[0])
			if temp != nil {
				rest := production[1:]
				for _, r := range temp {
					variants = append(variants, append(r, rest...))
				}
			}
		}

		return variants
	}

	// Expand all bodies for the head symbol
	for _, body := range grammar.Productions[*head] {
		first := body[0]
		variants := explore(&first)
		variants = append(variants, []Symbol{body[0]})
		variants = removeDuplicatesSlices(variants)
		rest := body[1:]

		for _, variant := range variants {
			// fmt.Println(variant)
			newBodyVariant := append(variant, rest...)
			result = append(result, newBodyVariant)
		}
	}

	return result
}
