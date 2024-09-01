package main

import (
	"bufio"
	"fmt"
	"os"

	balancer "github.com/DanielRasho/TC-1-ShuntingYard/internal/balancer"
)

func main() {
	// Abre el archivo de entrada
	file, err := os.Open("input_data/balancerInput.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Lee el archivo línea por línea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Procesando línea:", line)
		isBalanced, steps := balancer.IsBalanced(line)
		fmt.Println("Pasos de la pila:")
		for _, step := range steps {
			fmt.Println(step)
		}
		if isBalanced {
			fmt.Println("La expresión está bien balanceada.")
		} else {
			fmt.Println("La expresión no está bien balanceada.")
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
}
