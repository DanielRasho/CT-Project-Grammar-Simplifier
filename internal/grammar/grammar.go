package grammar

import (
	"fmt"
)

// Given a grammar it removes all epsilon productions
func SimplifyGrammar(grammar *Grammar, printSteps bool) *Grammar {

	if len(grammar.nonTerminals) == 0 {
		return grammar
	}

	// Asegurar que el símbolo inicial sea el primer no terminal en la lista nonTerminals
	startSymbol := grammar.nonTerminals[0]

	if printSteps {
		fmt.Println("\n1️⃣  Grammar BEFORE simplification")
		fmt.Println(grammar.String(true))
	}

	//
	fmt.Println("\n2️⃣  EPSILON REMOVAL ")

	// Paso 1: Identificar los símbolos directos anulables
	directNullables := identifyDirectNullables(grammar)
	if printSteps {
		fmt.Println("\n🔴  Direct Nullables found ")
		fmt.Printf("\t%v\n", *directNullables)
	}

	// Paso 2: Identificar todos los símbolos anulables (directos e indirectos)
	allNullables := identifyIndirectNullables(grammar, *directNullables)
	if printSteps {
		fmt.Println("\n🔴  All Nullables found ")
		fmt.Printf("\t%v\n", *allNullables)
	}

	// Paso 3: Reemplazar los símbolos anulables en las producciones
	grammarWithoutEpsilons := ReplaceNullables(grammar, *allNullables)
	if printSteps {
		fmt.Println("\n🔴 Grammar AFTER replacing nullables")
		fmt.Println(grammarWithoutEpsilons.String(true))
	}

	// Paso 4: Eliminar producciones épsilon
	finalGrammar1 := RemoveEpsilons(grammarWithoutEpsilons)
	if printSteps {
		fmt.Println("\n🔴  Grammar AFTER epsilon removal")
		fmt.Println(finalGrammar1.String(true))
	}

	// Paso 5: Eliminar producciones unarias
	finalGrammar2 := RemoveUnaryProductions(finalGrammar1, finalGrammar1.nonTerminals)
	if printSteps {
		fmt.Println("\n🔴  Grammar AFTER unary remove")
		fmt.Println(finalGrammar2.String(true))
	}

	// Paso 6: Eliminar simbolos inutiles
	finalGrammar3 := RemoveUselessSymbols(finalGrammar2, startSymbol)
	if printSteps {
		fmt.Println("\n🔴  Grammar AFTER remove useless symbol")
		fmt.Println(finalGrammar3.String(true))
	}
	return finalGrammar3
}
