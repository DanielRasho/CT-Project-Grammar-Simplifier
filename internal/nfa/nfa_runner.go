/*
Simulación de un autómata finito no determinista con una cadena
*/

package nfa

import "strings"

/**
 * epsilonClosure calcula la ε-cerradura para un único estado en un AFN.
 * La ε-cerradura de un estado es el conjunto de todos los estados que
 * pueden alcanzarse desde el estado dado utilizando transiciones epsilon (ε).
 *
 * Parámetros:
 *  - state: Un puntero a la estructura State que representa el estado inicial para calcular la ε-cerradura.
 *  - transitions: Un slice de Transition que representa todas las transiciones del AFN.
 *
 * Retorno:
 *  - Un slice de punteros a State que contiene todos los estados alcanzables desde el estado inicial utilizando transiciones ε.
 */
func epsilonClosure(state *State, transitions []Transition) []*State {
	closure := map[*State]bool{state: true}
	stack := []*State{state}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, t := range transitions {
			if t.From == current && t.Symbol == "ε" {
				for _, target := range t.To {
					if !closure[target] {
						closure[target] = true
						stack = append(stack, target)
					}
				}
			}
		}
	}

	var closureStates []*State
	for state := range closure {
		closureStates = append(closureStates, state)
	}

	return closureStates
}

/**
 * epsilonClosureOfSet calcula la ε-cerradura para un conjunto de estados en un AFN.
 * Esencialmente, extiende la operación de ε-cerradura para un solo estado a un conjunto de estados.
 *
 * Parámetros:
 *  - states: Un slice de punteros a State que representa los estados iniciales para calcular la ε-cerradura.
 *  - transitions: Un slice de Transition que representa todas las transiciones del AFN.
 *
 * Retorno:
 *  - Un slice de punteros a State que contiene todos los estados alcanzables desde los estados iniciales utilizando transiciones ε.
 */
func epsilonClosureOfSet(states []*State, transitions []Transition) []*State {
	closure := []*State{}
	for _, state := range states {
		closure = append(closure, epsilonClosure(state, transitions)...)
	}
	return closure
}

/**
 * mover realiza la operación de mover(T, a), que calcula el conjunto de estados alcanzables
 * desde un conjunto de estados dado utilizando un símbolo específico.
 *
 * Parámetros:
 *  - states: Un slice de punteros a State que representa los estados actuales desde los que se quiere mover.
 *  - symbol: Un string que representa el símbolo con el cual se realiza la transición.
 *  - transitions: Un slice de Transition que representa todas las transiciones del AFN.
 *
 * Retorno:
 *  - Un slice de punteros a State que contiene todos los estados alcanzables desde los estados iniciales utilizando el símbolo dado.
 */
func mover(states []*State, symbol string, transitions []Transition) []*State {
	var result []*State
	for _, state := range states {
		for _, t := range transitions {
			if t.From == state && t.Symbol == symbol {
				result = append(result, t.To...)
			}
		}
	}
	return result
}

/**
 * RunnerNFA simula el recorrido de un AFN (Autómata Finito No Determinista) con una cadena de entrada.
 * La función utiliza las operaciones de ε-cerradura y mover para determinar si la cadena es aceptada por el AFN.
 *
 * Parámetros:
 *  - nfa: Un puntero a la estructura NFA que representa el autómata finito no determinista.
 *  - cadena: Un string que representa la cadena de entrada que se quiere evaluar.
 *
 * Retorno:
 *  - Un string que indica si la cadena es aceptada ("Sí") o no ("No") por el AFN.
 */
func RunnerNFA(nfa *NFA, cadena string) bool {
	// Convertir la cadena a un slice de caracteres
	simbolos := strings.Split(cadena, "")

	// Inicializar el conjunto de estados actuales con la ε-cerradura del estado inicial
	currentStates := epsilonClosure(nfa.StartState, nfa.Transitions)

	// Procesar cada símbolo en la cadena
	for _, simbolo := range simbolos {
		currentStates = epsilonClosureOfSet(mover(currentStates, simbolo, nfa.Transitions), nfa.Transitions)
	}

	// Verificar si algún estado final es alcanzado
	for _, state := range currentStates {
		if state.IsFinal {
			return true
		}
	}

	return false
}
