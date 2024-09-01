package main

import (
	"fmt"
	"strings"

	ast "github.com/DanielRasho/TC-1-ShuntingYard/internal/abstract_syntax_tree"
	io "github.com/DanielRasho/TC-1-ShuntingYard/internal/io"
	shuttingyard "github.com/DanielRasho/TC-1-ShuntingYard/internal/shuntingyard"
)

func main() {

	// Llama a la función de lectura de archivo
	lines, err := io.ReaderTXT("input_data/ast.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Lista para almacenar los ASTs
	// var astList []ast.Node

	// Procesa cada línea leída del archivo
	for i, line := range lines {
		// Convierte la expresión regular a postfix usando Shunting Yard
		postfix, _, _ := shuttingyard.RegexToPostfix(line, false)
		// fmt.Print("Se hizo el postfix: ", postfix, "\n")

		// Construye el AST a partir del postfix
		root := ast.BuildAST(postfix)
		// fmt.Print("Su AST es: ", root, "\n\n")

		ast.GenerateImageFromRoot(root, fmt.Sprintf("./graphs/graph%d.png", i))

	}

}

// Imprime el árbol completo del AST
func printASTTree(node ast.Node, level int) {
	// Imprime el nodo actual
	switch n := node.(type) {
	case *ast.CharacterNode:
		fmt.Printf("%sCharacterNode: %s\n", indent(level), n.Value)
	case *ast.OperatorNode:
		fmt.Printf("%sOperatorNode: %s\n", indent(level), n.Value)
		for _, operand := range n.GetOperands() {
			printASTTree(operand, level+1)
		}
	}
}

// Devuelve un string de indentación para el nivel dado
func indent(level int) string {
	return strings.Repeat("  ", level)
}
