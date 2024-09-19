/*
Package abstract_syntax_tree proporciona una implementación para construir y manejar un Árbol de Sintaxis Abstracta (AST) a partir de expresiones en notación postfix.
*/

package abstract_syntax_tree

import (
	"strings"

	shuntingyard "github.com/DanielRasho/Computation-Theory/internal/shuntingyard"
)

/*
BuildAST construye un AST a partir de una expresión en notación postfix.

Parámetros:
  - postfix: La expresión en notación postfix como una cadena.

Retorno:
  - Un nodo (Node) que representa la raíz del AST construido a partir de la expresión postfix.

Panic:
 1. Si la expresión postfix es inválida, no está balanceada o en el stack hay menos símbolos de los que necesita un operador.
 2. Resultado del stack final no es un solo operador (tal que la cantidad de operadores relacionados es incorrecta y faltan o sobran caracteres).
*/
func BuildAST(postfix string) Node {
	var stack []Node
	isEscaped := false // Variable para rastrear si el carácter actual está escapado

	// fmt.Println("Expresión postfix:", postfix)

	// Recorrer todo el postfix
	for _, char := range strings.Split(postfix, "") {
		// fmt.Printf("Procesando carácter %d: '%s'\n", i, char)

		if isEscaped {
			// Si el carácter está escapado, lo añadimos al stack
			node := NewCharacterNode(char)
			stack = append(stack, node)
			isEscaped = false // Reiniciar el estado de escape
			// fmt.Printf("Añadiendo carácter escapado '%s' al stack\n", char)
			continue
		}

		if char == "\\" {
			// Si encontramos un backslash, marcar el siguiente carácter como escapado
			isEscaped = true
			// fmt.Printf("Carácter de escape encontrado, próximo carácter será escapado\n")
			continue
		}

		// Verifica si el carácter es un operador
		if symbol, isOperator := shuntingyard.OPERATORS[char].(*shuntingyard.Operator); isOperator {
			// fmt.Printf("'%s' es un operador\n", char)

			// Obtener la cantidad de símbolos que necesita el operador
			operandCount := symbol.GetOperands()
			// fmt.Printf("El operador '%s' necesita %d operandos\n", char, operandCount)

			if len(stack) < operandCount {
				panic("Expresión postfix inválida: falta operando")
			}

			// Añadir los símbolos que necesita el operador a operands
			operands := make([]Node, operandCount)
			for i := range operands {
				operands[i] = stack[len(stack)-1] // Agregar el valor a operands
				// fmt.Printf("Obteniendo operando '%s' del stack\n", operands[i])
				stack = stack[:len(stack)-1] // Eliminar ese operando del stack
			}

			// Invierte el orden de los operandos de operands para mantener el orden correcto
			// fmt.Println("Invirtiendo el orden de los operandos")
			for i, j := 0, len(operands)-1; i < j; i, j = i+1, j-1 {
				operands[i], operands[j] = operands[j], operands[i]
			}

			node := NewOperatorNode(char, operands)
			// fmt.Printf("Creando nodo operador '%s' con operandos\n", char)
			stack = append(stack, node)
		} else {
			// Si no es un operador, se trata de un carácter y se añade al stack
			node := NewCharacterNode(char)
			// fmt.Printf("Añadiendo carácter '%s' al stack\n", char)
			stack = append(stack, node)
		}
	}

	if len(stack) != 1 {
		panic("Expresión postfix inválida: el resultado final no es un solo nodo")
	}

	// fmt.Println("Construcción del AST completada exitosamente")
	return stack[0]
}
