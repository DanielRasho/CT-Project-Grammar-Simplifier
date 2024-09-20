package abstract_syntax_tree

import (
	"testing"

	shuntingyard "github.com/DanielRasho/Computation-Theory/internal/shuntingyard"
)

// Función auxiliar para comparar valores de nodos
func compareNodes(t *testing.T, got, want Node) {
	if got.String() != want.String() {
		t.Errorf("got %v, want %v", got.String(), want.String())
	}
}

// Test para una expresión regex simple evaluando una concatenación
func TestBuildASTSimpleConcatenation(t *testing.T) {
	regex := "abc"
	_, postfix, _ := shuntingyard.RegexToPostfix(regex, false)

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

// Test para una expresión regex con múltiples concatenaciones
func TestBuildASTMultipleConcatenations(t *testing.T) {
	regex := "abcd"
	_, postfix, _ := shuntingyard.RegexToPostfix(regex, false)

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

// Test para una expresión regex que combina una disyunción (OR) y una concatenación
func TestBuildASTDisjunctionAndConcatenation(t *testing.T) {
	regex := "ab|c"
	_, postfix, _ := shuntingyard.RegexToPostfix(regex, false)
	want := &OperatorNode{
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
	}

	got := BuildAST(postfix)
	compareNodes(t, got, want)
}

// Test para una expresión regex compleja evaluando combinaciones de concatenaciones y disyunciones
func TestBuildASTComplexExpression(t *testing.T) {
	regex := "ab|cd"
	_, postfix, _ := shuntingyard.RegexToPostfix(regex, false)

	want := &OperatorNode{
		Value: "|",
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

// Test para una expresión regex con varios operadores de concatenación encadenados
func TestBuildASTMultipleConcatenationsChain(t *testing.T) {
	regex := "abcd"
	_, postfix, _ := shuntingyard.RegexToPostfix(regex, false)

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

// Test para una expresión regex con caracteres escapados
func TestBuildASTWithEscapedCharacters(t *testing.T) {
	regex := "a\\|"
	_, postfix, _ := shuntingyard.RegexToPostfix(regex, false)

	want := &OperatorNode{
		Value: "·",
		Operands: []Node{
			&CharacterNode{Value: "a"},
			&CharacterNode{Value: "|"},
		},
	}

	got := BuildAST(postfix)
	compareNodes(t, got, want)
}
