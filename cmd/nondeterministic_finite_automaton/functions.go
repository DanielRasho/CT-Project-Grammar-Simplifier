/**
 * Funciones auxiliares para el manejo de expresiones regulares y simulación de AFN.
 */

package main

import (
	"fmt"
	"strings"

	ast "github.com/DanielRasho/TC-1-ShuntingYard/internal/abstract_syntax_tree"
	"github.com/DanielRasho/TC-1-ShuntingYard/internal/nfa"
)

/*
MenuNFATXT muestra un menú con las expresiones regulares cargadas desde un archivo,
permite al usuario seleccionar una, y luego solicita una cadena para simularla con el AFN correspondiente.
Parámetros:
- erList: Una lista de expresiones regulares (ER) leídas desde un archivo.
- nfaList: Una lista de AFNs generados a partir de las expresiones regulares.
Retorno: Ninguno.
*/
func MenuNFATXT(erList []string, nfaList []*nfa.NFA) {
	for {
		// Mostrar opciones disponibles
		fmt.Print("=====================================\n")
		fmt.Print("| EXPRESIONES REGULARES DISPONIBLES |\n")
		fmt.Print("=====================================\n")
		fmt.Println("0. Salir")
		for i := 0; i < len(erList); i++ {
			fmt.Println(fmt.Sprintf("%d. %s", i+1, erList[i]))
		}

		// Leer la selección del usuario
		fmt.Print("➡️ Escoge la expresión regular para leer la cadena: ")
		var respuesta_er int
		fmt.Scanln(&respuesta_er)

		// Salir si el usuario ingresa "0"
		if respuesta_er == 0 {
			fmt.Println("🚪 Saliendo del programa... 🚪")
			break
		}

		// Ajuste para el índice seleccionado
		respuesta_er -= 1

		// Validar si la opción está dentro del rango
		if respuesta_er < 0 || respuesta_er >= len(erList) {
			fmt.Println("Índice fuera de rango")
			continue
		}

		// Leer la cadena del usuario
		fmt.Print("➡️  Ingresa la cadena a evaluar: ")
		var cadena string
		fmt.Scanln(&cadena)
		fmt.Println("\t🤫 Susurro: escogiste la expresión regular '" + erList[respuesta_er] + "' para leer la cadena '" + cadena + "'")

		// Ejecutar la simulación del AFN con la cadena seleccionada
		resultado := nfa.RunnerNFA(nfaList[respuesta_er], cadena)

		// Mostrar el resultado de la simulación
		if resultado {
			fmt.Println("✅ Resultado de la simulación: la cadena '" + cadena + "' ∈ L(" + erList[respuesta_er] + ")\n")
		} else {
			fmt.Println("❌ Resultado de la simulación: la cadena '" + cadena + "' ∉ L(" + erList[respuesta_er] + ")\n")
		}
	}
}

/*
PrintAllResults muestra todos los resultados asociados a una expresión regular en particular,
incluyendo la expresión regular original, la notación postfix, el AST y el AFN.
Parámetros:
- index: Índice de la expresión regular en la lista.
- erList: Lista de expresiones regulares.
- postfixList: Lista de notaciones postfix de las expresiones regulares.
- astList: Lista de árboles de sintaxis abstracta (AST) generados a partir de las expresiones regulares.
- nfaList: Lista de AFNs generados a partir de los AST.
Retorno: Ninguno.
*/
func PrintAllResults(index int, erList []string, postfixList []string, astList []ast.Node, nfaList []*nfa.NFA) {
	if index < 0 || index >= len(erList) {
		fmt.Println("Índice fuera de rango")
		return
	}

	fmt.Printf("==================================\n")
	fmt.Printf("| RESULTADOS PARA LA POSICIÓN %d |\n", index)
	fmt.Printf("==================================\n")

	// Imprime la línea leída
	fmt.Printf("\nExpresión regular leída %d: %s\n", index+1, erList[index])

	// Imprime el postfix
	fmt.Printf("\nPostfix: %s\n", postfixList[index])

	// Imprime el AST
	fmt.Println("\nEl AST resultante es:")
	PrintASTTree(astList[index], 0)

	// Imprime el NFA
	fmt.Println("\nEl NFA resultante es:")
	PrintNFA(nfaList[index])
}

/*
PrintASTTree imprime el árbol de sintaxis abstracta (AST) de forma recursiva,
mostrando cada nodo y su nivel de profundidad en el árbol.
Parámetros:
- node: Nodo actual del AST.
- level: Nivel de profundidad actual en el árbol.
Retorno: Ninguno.
*/
func PrintASTTree(node ast.Node, level int) {
	// Imprime el nodo actual
	switch n := node.(type) {
	case *ast.CharacterNode:
		fmt.Printf("%sCharacterNode: %s\n", indent(level), n.Value)
	case *ast.OperatorNode:
		fmt.Printf("%sOperatorNode: %s\n", indent(level), n.Value)
		for _, operand := range n.GetOperands() {
			PrintASTTree(operand, level+1)
		}
	}
}

/*
PrintNFA imprime la estructura del AFN, mostrando el estado inicial, el estado final,
y todas las transiciones entre estados.

Parámetros:
  - nfa: Un puntero al AFN que se desea imprimir.

Retorno: Ninguno.
*/
func PrintNFA(nfa *nfa.NFA) {
	fmt.Printf("Estado inicial: %s\n", nfa.StartState.Name)
	fmt.Printf("Estado final: %s\n", nfa.EndState.Name)
	fmt.Println("Transiciones:")
	for _, t := range nfa.Transitions {
		toStates := make([]string, len(t.To))
		for i, s := range t.To {
			toStates[i] = s.Name
		}
		fmt.Printf("  Desde: %s -> Hasta: [%s] con símbolo: %s\n", t.From.Name, strings.Join(toStates, ", "), t.Symbol)
	}
}

/*
indent genera un string de indentación basado en el nivel de profundidad,
útil para formatear la salida de árboles o estructuras anidadas.
Parámetros:
- level: Nivel de profundidad para el cual se desea generar la indentación.
Retorno:
- Un string que representa la indentación.
*/
func indent(level int) string {
	return strings.Repeat("  ", level)
}
