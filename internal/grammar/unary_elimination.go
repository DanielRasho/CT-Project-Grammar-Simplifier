package grammar

/*
// Inicializa las parejas base (A, A) para cada no terminal A y las producciones unarias en su mismo conjunto
func initializeUnaryPairs(originalGrammar *Grammar) map[string][]string {
	nonTerminals := getNonTerminals(originalGrammar)
	unaryBase := make(map[string][]string)

	// Se crean las bases unarias por cada encabezado de producciones de las gramáticas
	for key := range nonTerminals {
		unaryBase[key] = []string{key} // Cada no terminal se inicializa con su propia pareja
	}

	// Añadir los unarios de una misma producción a unaryBase correspondiente a su encabezado
	for head, productions := range *originalGrammar {
		for _, production := range productions {
			// Si la producción es unaria
			if isUnary(production, nonTerminals) {
				// Si la producción no está en unaryBase, añadirla
				if !contains(unaryBase[head], production) {
					unaryBase[head] = append(unaryBase[head], production)
				}
			}
		}
	}

	return unaryBase
}

// Encuentra todas las parejas unarias de toda la gramática
func findUnaryPairs(unaryBase map[string][]string) map[string][]string {
	// Crear un nuevo mapa para almacenar las parejas unarias extendidas
	unaryPairs := make(map[string][]string)

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

							// Si los valores actuales no contienene los sub_valores añadirlo
							if !contains(unaryPairs[key], sv) {
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

// Elimina las producciones unarias y ajusta la gramática
func removeUnaryProductions(originalGrammar *Grammar, unaryPairs map[string][]string, nonTerminals map[string]struct{}) *Grammar {
	newGrammar := make(Grammar)

	// Hacer una nueva gramática y para el head de cada unaryPairs
	// traer las producciones no unarias
	for key, values := range unaryPairs {
		// Producciones para el no terminal actual
		var productions []string

		// Para cada cabeza de la gramática original
		for head := range *originalGrammar {
			if key == head {
				// Por cada value en las producciones unarias
				for _, value := range values {

					// Existe el value de unarypair en la gramatica original como encabezado
					if newProductions, exists := (*originalGrammar)[value]; exists {
						// Por cada producción del encabezado
						for _, newProduction := range newProductions {
							// Si no es unario añadirlo a las producciones
							if !isUnary(newProduction, nonTerminals) {
								productions = append(productions, newProduction)
							}
						}
					}
				}
			}
		}

		// Asigna las producciones no unarias a la nueva gramática
		newGrammar[key] = productions
	}

	return &newGrammar
}

// Comprueba si una producción es unaria (solo contiene un no terminal)
func isUnary(production string, nonTerminals map[string]struct{}) bool {
	// Comprobar si la producción es un no terminal en el mapa
	_, exists := nonTerminals[production]
	return exists
}
*/
