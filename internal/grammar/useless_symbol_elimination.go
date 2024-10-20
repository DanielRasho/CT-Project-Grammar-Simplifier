package grammar

/*
// Función que encuentra los símbolos que generan cadenas de terminales.
func findGeneratingSymbols(originalGrammar *Grammar) (map[string]struct{}, map[string]struct{}) {
	nonTerminals := getNonTerminals(originalGrammar)
	terminals := getTerminals(originalGrammar, nonTerminals)

	generatingSymbols := make(map[string]struct{})
	nonGeneratingSymbols := make(map[string]struct{})

	changed := true
	for changed {
		changed = false

		// Revisar cada no terminal.
		for head, productions := range *originalGrammar {

			// Si ya existe en la lista de simbolos generadores, se salta
			if _, exists := generatingSymbols[head]; exists {
				continue
			}

			// Revisar cada producción del no terminal.
			for _, production := range productions {
				// Verificar si toda la producción genera (es decir, si está compuesta solo de terminales o no terminales generadores).
				if isComposedOf(append(getKeys(terminals), getKeys(generatingSymbols)...), production) {
					// Si genera, añadimos el no terminal al conjunto de generadores.
					generatingSymbols[head] = struct{}{}
					changed = true
					break
				}
			}
		}
	}

	// Paso 4: Identificar los no generadores.
	for head := range nonTerminals {
		if _, exists := generatingSymbols[head]; !exists {
			nonGeneratingSymbols[head] = struct{}{}
		}
	}

	return generatingSymbols, nonGeneratingSymbols
}

// Función que encuentra los símbolos alcanzables a partir de un símbolo inicial.
func findReachableSymbols(originalGrammar *Grammar, startSymbol string) (map[string]struct{}, map[string]struct{}) {
	nonTerminals := getNonTerminals(originalGrammar)

	reachableSymbols := make(map[string]struct{})
	unreachableSymbols := make(map[string]struct{})

	// Utilizamos una cola para procesar los símbolos alcanzables.
	queue := []string{startSymbol}

	// Mientras haya elementos en la cola.
	for len(queue) > 0 {
		// Extraemos el primer elemento de la cola.
		current := queue[0]
		queue = queue[1:]

		// Si ya hemos marcado el símbolo como alcanzable, continuamos.
		if _, exists := reachableSymbols[current]; exists {
			continue
		}

		// Marcamos el símbolo actual como alcanzable.
		reachableSymbols[current] = struct{}{}

		// Recorremos todas las producciones del símbolo actual.
		for _, production := range (*originalGrammar)[current] {
			// Usamos la función splitProduction para separar la producción concatenada en símbolos individuales.
			symbols := splitProduction(production)

			// Procesamos cada símbolo en la producción.
			for _, symbol := range symbols {
				// Verificamos si el símbolo es un no terminal.
				if _, isNonTerminal := nonTerminals[symbol]; isNonTerminal {
					// Si el símbolo aún no está en los alcanzables, lo añadimos a la cola.
					if _, alreadyQueued := reachableSymbols[symbol]; !alreadyQueued {
						queue = append(queue, symbol)
					}
				}
			}
		}
	}

	// Identificamos los no terminales que no son alcanzables.
	for nonTerminal := range nonTerminals {
		if _, exists := reachableSymbols[nonTerminal]; !exists {
			unreachableSymbols[nonTerminal] = struct{}{}
		}
	}

	return reachableSymbols, unreachableSymbols
}
*/
