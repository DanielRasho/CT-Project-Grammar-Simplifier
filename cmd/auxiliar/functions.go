/**
 * Funciones auxiliares para el manejo de expresiones regulares y simulación de AFN.
 */

package auxiliar

import (
	"fmt"
	"strings"

	io "github.com/DanielRasho/TC-1-ShuntingYard/internal/IO"
	ast "github.com/DanielRasho/TC-1-ShuntingYard/internal/abstract_syntax_tree"
	dfaAutomata "github.com/DanielRasho/TC-1-ShuntingYard/internal/dfa"
	nfaAutomata "github.com/DanielRasho/TC-1-ShuntingYard/internal/nfa"
	runner "github.com/DanielRasho/TC-1-ShuntingYard/internal/runner_simulation"
	shuttingyard "github.com/DanielRasho/TC-1-ShuntingYard/internal/shuntingyard"
)

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
func PrintAllResults(index int, results []RegexResult) {
	if index < 0 || index >= len(results) {
		fmt.Println("Índice fuera de rango")
		return
	}

	result := results[index]

	fmt.Printf("==================================\n")
	fmt.Printf("| RESULTADOS PARA LA POSICIÓN %d |\n", index)
	fmt.Printf("==================================\n")

	// Imprime la expresión regular leída
	fmt.Printf("\nExpresión regular leída %d: %s\n", index+1, result.OriginalRegex)

	// Imprime el postfix
	fmt.Printf("\nPostfix: %s\n", result.Postfix)

	// Imprime el AST
	fmt.Println("\nEl AST resultante es:")
	PrintASTTree(result.AST, 0)

	// Imprime el NFA
	fmt.Println("\nEl NFA resultante es:")
	PrintNFA(result.NFA)

	// Imprime el DFA
	fmt.Println("\nEl DFA resultante es:")
	PrintDFA(result.DFA)

	// Imprime el DFA minimizado
	fmt.Println("\nEl DFA minimizado resultante es:")
	PrintDFA(result.DFA_M)
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
func PrintNFA(nfa *nfaAutomata.NFA) {
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

/*
PrintDFA imprime la estructura del DFA, mostrando el estado inicial, los estados finales,
y todas las transiciones entre estados, considerando que cada estado del DFA puede ser un conjunto de estados del NFA.

Parámetros:
  - dfa: Un puntero al DFA que se desea imprimir.

Retorno: Ninguno.
*/
func PrintDFA(dfa *dfaAutomata.DFA) {
	// Check if the DFA has any states
	if len(dfa.States) == 0 {
		panic("DFA has no states defined.")
	}
	if dfa.StartState == nil {
		panic("NO DFA start states defined.")
	}
	fmt.Printf("Estado inicial: %s\n", dfa.StartState.Name)
	fmt.Println("Estados finales:")
	for _, state := range dfa.States {
		if state.IsFinal {
			fmt.Printf("  %s\n", state.Name)
		}
	}

	fmt.Println("Transiciones:")
	for fromState, transitions := range dfa.Transitions {
		for symbol, toState := range transitions {
			fmt.Printf("  Desde: %s -> Hasta: [%s] con símbolo: %s\n", fromState.Name, toState.Name, symbol)
		}
	}
}

/*
InteractiveRegexSimulation es una función que permite al usuario interactuar con el programa
para convertir expresiones regulares a notación postfix, construir un AFN (Autómata Finito No Determinista)
y luego un AFD (Autómata Finito Determinista) a partir del AFN. Además, permite simular el AFN con una cadena
de entrada proporcionada por el usuario para verificar si pertenece al lenguaje definido por la expresión regular.

El proceso incluye los siguientes pasos:
 1. Solicitar al usuario una expresión regular.
 2. Convertir la expresión regular a notación postfix usando el algoritmo Shunting Yard.
 3. Construir un AST (Árbol Sintáctico Abstracto) a partir de la notación postfix.
 4. Construir un AFN a partir del AST.
 5. Convertir el AFN a un AFD.
 6. Renderizar y guardar la imagen del AFN generado.
 7. Solicitar al usuario una cadena para evaluar contra el AFN.
 8. Simular el AFN con la cadena proporcionada y mostrar el resultado de la simulación.

Si el usuario ingresa "0" como expresión regular, la función terminará la ejecución y saldrá del bucle.

Parámetros: Ninguno.

Retorno: Ninguno.
*/
func MyRegex() {
	for {
		fmt.Print("\n➡️  Ingresa una nueva expresión regular (utiliza ε para cadena vacía) o '0' para salir: ")
		var newRegex string
		fmt.Scanln(&newRegex)

		// Salir si el usuario ingresa "0"
		if newRegex == "0" {
			fmt.Println("\n🚪 Saliendo del programa... 🚪")
			break
		}

		// Convierte la expresión regular a postfix usando Shunting Yard
		postfix, _, _ := shuttingyard.RegexToPostfix(newRegex, false)
		// Construye el AST a partir del postfix
		root := ast.BuildAST(postfix)
		// Construye el AFN a partir del AST
		nfa := nfaAutomata.BuildNFA(root)
		// Construye el AFD
		dfa := dfaAutomata.BuildDFA(nfa)
		//	 Minimizar el AFD
		dfaM := dfaAutomata.MinimizeDFA(dfa)

		// Renderiza el NFA
		nfaFilename := fmt.Sprintf("./graphs/NFA/nfa_%s.png", newRegex)
		err := nfaAutomata.RenderAFN(nfa, nfaFilename)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("\t🌄 Grafo NFA generado exitosamente como '%s'!\n", nfaFilename)
		}

		// Renderiza el DFA
		dfaFilename := fmt.Sprintf("./graphs/DFA/dfa_%s.png", newRegex)
		err = dfaAutomata.RenderDFA(dfa, dfaFilename)
		if err != nil {
			fmt.Println("Error rendering DFA:", err)
		} else {
			fmt.Printf("\t🌄 Grafo DFA generado exitosamente como '%s'!\n", dfaFilename)
		}

		// Renderiza el DFA
		dfaMFilename := fmt.Sprintf("./graphs/DFA/dfaMIN_%s.png", newRegex)
		err = dfaAutomata.RenderDFA(dfaM, dfaMFilename)
		if err != nil {
			fmt.Println("Error rendering DFA:", err)
		} else {
			fmt.Printf("\t🌄 Grafo DFA generado exitosamente como '%s'!\n", dfaFilename)
		}

		// Simular el AFN con una cadena dada por el usuario
		fmt.Print("➡️  Ingresa la cadena a evaluar: ")
		var cadena string
		fmt.Scanln(&cadena)

		fmt.Printf("\t🤫 Susurro: escogiste la expresión regular '%s' para leer la cadena '%s'\n", newRegex, cadena)

		// Ejecutar la simulación del AFN y AFD con la cadena
		resultado_nfa := runner.RunnerNFA(nfa, cadena)
		resultado_dfa := runner.RunnerDFA(dfaM, cadena)

		// Mostrar el resultado de la simulación usando la nueva función
		RunnerSimulation(resultado_nfa, resultado_dfa, cadena, newRegex)
	}
}

/*
RunnerSimulation muestra el resultado de la simulación del AFN y del AFD con la cadena proporcionada por el usuario.
Dependiendo de si la cadena pertenece al lenguaje definido por la expresión regular o no, se imprime un mensaje correspondiente para cada uno.

Parámetros:
  - resultado_dfa: Resultado de la simulación del AFD, un booleano que indica si la cadena pertenece o no al lenguaje.
  - resultado_nfa: Resultado de la simulación del AFN, un booleano que indica si la cadena pertenece o no al lenguaje.
  - cadena: La cadena de entrada proporcionada por el usuario.
  - regex: La expresión regular utilizada para la simulación.

Retorno: Ninguno.
*/
func RunnerSimulation(resultado_dfa, resultado_nfa bool, cadena, regex string) {
	if resultado_nfa {
		fmt.Printf("✅ (AFN) Resultado de la simulación: la cadena '%s' ∈ L(%s)\n", cadena, regex)
	} else {
		fmt.Printf("❌ (AFN) Resultado de la simulación: la cadena '%s' ∉ L(%s)\n", cadena, regex)
	}

	if resultado_dfa {
		fmt.Printf("✅ (AFD) Resultado de la simulación: la cadena '%s' ∈ L(%s)\n", cadena, regex)
	} else {
		fmt.Printf("❌ (AFD) Resultado de la simulación: la cadena '%s' ∉ L(%s)\n", cadena, regex)
	}

	fmt.Println("\n-----------------------------------------")
}

/*
ProcessRegexFromFile lee expresiones regulares desde un archivo de texto, las convierte en postfix,
construye el AST, genera el NFA y DFA, y finalmente renderiza las imágenes correspondientes para
cada expresión regular. Además, guarda los resultados de cada paso en una lista.

Parámetros:
  - filePath: Ruta del archivo de texto que contiene las expresiones regulares.

Retorno:
  - []RegexProcessResult: Lista de resultados que incluye la expresión regular original, su conversión a postfix,
    el AST generado, el NFA y el DFA.
  - error: Error en caso de que ocurra algún problema durante la lectura del archivo o el procesamiento de las expresiones.
*/
func RegexFile(filePath string) ([]RegexResult, error) {
	var results []RegexResult

	// Llama a la función de lectura de archivo
	lines, err := io.ReaderTXT(filePath)
	if err != nil {
		return nil, err
	}

	// Procesa cada línea leída del archivo
	for index, line := range lines {
		fmt.Printf("\nExpresión Regular: %s\n", line)

		// Convertir a postfix
		postfix, _, _ := shuttingyard.RegexToPostfix(line, false)

		// Construir el AST
		root := ast.BuildAST(postfix)

		// Construir el NFA
		nfa := nfaAutomata.BuildNFA(root)

		// Convertir a DFA
		dfa := dfaAutomata.BuildDFA(nfa)

		// Minimizar DFA
		dfaM := dfaAutomata.MinimizeDFA(dfa)

		// Renderizar el NFA
		err := nfaAutomata.RenderAFN(nfa, fmt.Sprintf("./graphs/NFA/nfa_%d_%s.png", index, line))
		if err != nil {
			fmt.Println("Error renderizado de NFA:", err)
		}

		// Renderizar el DFA
		/*
			err = dfaAutomata.RenderDFA(dfaM, fmt.Sprintf("./graphs/DFA/dfa_%d_%s.png", index, line))
			if err != nil {
				fmt.Println("Error rendereizado de DFA:", err)
			}
		*/

		// Agregar el resultado al listado
		results = append(results, RegexResult{
			OriginalRegex: line,
			Postfix:       postfix,
			AST:           root,
			NFA:           nfa,
			DFA:           dfa,
			DFA_M:         dfaM,
		})
	}

	// Imprimir resultados para la primera expresión regular
	// PrintAllResults(0, results)

	return results, nil
}

/*
RegexProcessResult contiene los resultados del procesamiento de una expresión regular.

Campos:
  - OriginalRegex: La expresión regular original leída del archivo.
  - Postfix: La representación en postfix de la expresión regular.
  - AST: El árbol sintáctico abstracto (AST) construido a partir de la expresión en postfix.
  - NFA: El autómata finito no determinista (NFA) generado a partir del AST.
  - DFA: El autómata finito determinista (DFA) convertido desde el NFA.
*/
type RegexResult struct {
	OriginalRegex string
	Postfix       string
	AST           ast.Node
	NFA           *nfaAutomata.NFA
	DFA           *dfaAutomata.DFA
	DFA_M         *dfaAutomata.DFA
}

/*
MenuForSimulationFromResults muestra un menú con las expresiones regulares procesadas y permite al usuario
seleccionar una para simular su NFA y DFA. Luego, ejecuta la simulación y muestra los resultados.

Parámetros:
  - results: Lista de resultados procesados que incluye las expresiones regulares, el postfix, el AST, el NFA y el DFA.

Retorno:
  - Ninguno.
*/
func MenuRegexFile(results []RegexResult) {
	// Mostrar las expresiones regulares procesadas
	fmt.Println("\n🔍 Selecciona una expresión regular para simular:")
	for i, result := range results {
		fmt.Printf("%d. %s\n", i+1, result.OriginalRegex)
	}

	// Solicitar al usuario seleccionar una opción
	fmt.Print("➡️  Ingresa el número de la expresión regular que deseas simular (o '0' para salir): ")
	var choice int
	fmt.Scanln(&choice)

	// Salir si el usuario ingresa "0"
	if choice == 0 {
		fmt.Println("\n🚪 Saliendo del menú... 🚪")
		return
	}

	// Validar la selección del usuario
	if choice < 1 || choice > len(results) {
		fmt.Println("Opción inválida. Por favor selecciona un número válido.")
		return
	}

	// Obtener el resultado de la expresión regular seleccionada
	selectedResult := results[choice-1]

	// Simular el NFA y DFA con la expresión regular seleccionada
	fmt.Print("➡️  Ingresa la cadena a evaluar: ")
	var cadena string
	fmt.Scanln(&cadena)

	fmt.Printf("\t🤫 Susurro: escogiste la expresión regular '%s' para leer la cadena '%s'\n", selectedResult.OriginalRegex, cadena)

	// Ejecutar la simulación del NFA y DFA con la cadena
	resultado_nfa := runner.RunnerNFA(selectedResult.NFA, cadena)
	resultado_dfa := runner.RunnerDFA(selectedResult.DFA_M, cadena)

	// Mostrar el resultado de la simulación usando la función RunnerSimulation
	RunnerSimulation(resultado_dfa, resultado_nfa, cadena, selectedResult.OriginalRegex)
}
