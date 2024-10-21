package grammar

import (
	"fmt"
)

// Dada una gramática, elimina todas las producciones epsilon
func SimplifyGrammar(grammar *Grammar, printSteps bool) *Grammar {

	// Asegurar que el símbolo inicial sea el primer no terminal en la lista nonTerminals
	if len(grammar.nonTerminals) == 0 {
		return grammar
	}

	startSymbol := grammar.nonTerminals[0]

	if printSteps {
		fmt.Println("\n1️⃣  Gramática ANTES de la simplificación:")
		fmt.Println(grammar.String(true))
	}

	fmt.Println("\n2️⃣  ELIMINACIÓN DE EPSILON:")

	// Paso 1: Identificar los símbolos directos anulables
	directNullables := identifyDirectNullables(grammar)
	if printSteps {
		fmt.Println("\n🔴  2.1 Símbolos anulables directos encontrados:")
		fmt.Printf("\t%v\n", *directNullables)
	}

	// Paso 2: Identificar todos los símbolos anulables (directos e indirectos)
	allNullables := identifyIndirectNullables(grammar, *directNullables)
	if printSteps {
		fmt.Println("\n🔴  2.2 Todos los símbolos anulables encontrados:")
		fmt.Printf("\t%v\n", *allNullables)
	}

	// Paso 3: Reemplazar los símbolos anulables en las producciones
	grammarWithoutEpsilons := ReplaceNullables(grammar, *allNullables)
	if printSteps {
		fmt.Println("\n🔴  2.3 Gramática DESPUÉS de reemplazar los anulables:")
		fmt.Println(grammarWithoutEpsilons.String(true))
	}

	// Paso 4: Eliminar producciones épsilon
	finalGrammar1 := RemoveEpsilons(grammarWithoutEpsilons)
	if printSteps {
		fmt.Println("\n🔴  2.4 Gramática DESPUÉS de eliminar las producciones epsilon:")
		fmt.Println(finalGrammar1.String(true))
	}

	// Paso 5: Eliminar producciones unarias
	fmt.Println("\n3️⃣  ELIMINACIÓN DE PRODUCCIONES UNARIAS:")
	finalGrammar2 := RemoveUnaryProductions(finalGrammar1, finalGrammar1.nonTerminals)
	if printSteps {
		fmt.Println("\n🔴  3.1 Gramática DESPUÉS de eliminar producciones unarias:")
		fmt.Println(finalGrammar2.String(true))
	}

	// Paso 6: Eliminar símbolos inútiles
	fmt.Println("\n4️⃣  ELIMINACIÓN DE SIMBOLOS INUTILES:")
	finalGrammar3 := RemoveUselessSymbols(finalGrammar2, startSymbol)
	if printSteps {
		fmt.Println("\n🔴  4.1 Gramática DESPUÉS de eliminar símbolos inútiles:")
		fmt.Println(finalGrammar3.String(true))
	}

	// Paso 7: Normalizar paso 1
	fmt.Println("\n5️⃣  SIMPLIFICACIÓN A FORMA NORMAL DE CHOMSKY:")
	ncfGrammar1 := CNFTerminalSubstitution(finalGrammar3)
	if printSteps {
		fmt.Println("\n🔴  5.1 Gramática DESPUÉS de normalizar el paso 1 de Chomsky:")
		fmt.Println(ncfGrammar1.String(true))
	}

	// Paso 8: Normalizar paso 2
	ncfGrammar2 := CNFSplitLargeProductions(ncfGrammar1)
	if printSteps {
		fmt.Println("\n🔴  5.2 Gramática DESPUÉS de normalizar el paso 2 de Chomsky:")
		fmt.Println(ncfGrammar2.String(true))
	}

	sortGrammar := OrderProductionsByNonTerminals(ncfGrammar2)
	accepted := CYKParse(sortGrammar, "baaba", startSymbol)
	if accepted {
		fmt.Println("La cadena es aceptada por la gramática.")
	} else {
		fmt.Println("La cadena NO es aceptada por la gramática.")
	}

	return sortGrammar
}
