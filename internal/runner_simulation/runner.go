package runnersimulation

import (
	"strings"

	dfaAutomata "github.com/DanielRasho/TC-1-ShuntingYard/internal/dfa"
	nfaAutomata "github.com/DanielRasho/TC-1-ShuntingYard/internal/nfa"
)

/**
 * RunnerNFA simula el recorrido de un AFN (Autómata Finito No Determinista) con una cadena de entrada.
 * La función utiliza las operaciones de ε-cerradura y Mover para determinar si la cadena es aceptada por el AFN.
 *
 * Parámetros:
 *  - nfa: Un puntero a la estructura NFA que representa el autómata finito no determinista.
 *  - cadena: Un string que representa la cadena de entrada que se quiere evaluar.
 *
 * Retorno:
 *  - Un booleano que indica si la cadena es aceptada (true) o no (false) por el AFN.
 */
func RunnerNFA(nfa *nfaAutomata.NFA, cadena string) bool {
	// Convertir la cadena a un slice de caracteres
	simbolos := strings.Split(cadena, "")

	// Inicializar el conjunto de estados actuales con la ε-cerradura del estado inicial
	currentStates := nfaAutomata.EpsilonClosure(nfa.StartState, nfa.Transitions)

	// Procesar cada símbolo en la cadena
	for _, simbolo := range simbolos {
		currentStates = nfaAutomata.EpsilonClosureOfSet(nfaAutomata.Mover(currentStates, simbolo, nfa.Transitions), nfa.Transitions)
	}

	// Verificar si algún estado final es alcanzado
	for _, state := range currentStates {
		if state.IsFinal {
			return true
		}
	}

	return false
}

/**
 * RunnerDFA simula el recorrido de un DFA (Autómata Finito Determinista) con una cadena de entrada.
 * La función utiliza la operación de Mover para determinar si la cadena es aceptada por el DFA.
 *
 * Parámetros:
 *  - dfa: Un puntero a la estructura DFA que representa el autómata finito determinista.
 *  - cadena: Un string que representa la cadena de entrada que se quiere evaluar.
 *
 * Retorno:
 *  - Un booleano que indica si la cadena es aceptada (true) o no (false) por el DFA.
 */
func RunnerDFA(dfa *dfaAutomata.DFA, cadena string) bool {
	// Convertir la cadena a un slice de caracteres
	simbolos := strings.Split(cadena, "")

	// Inicializar el estado actual como el estado inicial del DFA
	currentState := dfa.StartState

	// Procesar cada símbolo en la cadena
	for _, simbolo := range simbolos {
		currentState = dfa.Mover(currentState, simbolo)
		if currentState == nil {
			// Si no hay transición definida para el símbolo, la cadena no es aceptada
			return false
		}
	}

	// Verificar si el estado final alcanzado es un estado aceptado
	return currentState.IsFinal
}
