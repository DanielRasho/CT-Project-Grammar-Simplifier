package grammar

// Función para determinar si una cadena es aceptada por una gramática en forma normal de Chomsky (CNF).
func CYKParse(grammar *Grammar, cadena string, initialSymbol Symbol) bool {
	lista_cadena := []rune(cadena)

	// Crear una matriz vacía de tamaño len(lista_cadena) x len(lista_cadena)
	matrixT := make([][][]string, len(lista_cadena))
	for i := range matrixT {
		matrixT[i] = make([][]string, len(lista_cadena))
	}

	// Llenar la matriz según la fila
	for i := range matrixT {
		for j := range lista_cadena {
			if i == 0 { // Llenar la fila 0 con los heads que producen directamente los terminales
				listado := FindHeadsProducingTerminal(grammar, string(lista_cadena[j]))
				matrixT[i][j] = removeDuplicatesString(listado) // Eliminar duplicados y asignar a la matriz

			} else if i == 1 && j < len(lista_cadena)-1 {
				values1 := matrixT[0][j]   // Obtener los valores de la columna actual y la primera fila
				values2 := matrixT[0][j+1] // Obtener los valores de la siguiente columna siguiente de la primera fila
				listado := []string{}      // Crear un slice para almacenar los resultados

				// Combinar valores de values1 y values2
				for _, v1 := range values1 {
					for _, v2 := range values2 {
						// Llamar a FindHeadsProducingNonTerminals para el par de valores (v1, v2)
						listado = append(listado, FindHeadsProducingNonTerminals(grammar, Symbol{Value: v1, IsTerminal: false}, Symbol{Value: v2, IsTerminal: false})...)
					}
				}

				// Almacenar el resultado en la fila 1, columna j
				matrixT[i][j] = removeDuplicatesString(listado) // Eliminar duplicados

			} else if i > 1 && j < len(lista_cadena)-i { // Para las filas superiores evitar la ultima columna
				listado := []string{} // Crear un slice para almacenar los resultados

				// Probar todas las posibles particiones de la subcadena
				for k := 0; k < i; k++ {
					values1 := matrixT[k][j]         // Valores de la partición izquierda
					values2 := matrixT[i-k-1][j+k+1] // Valores de la partición derecha

					// Combinar valores de values1 y values2
					for _, v1 := range values1 {
						for _, v2 := range values2 {
							// Llamar a FindHeadsProducingNonTerminals para el par de valores (v1, v2)
							listado = append(listado, FindHeadsProducingNonTerminals(grammar, Symbol{Value: v1, IsTerminal: false}, Symbol{Value: v2, IsTerminal: false})...)
						}
					}
				}

				// Almacenar el resultado en la posición correspondiente
				matrixT[i][j] = removeDuplicatesString(listado)

			} else {
				// Para las filas restantes, sino tiene nada dejar vacío
				matrixT[i][j] = []string{}
			}
		}

		// Imprimir el estado actual de la matriz después de completar la fila 'i'
		// fmt.Printf("Matriz después de completar la fila %d:\n", i)
		// for fi, fila := range matrixT {
		// 	fmt.Printf("Fila %d:\n", fi)
		// 	for j, heads := range fila {
		// 		fmt.Printf("  Columna %d: %v\n", j, heads)
		// 	}
		// }
	}

	lastCell := matrixT[len(lista_cadena)-1][0]

	// Si el símbolo inicial está en la última celda, entonces la cadena es aceptada
	for _, head := range lastCell {
		if head == initialSymbol.Value {
			return true
		}
	}

	return false
}
