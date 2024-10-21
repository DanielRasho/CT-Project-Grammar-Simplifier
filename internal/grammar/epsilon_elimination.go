package grammar

// Return a list of all the Direct nullables on the grammar
// Ex: A -> ε
func identifyDirectNullables(grammar *Grammar) *[]Symbol {

	directNullables := make([]Symbol, 0, 3)

	for head, bodies := range grammar.Productions {
		for _, body := range bodies {
			for _, symbol := range body {
				if symbol.value == "ε" && symbol.isTerminal == true {
					directNullables = append(directNullables, head)
					continue
				}
			}
		}
	}
	return &directNullables
}

// Identify indirect nullables
//
// Returns: List of all nullables (direct an indirect)
func identifyIndirectNullables(grammar *Grammar, nullabes []Symbol) *[]Symbol {

	pastNullables := make([]Symbol, len(nullabes))
	newNullables := make([]Symbol, len(nullabes))
	copy(pastNullables, nullabes)
	copy(newNullables, nullabes)

	// The algorithm works by having a list of PAST_NULLABLES and a list of NEW_NULLABLES
	// 1. On each step the new nullable productions found are added to NEW_NULLABLES
	// 2. If NEW_NULLABLES and PAST_NULLABLES remain equal after one step, it means that
	//    no new production where found, hence ALL NULLABLE PRODUCTION WHERE FOUND
	// 3. Else, repeat step 1 and 2.
	for {
		for head, bodies := range grammar.Productions {
			// If the production is already nullable dont analize it.
			if containsSymbol(pastNullables, head) {
				continue
			}
			// else check if it is nullable and add it to the NewNullables list
			for _, body := range bodies {
				isNullable := isComposedOfSymbols(pastNullables, body)
				if isNullable {
					newNullables = append(newNullables, head)
					break
				}
			}
		}

		if len(pastNullables) == len(newNullables) {
			break
		}

		pastNullables = newNullables
	}

	return &newNullables
}

// ReplaceNullables reemplaza las producciones que contienen símbolos anulables.
func ReplaceNullables(grammar *Grammar, nullables []Symbol) *Grammar {
	newGrammar := Grammar{
		terminals:    grammar.terminals,
		nonTerminals: grammar.nonTerminals,
		Productions:  make(map[Symbol][][]Symbol)}

	// Paso 1. Leer cada body de la gramática por cada head
	for head, symbols := range grammar.Productions {
		productionSet := make([][]Symbol, 0) // Mapa para rastrear producciones únicas

		// Cola para procesar producciones pendientes
		queue := append([][]Symbol{}, symbols...)

		// Procesar todas las producciones en la cola
		for len(queue) > 0 {
			production := queue[0]
			queue = queue[1:]

			// Caso 1: No hay símbolos anulables, se añade la producción tal cual si no se ha añadido antes
			if !containsSymbolSlice(productionSet, production) {
				productionSet = append(productionSet, production)
			}

			// Caso 2: Existen símbolos anulables en la producción
			for _, nullable := range nullables {
				if containsSymbol(production, nullable) {
					// Generar todas las combinaciones posibles reemplazando el símbolo nullable
					combinations := CombinationNullables(&nullable, &production)
					for _, newProd := range *combinations {
						// Evitar duplicados y procesar nuevas combinaciones
						if !containsSymbolSlice(productionSet, newProd) {
							productionSet = append(productionSet, production)
							queue = append(queue, newProd)
						}
					}
				}
			}
		}

		newGrammar.Productions[head] = removeDuplicatesSlices(productionSet)
	}

	return &newGrammar
}

// CombinationNullables genera todas las combinaciones posibles al reemplazar el símbolo nullable.
func CombinationNullables(nullable *Symbol, baseProduction *[]Symbol) *[][]Symbol {
	var newProductions [][]Symbol

	// Recorrer la producción y reemplazar el símbolo nullable por epsilon
	for i := 0; i < len(*baseProduction); i++ {
		if (*baseProduction)[i] == *nullable {
			// Crear una nueva producción con el símbolo reemplazado por epsilon
			newProd := make([]Symbol, len(*baseProduction))
			copy(newProd, *baseProduction)
			newProd[i] = EpsilonSymbol
			newProductions = append(newProductions, newProd)
		}
	}

	return &newProductions
}

// RemoveEpsilons elimina los caracteres epsilon de la producción y elimina duplicados
func RemoveEpsilons(grammar *Grammar) *Grammar {
	// Crear una nueva gramática para almacenar las producciones sin epsilon
	newGrammar := Grammar{
		terminals:    grammar.terminals,
		nonTerminals: grammar.nonTerminals,
		Productions:  make(map[Symbol][][]Symbol)}

	// Iterar sobre las cabezas de la gramática y sus producciones
	for head, bodies := range grammar.Productions {
		var newNonEpsilonBodies [][]Symbol

		for _, body := range bodies {
			// Reemplazar epsilon por una cadena vacía usando removeSymbols
			nonEpsilonProduction := removeSymbols(&body, &EpsilonSymbol)

			// Solo agregar producciones no vacías
			if len(*nonEpsilonProduction) > 0 {
				newNonEpsilonBodies = append(newNonEpsilonBodies, *nonEpsilonProduction)
			}
		}

		// Eliminar duplicados en las nuevas producciones
		newNonEpsilonBodies = removeDuplicatesSlices(newNonEpsilonBodies)

		// Evitar agregar entradas vacías en la nueva gramática
		if len(newNonEpsilonBodies) > 0 {
			newGrammar.Productions[head] = newNonEpsilonBodies
		}
	}

	// Eliminar el símbolo epsilon de la lista de terminales, si está presente
	newGrammar.terminals = *removeSymbols(&newGrammar.terminals, &EpsilonSymbol)

	// Retornar la nueva gramática
	return &newGrammar
}
