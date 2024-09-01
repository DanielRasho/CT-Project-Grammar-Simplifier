package dfa

import (
	"fmt"

	nfa "github.com/DanielRasho/TC-1-ShuntingYard/internal/nfa"
	runner "github.com/DanielRasho/TC-1-ShuntingYard/internal/runner_simulation"
)

func ConvertNFAtoAFD(nfaAutomaton *nfa.NFA) *DFA {
	dfa := NewDFA()
	queue := make([]*State, 0)
	symbols := nfa.ExtractSymbols(nfaAutomaton) // Obtiene los símbolos del NFA

	// Inicialización con el estado inicial del NFA
	initialClosure := runner.EpsilonClosure(nfaAutomaton.StartState, nfaAutomaton.Transitions)
	initialState := NewState("A", checkFinal(initialClosure), initialClosure)
	dfa.States[initialState.Name] = initialState
	dfa.InitialState = initialState
	queue = append(queue, initialState)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, symbol := range symbols { // Iterar sobre los símbolos extraídos
			targetStates := runner.Mover(current.States, symbol, nfaAutomaton.Transitions)
			targetClosure := runner.EpsilonClosureOfSet(targetStates, nfaAutomaton.Transitions)
			stateName := getStateName(targetClosure)
			if _, exists := dfa.States[stateName]; !exists {
				newState := NewState(stateName, checkFinal(targetClosure), targetClosure)
				dfa.States[stateName] = newState
				queue = append(queue, newState)
			}
			dfa.Transitions = append(dfa.Transitions, NewTransition(current, dfa.States[stateName], symbol))
		}
	}

	return dfa
}

// checkFinal verifica si alguno de los estados de NFA en el conjunto es final.
func checkFinal(states []*nfa.State) bool {
	for _, state := range states {
		if state.IsFinal {
			return true
		}
	}
	return false
}

// getStateName genera un nombre único para un conjunto de estados de NFA.
func getStateName(states []*nfa.State) string {
	name := "Q"
	for _, state := range states {
		name += fmt.Sprintf("_%s", state.Name) // Asegura que los nombres de estados son únicos y consistentes
	}
	return name
}
