package grammar

/*
Inicializa las parejas base (A, A) para cada no terminal A y las producciones unarias en su mismo conjunto
*/
func InitializeUnaryPairs(originalGrammar *Grammar) map[Symbol][]Symbol {

	nonTerminals := originalGrammar.NonTerminals
	unaryBase := make(map[Symbol][]Symbol)

	// Se crean las bases unarias por cada encabezado de producciones de las gramáticas
	for i := range nonTerminals {
		key := nonTerminals[i]
		unaryBase[key] = []Symbol{key} // Cada no terminal se inicializa con su propia pareja
	}

	// Añadir los unarios de una misma producción a unaryBase correspondiente a su encabezado
	for head, productions := range originalGrammar.Productions {
		for _, production := range productions {
			// Si la producción es unaria
			if isUnary(production, nonTerminals) {
				// Si la producción no está en unaryBase, añadirla
				if !containsSymbol(unaryBase[head], production[0]) {
					unaryBase[head] = append(unaryBase[head], production[0])
				}
			}
		}
	}

	return unaryBase
}

/*
Encuentra todas las parejas unarias de toda la gramática
*/
func FindUnaryPairs(unaryBase map[Symbol][]Symbol) map[Symbol][]Symbol {
	// Crear un nuevo mapa para almacenar las parejas unarias extendidas
	unaryPairs := make(map[Symbol][]Symbol)

	// Inicializar con las parejas unarias originales
	for key, value := range unaryBase {
		unaryPairs[key] = value
	}

	// Por cada clave, encontrar todas las producciones unarias posibles a partir de un no terminal
	for key := range unaryPairs {
		expanded := true // Bandera para indicar si se ha realizado una expansión

		// Seguir buscando producciones hasta que ya no existan producciones por añadir
		for expanded {
			expanded = false

			// Obtener los valores actuales
			currentValues := unaryPairs[key]
			if len(currentValues) > 0 {
				lastValue := currentValues[len(currentValues)-1] // Comenzar a partir de la última producción

				// Obtener los sub_valores de la sub_llave buscada,
				// (último valor del listado del encabezado que se lee)
				for sub_key, sub_value := range unaryPairs {
					if lastValue == sub_key {
						// Agregar todos los valores de la sub_llave a la llave original, asegurándose de no duplicar
						for _, sv := range sub_value {
							// Si los valores actuales no contienen los sub_valores añadirlo
							if !containsSymbol(unaryPairs[key], sv) {
								unaryPairs[key] = append(unaryPairs[key], sv)
								expanded = true // Si se agregó un nuevo elemento, se establece la bandera
							}
						}
					}
				}
			}
		}
	}

	// Retornar el nuevo mapa de unaryPairs
	return unaryPairs
}

/*
Elimina las producciones unarias y ajusta la gramática
*/
func RemoveUnaryProductions(originalGrammar *Grammar, nonTerminals []Symbol) *Grammar {
	unaryBase := InitializeUnaryPairs(originalGrammar)
	unaryPairs := FindUnaryPairs(unaryBase)

	newGrammar := &Grammar{
		terminals:    originalGrammar.terminals,
		NonTerminals: originalGrammar.NonTerminals,
		Productions:  make(map[Symbol][][]Symbol),
	}

	// Iterar sobre cada no terminal en unaryPairs
	for key, values := range unaryPairs {
		var productions [][]Symbol

		// Recorrer los no terminales en la gramática original
		for head := range originalGrammar.Productions {
			if key == head {
				// Para cada símbolo value en los pares unarios de la clave
				for _, value := range values {

					// Si el valor de unaryPairs existe como encabezado en la gramática original
					if newProductions, exists := originalGrammar.Productions[value]; exists {
						// Por cada producción de ese encabezado
						for _, newProduction := range newProductions {
							// Si no es una producción unaria y no está ya en la lista, añadirla a las producciones
							if !isUnary(newProduction, nonTerminals) && !containsProduction(productions, newProduction) {
								productions = append(productions, newProduction)
							}
						}
					}
				}
			}
		}

		// Asignar las producciones no unarias a la nueva gramática
		newGrammar.Productions[key] = productions
	}

	return newGrammar
}

/*
Comprueba si una producción es unaria (solo contiene un no terminal)
*/
func isUnary(production []Symbol, nonTerminals []Symbol) bool {
	// La producción debe tener exactamente un símbolo para ser unaria
	if len(production) != 1 {
		return false
	}

	// Obtener el único símbolo de la producción
	symbol := production[0]

	// Verificar si el símbolo es un no terminal
	if !symbol.isTerminal && containsSymbol(nonTerminals, symbol) {
		return true
	}

	// Si no es un no terminal, no es una producción unaria
	return false
}
