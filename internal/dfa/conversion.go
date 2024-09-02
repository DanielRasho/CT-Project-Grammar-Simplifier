/*
La conversión de un NFA a un DFA se realiza mediante el algoritmo de subconjuntos, donde se construye un DFA que simula el comportamiento de un NFA.

El método de subconjuntos consiste en agrupar los estados del NFA en un conjunto, tratándolos como un solo estado en el DFA. Se consideran todas las posibles transiciones para cada símbolo del alfabeto, asegurando que el DFA resultante sea determinista y completo.
*/

package dfa

import (
	nfaAutomata "github.com/DanielRasho/TC-1-ShuntingYard/internal/nfa"
	runner "github.com/DanielRasho/TC-1-ShuntingYard/internal/runner_simulation"
)

/**
 * BuildDFA realiza la conversión de un NFA a un DFA utilizando el método de subconjuntos.
 *
 * Parámetros:
 *  - nfa: Un puntero al NFA que se desea convertir en un DFA.
 *
 * Retorno:
 *  - Un puntero al DFA resultante.
 */
func BuildDFA(nfa *nfaAutomata.NFA) *DFA {
	dfa := NewDFA()
	initialClosure := runner.EpsilonClosureOfSet([]*nfaAutomata.State{nfa.StartState}, nfa.Transitions)

	initialSet := make(map[*nfaAutomata.State]bool)
	for _, state := range initialClosure {
		initialSet[state] = true
	}

	// Determinar si el conjunto inicial contiene algún estado de aceptación.
	isFinal := false
	for state := range initialSet {
		if state.IsFinal {
			isFinal = true
			break
		}
	}

	// Crear el estado inicial del DFA.
	initialState := dfa.addState(isFinal, initialSet, true)
	dfa.StartState = initialState

	// Lista de nuevos estados a procesar.
	unmarkedStates := []*DFAState{initialState}

	// Procesar estados no marcados.
	for len(unmarkedStates) > 0 {
		currentDFAState := unmarkedStates[0]
		unmarkedStates = unmarkedStates[1:]

		// Procesar cada símbolo del alfabeto del NFA.
		symbols := nfaAutomata.ExtractSymbols(nfa)
		for _, symbol := range symbols {
			nextStates := runner.Mover(stateSetToList(currentDFAState.StateSet), symbol, nfa.Transitions)
			closure := runner.EpsilonClosureOfSet(nextStates, nfa.Transitions)
			nextSet := make(map[*nfaAutomata.State]bool)
			for _, s := range closure {
				nextSet[s] = true
			}

			// Ver si este conjunto de estados ya está en el DFA.
			nextDFAState := findStateInDFA(dfa, nextSet)
			if nextDFAState == nil {
				// Marcar si es un estado final.
				isNextFinal := false
				for s := range nextSet {
					if s.IsFinal {
						isNextFinal = true
						break
					}
				}
				nextDFAState = dfa.addState(isNextFinal, nextSet, true)
				unmarkedStates = append(unmarkedStates, nextDFAState)
			}

			// Agregar la transición al DFA.
			dfa.addTransition(currentDFAState, symbol, nextDFAState)
		}
	}

	return dfa
}

/**
 * stateSetToList convierte un mapa de estados (set) a una lista de punteros a State.
 *
 * Parámetros:
 *  - set: Un mapa que representa un conjunto de estados.
 *
 * Retorno:
 *  - Un slice de punteros a State que representa los estados en el conjunto.
 */
func stateSetToList(set map[*nfaAutomata.State]bool) []*nfaAutomata.State {
	var list []*nfaAutomata.State
	for state := range set {
		list = append(list, state)
	}
	return list
}

/**
 * findStateInDFA busca en el DFA para encontrar un estado existente que coincida con un conjunto dado de estados del NFA.
 *
 * Parámetros:
 *  - dfa: Un puntero al DFA en el que se realiza la búsqueda.
 *  - set: Un mapa que representa el conjunto de estados del NFA a buscar.
 *
 * Retorno:
 *  - Un puntero al DFAState correspondiente si se encuentra, o nil si no se encuentra.
 */
func findStateInDFA(dfa *DFA, set map[*nfaAutomata.State]bool) *DFAState {
	for _, state := range dfa.States {
		if equalStateSets(state.StateSet, set) {
			return state
		}
	}
	return nil
}

/**
 * equalStateSets compara dos conjuntos de estados para determinar si son iguales.
 *
 * Parámetros:
 *  - set1: Un mapa que representa el primer conjunto de estados.
 *  - set2: Un mapa que representa el segundo conjunto de estados.
 *
 * Retorno:
 *  - Un booleano que indica si los dos conjuntos son iguales (true) o no (false).
 */
func equalStateSets(set1, set2 map[*nfaAutomata.State]bool) bool {
	if len(set1) != len(set2) {
		return false
	}
	for s := range set1 {
		if !set2[s] {
			return false
		}
	}
	return true
}
