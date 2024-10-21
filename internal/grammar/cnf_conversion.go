package grammar

/*
Reemplazar los cuerpos de las producciones de longitud mayor o igual a 2 que contienen terminales, creando nuevos no terminales para cada terminal.
*/
func CNFTerminalSubstitution(originalGrammar *Grammar) *Grammar {
	// Paso 1: Crear una copia de la gramática original para añadir las nuevas producciones
	newGrammar := &Grammar{
		terminals:    originalGrammar.terminals,
		nonTerminals: originalGrammar.nonTerminals,
		Productions:  make(map[Symbol][][]Symbol),
	}

	// Paso 2: Crear un nuevo no terminal por cada terminal que aparece en producciones de longitud >= 2
	terminalToNonTerminal := make(map[Symbol]Symbol)
	for _, terminal := range originalGrammar.terminals {
		// Crear un nuevo símbolo no terminal que reemplazará al terminal
		newNonTerminal := Symbol{isTerminal: false, value: terminal.value, id: terminal.id + 1}
		terminalToNonTerminal[terminal] = newNonTerminal

		// Añadir el nuevo no terminal a la lista de no terminales
		newGrammar.nonTerminals = append(newGrammar.nonTerminals, newNonTerminal)

		// Añadir la producción de nuevo no terminal -> terminal
		newGrammar.Productions[newNonTerminal] = [][]Symbol{{terminal}}
	}

	// Paso 3: Modificar las producciones de la gramática original, reemplazando terminales por no terminales
	for head, productions := range originalGrammar.Productions {
		newProductions := make([][]Symbol, 0)

		for _, production := range productions {
			// Si la producción tiene longitud >= 2 y contiene terminales, reemplazar terminales
			if len(production) >= 2 {
				modifiedProduction := replaceTerminalsWithNonTerminals(production, terminalToNonTerminal)
				newProductions = append(newProductions, modifiedProduction)
			} else {
				// Mantener las producciones que ya están en forma normal
				newProductions = append(newProductions, production)
			}
		}

		// Actualizar las producciones en la nueva gramática
		newGrammar.Productions[head] = newProductions
	}

	return newGrammar
}

/*
Reemplazar los terminales en una producción con sus no terminales correspondientes.
*/
func replaceTerminalsWithNonTerminals(production []Symbol, terminalToNonTerminal map[Symbol]Symbol) []Symbol {
	newProduction := make([]Symbol, len(production))
	copy(newProduction, production)

	for i, symbol := range newProduction {
		// Si el símbolo es un terminal y tiene un equivalente no terminal, reemplazarlo
		if replacement, exists := terminalToNonTerminal[symbol]; exists {
			newProduction[i] = replacement
		}
	}

	return newProduction
}

/*
Divide producciones largas (de longitud > 2) en producciones binarias.
*/
func CNFSplitLargeProductions(originalGrammar *Grammar) *Grammar {
	// Crear una nueva gramática para almacenar las producciones resultantes
	newGrammar := &Grammar{
		terminals:    originalGrammar.terminals,
		nonTerminals: originalGrammar.nonTerminals,
		Productions:  make(map[Symbol][][]Symbol),
	}

	// Iterar sobre las producciones de la gramática original
	for head, productions := range originalGrammar.Productions {
		for _, production := range productions {
			// Si la producción tiene más de 2 símbolos, se debe dividir
			for len(production) > 2 {
				// Obtener los dos últimos símbolos de la producción
				lastSymbol1 := production[len(production)-1]
				lastSymbol2 := production[len(production)-2]

				// Remover los dos últimos símbolos de la producción original
				production = production[:len(production)-2]

				// Crear un nuevo símbolo no terminal que represente a estos dos símbolos
				newSymbol := Symbol{
					isTerminal: false,
					value:      lastSymbol2.value + "_" + lastSymbol1.value, // Nombre combinado de los símbolos
					id:         0,                                           // id 0 porque son nuevas producciones que no derivan de nada
				}

				// Añadir el nuevo no terminal a la lista de no terminales si no está presente
				newGrammar.nonTerminals = append(newGrammar.nonTerminals, newSymbol)

				// Añadir la nueva producción a la gramática: newSymbol -> lastSymbol2 lastSymbol1
				newGrammar.Productions[newSymbol] = [][]Symbol{{lastSymbol2, lastSymbol1}}

				// Reemplazar los dos últimos símbolos por el nuevo símbolo en la producción actual
				production = append(production, newSymbol)
			}

			// Añadir la producción (de longitud 2 o menos) a la nueva gramática
			newGrammar.Productions[head] = append(newGrammar.Productions[head], production)
		}
	}

	return newGrammar
}
