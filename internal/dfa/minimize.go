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
	fmt.Println("Iniciando la minimización del DFA...")

	// Categorías para almacenar los estados
	var acceptedStates []*DFAState
	var nonAcceptedStates []*DFAState

	// Recorre todos los estados del DFA
	for _, state := range dfa.States {
		if state.IsFinal {
			acceptedStates = append(acceptedStates, state)
		} else {
			nonAcceptedStates = append(nonAcceptedStates, state)
		}
	}

	fmt.Println("Estados de aceptación:", len(acceptedStates))
	fmt.Println("Estados no aceptados:", len(nonAcceptedStates))

	// Crear punteros para los subsets
	subset1 := &subset{States: acceptedStates, ID: 1}
	subset2 := &subset{States: nonAcceptedStates, ID: 2}

	// Crear partición inicial con punteros a los subsets
	initialPartition := partition{Subsets: []*subset{subset1, subset2}, ID: 0}
	fmt.Println("Partición inicial creada con 2 subconjuntos.")

	finalPartitions := doPartition(initialPartition, dfa.Transitions)
	fmt.Println("Particiones finales obtenidas:", len(finalPartitions))

	// Crear el nuevo DFA
	newDFA := NewDFA()
	stateMap := make(map[string]*DFAState)

	// Crear nuevos estados en el DFA basado en las particiones
	for _, p := range finalPartitions {
		for _, mysubset := range p.Subsets {
			stateSet := make(map[*nfaAutomata.State]bool)
			isFinal := false

			for _, state := range mysubset.States {
				for nfaState := range state.StateSet {
					stateSet[nfaState] = true
				}
				if state.IsFinal {
					isFinal = true
				}
			}

			stateName := generateStateName([]*subset{mysubset})
			newState := newDFA.addState(isFinal, stateSet, false, []string{stateName})
			if newState == nil {
				fmt.Printf("Error al crear el nuevo estado: %s\n", stateName)
				continue
			}
			stateMap[stateName] = newState
			fmt.Printf("Nuevo estado creado: %s, es final: %v\n", stateName, isFinal)
		}
	}

	// Definir las transiciones para el nuevo DFA
	for _, p := range finalPartitions {
		for _, mysubset := range p.Subsets {
			for _, state := range mysubset.States {
				for symbol, nextState := range dfa.Transitions[state] {
					nextPartition := findPartition(nextState, finalPartitions)
					if nextPartition == nil {
						fmt.Printf("No se encontró la partición para el estado de destino: %s\n", nextState.Name)
						continue
					}

					nextStateName := generateStateName([]*subset{nextPartition})
					fromStateName := generateStateName([]*subset{mysubset})

					// Verificar existencia de estados antes de añadir la transición
					fromState, ok1 := stateMap[fromStateName]
					toState, ok2 := stateMap[nextStateName]
					if !ok1 || !ok2 {
						fmt.Printf("Error: Transición no válida desde %s a %s con símbolo %s\n", fromStateName, nextStateName, symbol)
						continue
					}

					// Agregar la transición al nuevo DFA
					newDFA.addTransition(fromState, symbol, toState)
					fmt.Printf("Transición añadida: %s --%s--> %s\n", fromStateName, symbol, nextStateName)
				}
			}
		}
	}

	// Definir el estado inicial del nuevo DFA
	initialSubset := findPartition(dfa.StartState, finalPartitions)
	if initialSubset == nil {
		fmt.Println("Error: No se encontró la partición para el estado inicial.")
		return nil
	}
	newDFA.StartState = stateMap[generateStateName([]*subset{initialSubset})]
	if newDFA.StartState == nil {
		fmt.Println("Error: Estado inicial no válido.")
		return nil
	}
	fmt.Printf("Estado inicial del nuevo DFA: %s\n", newDFA.StartState.Name)

	fmt.Println("Minimización completada.")
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
		var newSubsets []*subset
		partChanged := false

		for _, s := range partitions[0].Subsets {
			subsetMap := make(map[string][]*DFAState)

			for _, state := range s.States {
				key := generateKey(state, transitions)
				subsetMap[key] = append(subsetMap[key], state)
			}

			for _, states := range subsetMap {
				newSubsets = append(newSubsets, &subset{States: states, ID: len(newSubsets) + 1})
			}

			if len(newSubsets) > len(partitions[0].Subsets) {
				partChanged = true
			}
		}

		if !partChanged {
			break
		}

		partitions = []partition{{Subsets: newSubsets, ID: len(partitions) + 1}}
	}

	return partitions
}
