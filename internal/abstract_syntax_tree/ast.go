package abstract_syntax_tree

import (
	"github.com/DanielRasho/Computation-Theory/internal/shuntingyard"
)

/*
BuildAST construye un AST a partir de una lista de símbolos en notación postfix.
Parámetros:
  - postfixSymbols: La expresión en notación postfix como una lista de símbolos (Symbol).

Retorno:
  - Un nodo (Node) que representa la raíz del AST construido a partir de la expresión postfix.

Panic:
 1. Si la expresión postfix es inválida, no está balanceada o en el stack hay menos símbolos de los que necesita un operador.
 2. Resultado del stack final no es un solo nodo (tal que la cantidad de operadores relacionados es incorrecta y faltan o sobran símbolos).
*/
func BuildAST(postfixSymbols []shuntingyard.Symbol) Node {
	var stack []Node

	// Recorrer toda la lista de símbolos en notación postfix
	for _, symbol := range postfixSymbols {

		// Verifica si el símbolo es un operador
		if operator, isOperator := symbol.(*shuntingyard.Operator); isOperator {

			// Obtener la cantidad de símbolos que necesita el operador
			operandCount := operator.GetOperands()
			if len(stack) < operandCount {
				panic("Expresión postfix inválida: falta operando")
			}

			// Añadir los símbolos que necesita el operador a operands
			operands := make([]Node, operandCount)
			for i := range operands {
				operands[i] = stack[len(stack)-1] // Agregar el valor a operands
				stack = stack[:len(stack)-1]      // Eliminar ese operando del stack
			}

			// Invierte el orden de los operandos de operands para mantener el orden correcto
			for i, j := 0, len(operands)-1; i < j; i, j = i+1, j-1 {
				operands[i], operands[j] = operands[j], operands[i]
			}
			// Crear un nodo operador con los operandos
			node := NewOperatorNode(operator.String(), operands)
			stack = append(stack, node)

		} else {
			// Si no es un operador, es un carácter (Symbol) y se añade al stack
			node := NewCharacterNode(symbol.String())
			stack = append(stack, node)
		}
	}

	if len(stack) != 1 {
		panic("Expresión postfix inválida: el resultado final no es un solo nodo")
	}
	return stack[0]
}
