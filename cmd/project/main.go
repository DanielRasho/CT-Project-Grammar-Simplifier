package main

import (
	"fmt"

	functions "github.com/DanielRasho/TC-1-ShuntingYard/cmd/auxiliar"
)

func main() {
	fmt.Println("🔄 Simulador de expresiones regulares con AFN y AFD 🔄")

	for {
		// Mostrar el menú principal
		fmt.Println("\n🌟 Menú Principal 🌟")
		fmt.Println("1. Ingresar una nueva expresión regular")
		fmt.Println("2. Procesar expresiones regulares desde un archivo")
		fmt.Println("0. Salir")
		fmt.Print("➡️  Selecciona una opción: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error leyendo la opción:", err)
			continue
		}

		switch choice {
		case 1:
			// Llamar a la función para ingresar una nueva expresión regular
			functions.MyRegex()

		case 2:
			// Procesar el archivo con expresiones regulares
			results, err := functions.RegexFile("input_data/test.txt")
			if err != nil {
				fmt.Println("Error procesando el archivo:", err)
				continue
			}

			// Mostrar el menú de selección para la simulación
			functions.MenuRegexFile(results)

		case 0:
			fmt.Println("\n🚪 Saliendo del programa... 🚪")
			return

		default:
			fmt.Println("Opción inválida. Por favor selecciona un número válido.")
		}
	}
}
