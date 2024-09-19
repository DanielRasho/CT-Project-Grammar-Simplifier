package abstract_syntax_tree

import (
	"testing"
)

// Función auxiliar para comparrar valores de nodos
func compareNodes(t *testing.T, got, want Node) {
	if got.String() != want.String() {
		t.Errorf("got %v, want %v", got.String(), want.String())
	}
}

// Test para una expresión postfix simple evaluando una concatenación
func TestBuildASTSimpleConcatenation(t *testing.T) {
	postfix := "ab·c·"
	want := &OperatorNode{
		Value: "·",
		Operands: []Node{
			&OperatorNode{
				Value: "·",
				Operands: []Node{
					&CharacterNode{Value: "a"},
					&CharacterNode{Value: "b"},
				},
			},
			&CharacterNode{Value: "c"},
		},
	}

	got := BuildAST(postfix)
	compareNodes(t, got, want)
}

// Test para una expresión postfix con dos concatenaciones evaluando múltiples concatenaciones
func TestBuildASTMultipleConcatenations(t *testing.T) {
	postfix := "ab·cd··"
	want := &OperatorNode{
		Value: "·",
		Operands: []Node{
			&OperatorNode{
				Value: "·",
				Operands: []Node{
					&CharacterNode{Value: "a"},
					&CharacterNode{Value: "b"},
				},
			},
			&OperatorNode{
				Value: "·",
				Operands: []Node{
					&CharacterNode{Value: "c"},
					&CharacterNode{Value: "d"},
				},
			},
		},
	}

	got := BuildAST(postfix)
	compareNodes(t, got, want)
}

// Test para una expresión postfix que combina una disyunción (OR) y una concatenación
func TestBuildASTDisjunctionAndConcatenation(t *testing.T) {
	postfix := "ab|c·"
	want := &OperatorNode{
		Value: "·",
		Operands: []Node{
			&OperatorNode{
				Value: "|",
				Operands: []Node{
					&CharacterNode{Value: "a"},
					&CharacterNode{Value: "b"},
				},
			},
			&CharacterNode{Value: "c"},
		},
	}

	got := BuildAST(postfix)
	compareNodes(t, got, want)
}

// Test para una expresión postfix compleja evaluando combinaciones de concatenaciones y disyunciones
func TestBuildASTComplexExpression(t *testing.T) {
	postfix := "abc·d|·"
	want := &OperatorNode{
		Value: "·",
		Operands: []Node{
			&OperatorNode{
				Value: "|",
				Operands: []Node{
					&OperatorNode{
						Value: "·",
						Operands: []Node{
							&CharacterNode{Value: "a"},
							&CharacterNode{Value: "b"},
						},
					},
					&CharacterNode{Value: "c"},
				},
			},
			&CharacterNode{Value: "d"},
		},
	}

	got := BuildAST(postfix)
	compareNodes(t, got, want)
}

// Test para una expresión postfix con varios operadores de concatenación encadenados
func TestBuildASTMultipleConcatenationsChain(t *testing.T) {
	postfix := "ab·c·d·"
	want := &OperatorNode{
		Value: "·",
		Operands: []Node{
			&OperatorNode{
				Value: "·",
				Operands: []Node{
					&OperatorNode{
						Value: "·",
						Operands: []Node{
							&CharacterNode{Value: "a"},
							&CharacterNode{Value: "b"},
						},
					},
					&CharacterNode{Value: "c"},
				},
			},
			&CharacterNode{Value: "d"},
		},
	}

	got := BuildAST(postfix)
	compareNodes(t, got, want)
}

// Test para una expresión postfix con caracteres escapados
func TestBuildASTWithEscapedCharacters(t *testing.T) {
	// Expresión postfix con caracteres escapados (barra invertida y otros símbolos)
	postfix := "a\\|·"

	// El árbol de sintaxis esperado debe reflejar los caracteres correctamente escapados.
	want := &OperatorNode{
		Value: "·",
		Operands: []Node{
			&OperatorNode{
				Value: "|", // Operador '|' que aparece en el postfix
				Operands: []Node{
					&CharacterNode{Value: "a"},  // Carácter 'a'
					&CharacterNode{Value: "\\"}, // Carácter barra invertida escapado
				},
			},
			&CharacterNode{Value: "\\"}, // Carácter barra invertida escapado
		},
	}

	// Llamada a la función para construir el AST desde la expresión postfix
	got := BuildAST(postfix)

	// Verificación de que el AST construido es igual al árbol esperado
	compareNodes(t, got, want)
}
