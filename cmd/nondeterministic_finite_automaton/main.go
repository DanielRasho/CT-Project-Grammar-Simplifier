package main

import (
	"fmt"

	io "github.com/DanielRasho/TC-1-ShuntingYard/internal/IO"
	ast "github.com/DanielRasho/TC-1-ShuntingYard/internal/abstract_syntax_tree"
	nfaAutomata "github.com/DanielRasho/TC-1-ShuntingYard/internal/nfa"
	runner "github.com/DanielRasho/TC-1-ShuntingYard/internal/runner_simulation"
	shuttingyard "github.com/DanielRasho/TC-1-ShuntingYard/internal/shuntingyard"
)

/**
 * main es la función principal del programa, encargada de gestionar la lógica de procesamiento
 * de expresiones regulares, convertirlas a postfix, construir el árbol de sintaxis abstracta (AST),
 * construir el autómata finito no determinista (AFN) y simular la aceptación de cadenas por los AFNs.
 *
 * El flujo del programa incluye:
 * 1. Cargar expresiones regulares desde un archivo.
 * 2. Convertir cada expresión regular a postfix, construir el AST y el AFN correspondiente.
 * 3. Permitir al usuario ingresar nuevas expresiones regulares y simular cadenas con ellas.
 * 4. Imprimir resultados detallados para una expresión regular específica.
 * 5. Graficar todos los AFNs generados (comentado en este ejemplo).
 *
 * No recibe parámetros ni devuelve valores, pues actúa directamente sobre la entrada/salida estándar.
 */
func main() {
	// Llama a la función de lectura de archivo
	lines, err := io.ReaderTXT("input_data/thompson.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Procesa cada línea leída del archivo
	for index, line := range lines {

		// Convierte la expresión regular a postfix usando Shunting Yard
		postfix, _, _ := shuttingyard.RegexToPostfix(line, false)

		// Construye el AST a partir del postfix
		root := ast.BuildAST(postfix)

		// Construye el AFN a partir del AST
		nfa := nfaAutomata.BuildNFA(root)

		// Render initial automatons
		err := nfaAutomata.RenderAFN(nfa, fmt.Sprintf("./graphs/nfa%d.png", index))
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("\t🌄 NFA image generated successfully!")
		}
	}

	fmt.Println("🔄 Simulador de expresiones regulares con AFN 🔄")

	// Bucle principal para pedir una nueva expresión regular y cadena
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

		// Renderiza el automata generado.
		err := nfaAutomata.RenderAFN(nfa, "./graphs/nfa.png")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("\t🌄 NFA image generated successfully!")
		}

		// Simular el AFN con una cadena dada por el usuario
		fmt.Print("➡️  Ingresa la cadena a evaluar: ")
		var cadena string
		fmt.Scanln(&cadena)

		fmt.Printf("\t🤫 Susurro: escogiste la expresión regular '%s' para leer la cadena '%s'\n", newRegex, cadena)

		// Ejecutar la simulación del AFN con la cadena
		resultado := runner.RunnerNFA(nfa, cadena)

		// Mostrar el resultado de la simulación
		if resultado {
			fmt.Printf("✅ Resultado de la simulación: la cadena '%s' ∈ L(%s)\n", cadena, newRegex)

		} else {
			fmt.Printf("❌ Resultado de la simulación: la cadena '%s' ∉ L(%s)\n", cadena, newRegex)
		}

		fmt.Println("\n-----------------------------------------")
	}
}
