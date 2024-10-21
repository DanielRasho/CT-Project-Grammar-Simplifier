package grammar

/*
Función que encuentra los símbolos que generan cadenas de terminales.
*/
func findGeneratingSymbols(originalGrammar *Grammar) ([]Symbol, []Symbol) {
	// Obtener no terminales y terminales
	nonTerminals := originalGrammar.nonTerminals
	terminals := originalGrammar.terminals

	// Listas para almacenar los símbolos generadores y no generadores
	var generatingSymbols []Symbol
	var nonGeneratingSymbols []Symbol

	changed := true
	for changed {
		changed = false

		// Revisar cada no terminal y sus producciones
		for _, nonTerminal := range nonTerminals {
			// Si ya existe en la lista de símbolos generadores, se salta
			if containsSymbol(generatingSymbols, nonTerminal) {
				continue
			}

			// Verificar si el no terminal tiene producciones
			productions, exists := originalGrammar.productions[nonTerminal]
			if !exists {
				// Si no tiene producciones, se considera no generador
				if !containsSymbol(nonGeneratingSymbols, nonTerminal) {
					nonGeneratingSymbols = append(nonGeneratingSymbols, nonTerminal)
				}
				continue
			}

			// Revisar cada producción del no terminal
			for _, production := range productions {
				// Verificar si toda la producción genera (es decir, si está compuesta solo de terminales o no terminales generadores)
				if isComposedOfSymbol(append(terminals, generatingSymbols...), production) {
					// Si genera, añadimos el no terminal al conjunto de generadores
					generatingSymbols = append(generatingSymbols, nonTerminal)
					changed = true
					break
				}
			}
		}
	}

	// Identificar los no generadores que no se hayan detectado anteriormente
	for _, nonTerminal := range nonTerminals {
		if !containsSymbol(generatingSymbols, nonTerminal) && !containsSymbol(nonGeneratingSymbols, nonTerminal) {
			nonGeneratingSymbols = append(nonGeneratingSymbols, nonTerminal)
		}
	}

	return generatingSymbols, nonGeneratingSymbols
}

/*
Función que encuentra los símbolos alcanzables a partir de un símbolo inicial.
*/
func findReachableSymbols(originalGrammar *Grammar, startSymbol Symbol) ([]Symbol, []Symbol) {
	nonTerminals := originalGrammar.nonTerminals

	// Listas para símbolos alcanzables y no alcanzables
	var reachableSymbols []Symbol
	var unreachableSymbols []Symbol

	// Utilizamos una cola para procesar los símbolos alcanzables.
	queue := []Symbol{startSymbol}

	// Mientras haya elementos en la cola.
	for len(queue) > 0 {
		// Extraemos el primer elemento de la cola.
		current := queue[0]
		queue = queue[1:]

		// Si ya hemos marcado el símbolo como alcanzable, continuamos.
		if containsSymbol(reachableSymbols, current) {
			continue
		}

		// Marcamos el símbolo actual como alcanzable.
		reachableSymbols = append(reachableSymbols, current)

		// Verificamos si el símbolo actual tiene producciones
		productions, exists := originalGrammar.productions[current]
		if !exists {
			// Si no tiene producciones, se ignora (el símbolo es alcanzable pero no tiene producciones)
			continue
		}

		// Recorremos todas las producciones del símbolo actual.
		for _, production := range productions {
			// Procesamos cada símbolo en la producción.
			for _, symbol := range production {
				// Verificamos si el símbolo es un no terminal.
				if containsSymbol(nonTerminals, symbol) {
					// Si el símbolo aún no está en los alcanzables, lo añadimos a la cola.
					if !containsSymbol(reachableSymbols, symbol) {
						queue = append(queue, symbol)
					}
				}
			}
		}
	}

	// Identificamos los no terminales que no son alcanzables.
	for _, nonTerminal := range nonTerminals {
		if !containsSymbol(reachableSymbols, nonTerminal) {
			unreachableSymbols = append(unreachableSymbols, nonTerminal)
		}
	}

	return reachableSymbols, unreachableSymbols
}

