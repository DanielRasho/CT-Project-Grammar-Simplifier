/*
El DFA (Autómata Finito Determinista) se compone de estados, transiciones, un estado inicial y estados de aceptación.
Los estados representan conjuntos de estados de un NFA, y las transiciones definen cómo se mueve el DFA entre estos estados basados en los símbolos del alfabeto.

Un DFA garantiza que para cada estado y símbolo del alfabeto, exista a lo sumo una transición hacia otro estado, eliminando la indeterminación presente en un NFA.
*/

package dfa

import (
	nfaAutomata "github.com/DanielRasho/TC-1-ShuntingYard/internal/nfa"
)

/**
 * DFAState representa un conjunto de estados del NFA que forman un único estado en el DFA.
 *
 * Atributos:
 *  - Name: Nombre del estado en el DFA.
 *  - IsFinal: Indica si el estado es final.
 *  - StateSet: Conjunto de estados del NFA que componen este estado del DFA.
 */
type DFAState struct {
	Name     string
	IsFinal  bool
	StateSet map[*nfaAutomata.State]bool
}

/**
 * DFA representa un autómata finito determinista.
 *
 * Atributos:
 *  - StartState: Estado inicial del DFA.
 *  - States: Lista de estados que componen el DFA.
 *  - Transitions: Mapa de transiciones entre los estados del DFA.
 */
type DFA struct {
	StartState  *DFAState
	States      []*DFAState
	Transitions map[*DFAState]map[string]*DFAState
}

/**
 * NewDFA crea y retorna un nuevo DFA vacío.
 *
 * Retorno:
 *  - Un puntero a la estructura DFA inicializada.
 */
func NewDFA() *DFA {
	return &DFA{
		Transitions: make(map[*DFAState]map[string]*DFAState),
	}
}

/**
 * addState agrega un nuevo estado al DFA con un nombre generado automáticamente.
 *
 * Parámetros:
 *  - isFinal: Un booleano que indica si el estado es final.
 *  - stateSet: Un mapa que representa el conjunto de estados del NFA que conforman este estado en el DFA.
 *
 * Retorno:
 *  - Un puntero al nuevo estado del DFA creado.
 */
func (dfa *DFA) addState(isFinal bool, stateSet map[*nfaAutomata.State]bool) *DFAState {
	newState := &DFAState{
		Name:     getStateName(),
		IsFinal:  isFinal,
		StateSet: stateSet,
	}
	dfa.States = append(dfa.States, newState)
	return newState
}

/**
 * addTransition agrega una transición entre estados al DFA.
 *
 * Parámetros:
 *  - from: Un puntero al estado origen de la transición.
 *  - symbol: El símbolo de entrada que dispara la transición.
 *  - to: Un puntero al estado destino de la transición.
 */
func (dfa *DFA) addTransition(from *DFAState, symbol string, to *DFAState) {
	if dfa.Transitions[from] == nil {
		dfa.Transitions[from] = make(map[string]*DFAState)
	}
	dfa.Transitions[from][symbol] = to
}
