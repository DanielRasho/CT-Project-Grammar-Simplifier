package grammar

import "fmt"

// Inicializa las parejas base (A, A) para cada no terminal A y las producciones unarias en su mismo conjunto
func initializeUnaryPairs(originalGrammar *Grammar) map[string][]string {
	nonTerminals := getNonTerminals(originalGrammar)
	unaryBase := make(map[string][]string)

	// Inicializa las parejas unarias y extiende con producciones unarias
	for key := range nonTerminals {
		unaryBase[key] = []string{key} // Cada no terminal se inicializa con su propia pareja
	}

	for head, productions := range *originalGrammar {
		for _, production := range productions {
			if isUnary(production, nonTerminals) {
				unaryBase[head] = append(unaryBase[head], production)
			}
		}
	}

	// Elimina duplicados de las producciones unarias
	for key, productions := range unaryBase {
		unaryBase[key] = RemoveDuplicates(productions)
	}

	return unaryBase
}

// Encuentra todas las parejas unarias
func findUnaryPairs(unaryBase map[string][]string) map[string][]string {
	// Crear un nuevo mapa para almacenar las parejas unarias extendidas
	unaryPairs := make(map[string][]string)

	// Inicializar con las parejas unarias originales
	for key, value := range unaryBase {
		unaryPairs[key] = value
	}

	// Por cada clave, expandir sus valores
	for key := range unaryPairs {
		expanded := true // Bandera para indicar si se ha realizado una expansión
		for expanded {
			expanded = false // Resetear la bandera

			// Obtener los valores actuales
			currentValues := unaryPairs[key]
			if len(currentValues) > 0 {
				lastValue := currentValues[len(currentValues)-1] // Obtener el último valor

				// Por cada llave, agregar valores
				for skey, svalue := range unaryPairs {
					if lastValue == skey {
						// Agregar todos los valores de skey a la key, asegurándose de no duplicar
						for _, sv := range svalue {
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
		// Inicializa las producciones para el no terminal actual
		var productions []string

		// Para cada cabeza de la gramática original
		for head := range *originalGrammar {
			if key == head {
				// Por cada value en las producciones unarias
				for _, value := range values {
					// Eliminar ese value en los values de key
					// Y obtener las producciones de value eliminado
					// Siempre que no sea un unario por medio de isUnary
					if !isUnary(value, nonTerminals) {
						// Si no es unario, añade las producciones a la lista
						productions = append(productions, value)
					} else {
						// Agregar las producciones de value a key
						if newProductions, exists := (*originalGrammar)[value]; exists {
							for _, newProduction := range newProductions {
								// Verifica si la nueva producción no es un unario antes de añadir
								if !isUnary(newProduction, nonTerminals) {
									productions = append(productions, newProduction)
								}
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

// imprime el resultado de parejas unarias
func printResultUnarPairs(unaryPairs map[string][]string) {
	// Imprimir el contenido de unaryPairs
	for key, values := range unaryPairs {
		// Imprimir la clave
		fmt.Print(key + ":")
		// Imprimir los valores asociados a la clave
		for _, value := range values {
			fmt.Print(value + "-")
		}
		fmt.Println() // Nueva línea al final de cada clave
	}
}

// printGrammar imprime la gramática de manera legible
func printGrammar(g *Grammar) {
	for key, productions := range *g {
		fmt.Printf("%s -> ", key)
		for i, production := range productions {
			if i > 0 {
				fmt.Print(" | ") // separador entre producciones
			}
			fmt.Print(production)
		}
		fmt.Println() // nueva línea después de cada no terminal
	}
}