// Función que elimina los símbolos no generadores de la gramática y retorna una nueva gramática.
func RemoveNonGeneratingSymbols(originalGrammar *Grammar) *Grammar {
	// Obtener los símbolos no generadores
	_, nonGeneratingSymbols := findGeneratingSymbols(originalGrammar)

	// Crear una nueva gramática vacía
	newGrammar := &Grammar{
		terminals:    []Symbol{},
		nonTerminals: []Symbol{},
		productions:  make(map[Symbol][][]Symbol),
	}

	// Procesar las producciones de la gramática original
	for head, productions := range originalGrammar.productions {
		// Lista de producciones válidas (sin símbolos no generadores)
		var validProductions [][]Symbol

		// Verificar cada producción
		for _, production := range productions {
			containsNonGenerating := false

			// Verificar si la producción contiene algún símbolo no generador
			for _, symbol := range production {
				if containsSymbol(nonGeneratingSymbols, symbol) {
					containsNonGenerating = true
					break
				}
			}

			// Si no contiene símbolos no generadores, la producción es válida
			if !containsNonGenerating {
				validProductions = append(validProductions, production)
			}
		}

		// Si hay producciones válidas, añadirlas a la nueva gramática
		if len(validProductions) > 0 {
			newGrammar.productions[head] = validProductions
			newGrammar.nonTerminals = append(newGrammar.nonTerminals, head)
		}
	}

	// Recorrer la nueva gramática para identificar los terminales y no terminales
	for _, productions := range newGrammar.productions {
		for _, production := range productions {
			for _, symbol := range production {
				if symbol.isTerminal {
					// Añadir a la lista de terminales si no está ya presente
					if !containsSymbol(newGrammar.terminals, symbol) {
						newGrammar.terminals = append(newGrammar.terminals, symbol)
					}
				} else {
					// Añadir a la lista de no terminales si no está ya presente
					if !containsSymbol(newGrammar.nonTerminals, symbol) {
						newGrammar.nonTerminals = append(newGrammar.nonTerminals, symbol)
					}
				}
			}
		}
	}

	// Retornar la nueva gramática sin símbolos no generadores
	return newGrammar
}

// Función que elimina los símbolos no alcanzables de la gramática y retorna una nueva gramática.
func RemoveNonReachableSymbols(originalGrammar *Grammar) *Grammar {
	// Inicializar la variable startSymbol (el primer símbolo de las producciones)
	var startSymbol Symbol
	for symbol := range originalGrammar.productions {
		startSymbol = symbol
		break
	}

	// Obtener los símbolos no alcanzables
	_, unreachableSymbols := findReachableSymbols(originalGrammar, startSymbol)

	// Crear una nueva gramática vacía
	newGrammar := &Grammar{
		terminals:    []Symbol{},
		nonTerminals: []Symbol{},
		productions:  make(map[Symbol][][]Symbol),
	}

	// Procesar las producciones de la gramática original
	for head, productions := range originalGrammar.productions {
		// Si el head (no terminal) está en los símbolos no alcanzables, lo omitimos
		if containsSymbol(unreachableSymbols, head) {
			continue
		}

		// Añadir las producciones válidas a la nueva gramática
		newGrammar.productions[head] = productions
		newGrammar.nonTerminals = append(newGrammar.nonTerminals, head)
	}

	// Recorrer la nueva gramática para identificar los terminales y no terminales
	for _, productions := range newGrammar.productions {
		for _, production := range productions {
			for _, symbol := range production {
				if symbol.isTerminal {
					// Añadir a la lista de terminales si no está ya presente
					if !containsSymbol(newGrammar.terminals, symbol) {
						newGrammar.terminals = append(newGrammar.terminals, symbol)
					}
				} else {
					// Añadir a la lista de no terminales si no está ya presente
					if !containsSymbol(newGrammar.nonTerminals, symbol) {
						newGrammar.nonTerminals = append(newGrammar.nonTerminals, symbol)
					}
				}
			}
		}
	}

	// Retornar la nueva gramática sin símbolos no alcanzables
	return newGrammar
}

func RemoveUselessSymbols(originalGrammar *Grammar) *Grammar {
	grammarWithoutGeneratingSymbols := RemoveNonGeneratingSymbols(originalGrammar)
	grammarWithoutReachableSymbos := RemoveNonReachableSymbols(grammarWithoutGeneratingSymbols)
	return grammarWithoutReachableSymbos
}
