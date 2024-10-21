package grammar

import (
	"fmt"
)

// Given a grammar it removes all epsilon productions
func SimplifyGrammar(grammar *Grammar, printSteps bool) *Grammar {
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
	finalGrammar := RemoveEpsilons(grammarWithoutEpsilons)
	if printSteps {
		fmt.Println("\n🔴  Grammar AFTER epsilon removal")
		fmt.Println(finalGrammar.String(true))
	}

	return finalGrammar
}
