package dfa

import (
	"fmt"
	"sort"
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
	// Categorías para almacenar los estados
	var acceptedStates []*DFAState
	var nonAcceptedStates []*DFAState

	// Recorre todos los estados del DFA
	for _, state := range dfa.States {
		if state.IsFinal {
			// Si el estado es de aceptación, añadir a la lista de aceptados
			acceptedStates = append(acceptedStates, state)
		} else {
			// Si el estado no es de aceptación, añadir a la lista de no aceptados
			nonAcceptedStates = append(nonAcceptedStates, state)
		}
	}

	// Crear punteros para los subsets
	subset1 := &subset{States: acceptedStates, ID: 1}
	subset2 := &subset{States: nonAcceptedStates, ID: 2}

	// Crear partición inicial con punteros a los subsets
	initialPartition := partition{Subsets: []*subset{subset1, subset2}, ID: 0}
	finalPartitions := doPartition(initialPartition, dfa.Transitions)

	// Crear el nuevo DFA
	newDFA := NewDFA()
	stateMap := make(map[string]*DFAState)

	// Recopilar nombres de estados existentes
	var existingNames []string
	for _, state := range dfa.States {
		existingNames = append(existingNames, state.Name)
	}

	// Crear nuevos estados en el DFA basado en las particiones
	for _, p := range finalPartitions {
		stateSet := make(map[*nfaAutomata.State]bool)
		for _, subset := range p.Subsets {
			for _, state := range subset.States {
				for nfaState := range state.StateSet {
					stateSet[nfaState] = true
				}
			}
		}

		// Determinar si el nuevo estado es final
		isFinal := false
		for _, subset := range p.Subsets {
			for _, state := range subset.States {
				if state.IsFinal {
					isFinal = true
					break
				}
			}
		}

		// Crear un nuevo estado y agregarlo al nuevo DFA
		stateName := generateStateName(p.Subsets) // Esto ahora usará los nombres del DFA original
		newState := newDFA.addState(isFinal, stateSet, false, existingNames)
		stateMap[stateName] = newState
	}

	// Definir las transiciones para el nuevo DFA
	for _, p := range finalPartitions {
		for _, mysubset := range p.Subsets {
			for _, state := range mysubset.States {
				for symbol, nextState := range dfa.Transitions[state] {
					// Encuentra el estado de la partición correspondiente
					nextPartition := findPartition(nextState, finalPartitions)
					nextStateName := generateStateName([]*subset{nextPartition})
					fromStateName := generateStateName([]*subset{mysubset})

					// Agregar la transición al nuevo DFA
					newDFA.addTransition(stateMap[fromStateName], symbol, stateMap[nextStateName])
				}
			}
		}
	}

	// Definir el estado inicial del nuevo DFA
	initialSubset := findPartition(dfa.States[0], finalPartitions)
	newDFA.StartState = stateMap[generateStateName([]*subset{initialSubset})]

	return newDFA
}

/**
 * generateStateName genera un nombre único para un estado basado en los nombres de los estados del DFA original.
 *
 * Parámetros:
 *  - subsets: Un slice de punteros a estructuras subset que representan los estados en una partición.
 *
 * Retorno:
 *  - string: Un nombre único generado para el nuevo estado del DFA.
 */
func generateStateName(subsets []*subset) string {
	var stateNames []string
	for _, s := range subsets {
		for _, state := range s.States {
			stateNames = append(stateNames, state.Name) // Utiliza el nombre del DFA, no los nombres de los estados del NFA
		}
	}
	sort.Strings(stateNames) // Asegúrate de que los nombres estén ordenados para evitar inconsistencias
	return "{" + strings.Join(stateNames, ",") + "}"
}

/**
 * findPartition encuentra la partición a la que pertenece un estado.
 *
 * Parámetros:
 *  - state: Un puntero al estado del DFA que se busca.
 *  - partitions: Un slice de estructuras partition que representan las particiones actuales del DFA.
 *
 * Retorno:
 *  - *subset: Un puntero al subset que contiene el estado dado.
 */
func findPartition(state *DFAState, partitions []partition) *subset {
	for _, p := range partitions {
		for _, s := range p.Subsets {
			for _, st := range s.States {
				if st == state {
					return s
				}
			}
		}
	}
	return nil
}

/**
 * generateKey genera una clave única para un estado basado en sus transiciones.
 *
 * Parámetros:
 *  - state: Un puntero al estado del DFA para el cual se genera la clave.
 *  - transitions: Un mapa de transiciones del DFA que asocia estados con símbolos de entrada.
 *
 * Retorno:
 *  - string: Una clave única generada para el estado basado en sus transiciones.
 */
func generateKey(state *DFAState, transitions map[*DFAState]map[string]*DFAState) string {
	transitionsMap := transitions[state]
	var keyParts []string
	for symbol, nextState := range transitionsMap {
		keyParts = append(keyParts, fmt.Sprintf("%s:%s", symbol, nextState.Name))
	}
	return strings.Join(keyParts, ",")
}

/**
 * doPartition divide los estados en particiones más pequeñas basadas en las transiciones.
 *
 * Parámetros:
 *  - initialPartition: La partición inicial que contiene los estados separados en aceptados y no aceptados.
 *  - transitions: Un mapa de transiciones del DFA que asocia estados con símbolos de entrada.
 *
 * Retorno:
 *  - []partition: Un slice de particiones resultantes después de aplicar la minimización.
 */
func doPartition(initialPartition partition, transitions map[*DFAState]map[string]*DFAState) []partition {
	var partitions []partition
	partitions = append(partitions, initialPartition)

	for {
		var newPartitions []*subset
		partChanged := false

		// Recorre cada partición
		for _, s := range partitions[0].Subsets {
			subsetMap := make(map[string][]*DFAState)

			// Recorre cada estado en el subset
			for _, state := range s.States {
				key := generateKey(state, transitions)
				subsetMap[key] = append(subsetMap[key], state)
			}

			// Crea nuevos subsets a partir de los grupos
			for _, states := range subsetMap {
				newPartitions = append(newPartitions, &subset{States: states, ID: len(newPartitions) + 1})
			}

			if len(newPartitions) > len(partitions[0].Subsets) {
				partChanged = true
			}
		}

		if !partChanged {
			break
		}

		partitions = []partition{{Subsets: newPartitions, ID: len(partitions) + 1}}
	}

	return partitions
}
