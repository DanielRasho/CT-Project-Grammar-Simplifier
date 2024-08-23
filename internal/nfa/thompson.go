/*
Package nondeterministic_finite_automaton proporciona una implementación para construir y manejar Autómatas Finitos No Deterministas (nfa) a partir de expresiones representadas en un Árbol de Sintaxis Abstracta (AST).
*/
package nfa

import (
	"fmt"

	"github.com/DanielRasho/TC-1-ShuntingYard/internal/abstract_syntax_tree"
)

/*
renameStates renombra los estados del nfa secuencialmente desde 0 hasta n.

Parámetros:
  - nfa: Un puntero a la estructura NFA que representa el autómata a renombrar.
*/
func renameStates(nfa *NFA) {
	stateMap := make(map[*State]string)
	visited := make(map[*State]bool)
	counter := 0

	// Función auxiliar para asignar un nombre si el estado no ha sido nombrado
	assignName := func(s *State) {
		if _, exists := stateMap[s]; !exists {
			stateMap[s] = fmt.Sprintf("q%d", counter)
			counter++
		}
	}

	// Función de DFS para recorrer y renombrar estados
	var dfs func(s *State)
	dfs = func(s *State) {
		if visited[s] {
			return
		}
		visited[s] = true
		assignName(s)

		// Recorrer las transiciones desde el estado actual
		for _, t := range nfa.Transitions {
			if t.From == s {
				for _, toState := range t.To {
					dfs(toState) // Recursivamente nombrar los estados alcanzables
				}
			}
		}
	}

	// Iniciar DFS desde el estado inicial
	dfs(nfa.StartState)

	// Actualizar los nombres en los estados
	for state, name := range stateMap {
		state.Name = name
	}
}

/*
BuildNFA construye un nfa a partir de un nodo AST (árbol de sintaxis abstracta).

Parámetros:
  - node: Un nodo del tipo abstract_syntax_tree.Node que representa la raíz del AST.

Retorno:
  - Un puntero a la estructura NFA que representa el autómata construido.

Panic:
  - Si el nodo AST contiene un operador no soportado o si la cantidad de operandos no es la esperada.
*/
func BuildNFA(node abstract_syntax_tree.Node) *NFA {
	var nfa *NFA

	switch n := node.(type) {
	case *abstract_syntax_tree.CharacterNode:
		nfa = nfaBasic(n.Value)
	case *abstract_syntax_tree.OperatorNode:
		switch n.Value {
		case "·":
			nfa = nfaConcatenation(n.Operands)
		case "|":
			nfa = nfaUnion(n.Operands)
		case "*":
			nfa = nfaKleene(n.Operands)
		}
	}

	// Si se construyó un nfa, desmarcar todos los estados como finales y
	// marcar el último estado como final
	if nfa != nil {
		// Desmarcar todos los estados como no finales
		for _, t := range nfa.Transitions {
			t.From.IsFinal = false
			for _, toState := range t.To {
				toState.IsFinal = false
			}
		}

		// Marcar solo el estado final del NFA completo como final
		nfa.EndState.IsFinal = true

		// Renombrar los estados
		renameStates(nfa)
	}

	return nfa
}

/*
nfaBasic crea un nfa para un solo carácter usando un label para la transición.

Parámetros:
  - label: Una cadena que representa el carácter para la transición.

Retorno:
  - Un puntero a la estructura NFA que representa el autómata básico creado.
*/
func nfaBasic(label string) *NFA {
	// Crear los estados
	start := NewState(fmt.Sprintf("Start_%s", label), false)
	end := NewState(fmt.Sprintf("End_%s", label), false)

	// Crear la transición con el label proporcionado
	transition := NewTransition(start, []*State{end}, label)

	// Crear y devolver el nfa
	return NewNFA(start, end, []Transition{transition})
}

