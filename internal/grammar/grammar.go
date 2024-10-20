package grammar

/*
import (
	"fmt"
	"strings"
)

// Given a grammar it removes all epsilon productions
func SimplifyGrammar(grammar *Grammar, printSteps bool) *Grammar {
	if printSteps {
		fmt.Println("\n1️⃣  Grammar BEFORE simplification")
		for head, productions := range *grammar {
			fmt.Printf("%s -> %s\n", head, strings.Join(productions, " | "))
		}
	}

	// Paso 1: Identificar los símbolos directos anulables
	directNullables := identifyDirectNullables(grammar)
	if printSteps {
		fmt.Println("\n2️⃣  Direct Nullables found ")
		fmt.Printf("\t%v\n", *directNullables)
	}

	// Paso 2: Identificar todos los símbolos anulables (directos e indirectos)
	allNullables := identifyIndirectNullables(grammar, *directNullables)
	if printSteps {
		fmt.Println("\n3️⃣  All Nullables found ")
		fmt.Printf("\t%v\n", *allNullables)
	}

	// Paso 3: Reemplazar los símbolos anulables en las producciones
	grammarWithoutEpsilons := ReplaceNullables(grammar, *allNullables)
	if printSteps {
		fmt.Println("\n4️⃣  Grammar AFTER replacing nullables")
		for head, productions := range *grammarWithoutEpsilons {
			fmt.Printf("%s -> %s\n", head, strings.Join(productions, " | "))
		}
	}

	// Paso 4: Eliminar producciones épsilon
	finalGrammar := RemoveEpsilons(grammarWithoutEpsilons)
	if printSteps {
		fmt.Println("\n5️⃣  Grammar AFTER epsilon removal")
		for head, productions := range *finalGrammar {
			fmt.Printf("%s -> %s\n", head, strings.Join(productions, " | "))
		}
	}

	return finalGrammar
}
*/
