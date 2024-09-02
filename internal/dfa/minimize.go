package dfa

import (
	"fmt"
	"strings"

	nfaAutomata "github.com/DanielRasho/TC-1-ShuntingYard/internal/nfa"
)

/**
 * MinimizeDFA minimiza un DFA dado utilizando el algoritmo de minimización de particiones.
 *
 * Parámetros:
 *  - dfa: Un puntero a la estructura DFA que se va a minimizar.
 *
 * Retorno:
 *  - *DFA: Un puntero al DFA minimizado.
 */
func MinimizeDFA(dfa *DFA) *DFA {
	if len(dfa.States) == 0 {
		return dfa // Empty DFA, return as is
	}

	// Step 1: Partition states into final and non-final
	finalStates := make([]*DFAState, 0)
	nonFinalStates := make([]*DFAState, 0)
	for _, state := range dfa.States {
		if state.IsFinal {
			finalStates = append(finalStates, state)
		} else {
			nonFinalStates = append(nonFinalStates, state)
		}
	}

	// Step 2: Initialize partitions
	partitions := []*subset{
		{States: finalStates, ID: 1},
		{States: nonFinalStates, ID: 2},
	}

	// Step 3: Refine partitions until no further changes
	changed := true
	for changed {
		changed = false
		newPartitions := []*subset{}

		for _, partition := range partitions {
			partitionMap := make(map[string]*subset)
			for _, state := range partition.States {
				key := ""
				for symbol, targetState := range dfa.Transitions[state] {
					targetPartition := findPartitionForState(targetState, partitions)
					key += symbol + fmt.Sprintf("%d", targetPartition.ID)
				}

				if partitionMap[key] == nil {
					newSubset := &subset{ID: len(newPartitions) + 1}
					partitionMap[key] = newSubset
					newPartitions = append(newPartitions, newSubset)
				}
				partitionMap[key].States = append(partitionMap[key].States, state)
			}
		}

		if len(newPartitions) != len(partitions) {
			changed = true
			partitions = newPartitions
		}
	}

	// Step 4: Create a new DFA based on the refined partitions
	minimizedDFA := NewDFA()
	stateMap := make(map[*subset]*DFAState)
	for _, partition := range partitions {
		isFinal := false
		stateSet := make(map[*nfaAutomata.State]bool)
		var subsetNames []string

		for _, state := range partition.States {
			if state.IsFinal {
				isFinal = true
			}
			for nfaState := range state.StateSet {
				stateSet[nfaState] = true
				subsetNames = append(subsetNames, nfaState.Name)
			}
		}

		newStateName := "{" + strings.Join(subsetNames, ",") + "}"
		newState := &DFAState{
			Name:     newStateName,
			IsFinal:  isFinal,
			StateSet: stateSet,
		}

		minimizedDFA.States = append(minimizedDFA.States, newState)
		stateMap[partition] = newState
	}

	// Step 5: Add transitions to the new DFA
	for _, partition := range partitions {
		for _, state := range partition.States {
			for symbol, targetState := range dfa.Transitions[state] {
				fromState := stateMap[partition]
				toPartition := findPartitionForState(targetState, partitions)
				toState := stateMap[toPartition]
				minimizedDFA.addTransition(fromState, symbol, toState)
			}
		}
	}

	// Set the start state for the minimized DFA
	initialPartition := findPartitionForState(dfa.StartState, partitions)
	minimizedDFA.StartState = stateMap[initialPartition]

	return minimizedDFA
}

// Helper function to find the partition for a given state
func findPartitionForState(state *DFAState, partitions []*subset) *subset {
	for _, partition := range partitions {
		for _, s := range partition.States {
			if s == state {
				return partition
			}
		}
	}
	return nil
}

func getStateName(isBuildingDFA bool, stateSet map[*nfaAutomata.State]bool, existingNames []string) string {
	if !isBuildingDFA {
		// If we're minimizing, generate a name based on the state set
		var names []string
		for state := range stateSet {
			names = append(names, state.Name)
		}
		return "{" + strings.Join(names, ",") + "}"
	}

	// Generate a sequential state name during DFA construction
	counter := stateNameCounter
	stateNameCounter++

	var name strings.Builder
	repeat := counter / 26 // Calculate how many full alphabet cycles we've completed

	if repeat == 0 {
		return string(rune('A' + counter%26)) // Simple case for the first 26 states
	}

	// For states beyond 'Z', we add 'A' repeated and then the next character
	for i := 0; i < repeat; i++ {
		name.WriteByte('A')
	}
	name.WriteByte(byte('A' + counter%26))

	return name.String()
}
