/*
Package abstract_syntax_tree proporciona una implementación para construir y manejar un Árbol de Sintaxis Abstracta (AST) a partir de expresiones en notación postfix.
*/

package abstract_syntax_tree

import (
	"strings"

	"github.com/DanielRasho/TC-1-ShuntingYard/internal/shuntingyard"
)

/*
BuildAST construye un AST a partir de una expresión en notación postfix.
Parámetros:
  - postfix: La expresión en notación postfix como una cadena.

Retorno:
  - Un nodo (Node) que representa la raíz del AST construido a partir de la expresión postfix.

Panic:
 1. Si la expresión postfix es inválida, no está balanceada o en el stack hay menos simbolos de los que necesita un operador.
 2. Resultado del stack final no es un solo operador (tal que la cantidad de operadores relacionados es incorrecta y faltan o sobran caracteres)
*/
func BuildAST(postfix string) Node {
	var stack []Node

	// Recorrer todo el postfix
	for _, char := range strings.Split(postfix, "") {

		// Verifica si el carácter es un operador
		if symbol, isOperator := shuntingyard.OPERATORS[char].(*shuntingyard.Operator); isOperator {

			// Obtener la cantidad de simbolos que necesita el operador
			operandCount := symbol.GetOperands()
			if len(stack) < operandCount {
				panic("Expresión postfix inválida: falta operando")
			}

			// Añadir los simbolos que necesita el operador a operands
			operands := make([]Node, operandCount)
			for i := range operands {
				operands[i] = stack[len(stack)-1] // Agregar el valor a operands
				stack = stack[:len(stack)-1]      // Eliminar ese operando del stack
			}

			// Invierte el orden de los operandos de operands para mantener el orden correcto
			for i, j := 0, len(operands)-1; i < j; i, j = i+1, j-1 {
				operands[i], operands[j] = operands[j], operands[i]
			}
			node := NewOperatorNode(char, operands)
			stack = append(stack, node)
		} else {
			// Si no es un operador, se trata de un carácter y se añade al stack
			node := NewCharacterNode(char)
			stack = append(stack, node)
		}
	}

	if len(stack) != 1 {
		panic("Expresión postfix inválida: el resultado final no es un solo nodo")
	}
	return stack[0]
}