/*
nfaConcatenation crea un nfa para la concatenación de dos nfas.

Parámetros:
  - operands: Un slice de nodos AST que representan los operandos para la concatenación.

Retorno:
  - Un puntero a la estructura NFA que representa el autómata de concatenación creado.

Panic:
  - Si el número de operandos no es exactamente 2.
*/
func nfaConcatenation(operands []abstract_syntax_tree.Node) *NFA {
	if len(operands) != 2 {
		panic("nfaConcatenation expects exactly 2 operands")
	}

	// Construir los nfas para los operandos
	nfa1 := BuildNFA(operands[0])
	nfa2 := BuildNFA(operands[1])

	// Crear un nuevo estado para la transición ε
	newStart := nfa1.StartState
	newEnd := nfa2.EndState

	// Crear la transición ε entre el estado final del primer NFA y el estado inicial del segundo NFA
	epsilonTransition := NewTransition(nfa1.EndState, []*State{nfa2.StartState}, "ε")

	// Combinar transiciones de ambos nfas
	allTransitions := append(nfa1.Transitions, nfa2.Transitions...)
	allTransitions = append(allTransitions, epsilonTransition)

	// Crear el nuevo NFA
	return &NFA{
		StartState:  newStart,
		EndState:    newEnd,
		Transitions: allTransitions,
	}
}

/*
nfaUnion crea un nfa para la unión de dos nfas.

Parámetros:
  - operands: Un slice de nodos AST que representan los operandos para la unión.

Retorno:
  - Un puntero a la estructura NFA que representa el autómata de unión creado.

Panic:
  - Si el número de operandos no es exactamente 2.
*/
func nfaUnion(operands []abstract_syntax_tree.Node) *NFA {
	if len(operands) != 2 {
		panic("nfaUnion expects exactly 2 operands")
	}

	// Construir los nfas para los operandos
	nfa1 := BuildNFA(operands[0])
	nfa2 := BuildNFA(operands[1])

	// Crear un nuevo estado para la transición ε inicial
	start := NewState("Start_|", false)
	end := NewState("End_|", false)

	// Crear transiciones ε desde el nuevo estado inicial a los estados iniciales de los nfas
	epsilonTransition1 := NewTransition(start, []*State{nfa1.StartState, nfa2.StartState}, "ε")

	// Crear transiciones ε desde los estados finales de los nfas al nuevo estado final
	epsilonTransitionEnd1 := NewTransition(nfa1.EndState, []*State{end}, "ε")
	epsilonTransitionEnd2 := NewTransition(nfa2.EndState, []*State{end}, "ε")

	// Combinar transiciones de ambos nfas con las nuevas transiciones ε
	allTransitions := append(nfa1.Transitions, nfa2.Transitions...)
	allTransitions = append(allTransitions, epsilonTransition1, epsilonTransitionEnd1, epsilonTransitionEnd2)

	// Crear el nuevo NFA
	return &NFA{
		StartState:  start,
		EndState:    end,
		Transitions: allTransitions,
	}
}

/*
nfaKleene crea un nfa para la cerradura de Kleene.

Parámetros:
  - operands: Un slice de nodos AST que representan el operando para la cerradura de Kleene.

Retorno:
  - Un puntero a la estructura NFA que representa el autómata de cerradura de Kleene creado.

Panic:
  - Si el número de operandos no es exactamente 1.
*/
func nfaKleene(operands []abstract_syntax_tree.Node) *NFA {
	if len(operands) != 1 {
		panic("nfaKleene expects exactly 1 operand")
	}

	// Construir el nfa del operando
	nfa1 := BuildNFA(operands[0])

	// Crear nuevos estados para la transición ε inicial y final
	start := NewState("Start_*", false)
	end := NewState("End_*", false)

	// Crear transiciones ε:
	// 1. Desde el nuevo estado inicial al estado inicial del NFA y al nuevo estado final.
	epsilonTransition1 := NewTransition(start, []*State{nfa1.StartState, end}, "ε")

	// 2. Desde el estado final del NFA al nuevo estado final y al estado inicial del NFA.
	epsilonTransition2 := NewTransition(nfa1.EndState, []*State{nfa1.StartState, end}, "ε")

	// Combinar todas las transiciones
	allTransitions := append(nfa1.Transitions, epsilonTransition1, epsilonTransition2)

	// Crear el nuevo NFA con los estados y transiciones actualizados
	return &NFA{
		StartState:  start,
		EndState:    end,
		Transitions: allTransitions,
	}
}
