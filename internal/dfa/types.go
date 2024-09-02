/*
El DFA (Autómata Finito Determinista) se compone de estados, transiciones, un estado inicial y estados de aceptación.
Los estados representan conjuntos de estados de un NFA, y las transiciones definen cómo se mueve el DFA entre estos estados basados en los símbolos del alfabeto.

Un DFA garantiza que para cada estado y símbolo del alfabeto, exista a lo sumo una transición hacia otro estado, eliminando la indeterminación presente en un NFA.
*/

package dfa

import (
	"fmt"
	"strings"

	nfaAutomata "github.com/DanielRasho/TC-1-ShuntingYard/internal/nfa"
)

var stateNameCounter = 0

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

// subset representa un conjunto de estados del DFA.
type subset struct {
	States []*DFAState
	ID     int
}

// partition representa una partición del conjunto de estados del DFA.
type partition struct {
	Subsets []*subset
	ID      int
}

// String method for DFAState to return a string representation of the DFAState
func (state *DFAState) String() string {
	stateNames := []string{}
	for nfaState := range state.StateSet {
		stateNames = append(stateNames, nfaState.Name)
	}

	// Join the state names and return a formatted string
	return fmt.Sprintf("State: %s, IsFinal: %t, StateSet: {%s}", state.Name, state.IsFinal, strings.Join(stateNames, ", "))
}

// String method for subset to return a string representation of the subset
func (s *subset) String() string {
	stateStrings := []string{}
	for _, state := range s.States {
		stateStrings = append(stateStrings, state.String())
	}

	// Join all the DFAState strings and return a formatted string
	return fmt.Sprintf("Subset ID: %d, States: [%s]", s.ID, strings.Join(stateStrings, "; "))
}

/**
 * NewDFA crea y retorna un nuevo DFA vacío.
 *
 * Retorno:
 *  - *DFA: Un puntero a la estructura DFA inicializada.
 */
func NewDFA() *DFA {
	return &DFA{
		Transitions: make(map[*DFAState]map[string]*DFAState),
	}
}

/**
 * getStateName genera nombres de estados secuenciales como A, B, C, ..., Z, AA, AB, etc., manejando correctamente el ciclo después de Z.
 *
 * Parámetros:
 *  - isBuildingDFA: Un booleano que indica si estamos en la fase de construcción del DFA (true) o minimización (false).
 *  - stateSet: Un mapa que representa el conjunto de estados del NFA que conforman este estado en el DFA. Se usa solo si isBuildingDFA es false.
 *  - existingNames: Un slice de strings que contiene los nombres de los estados existentes. Se usa para evitar duplicados y crear nombres únicos.
 *
 * Retorno:
 *  - string: Un string que representa el nombre del nuevo estado.
 */

func getStateName(isBuildingDFA bool,
	// stateSet map[*nfaAutomata.State]bool,
	existingNames []string) string {
	if !isBuildingDFA {
		// Si no estamos construyendo el DFA, generar el nombre basado en el conjunto de estados
		// var names []string
		//for state := range stateSet {
		//	names = append(names, state.Name)
		// }
		return strings.Join(existingNames, ",")
	}

	// Generar un nombre secuencial para la construcción del DFA
	counter := stateNameCounter
	stateNameCounter++

	var name strings.Builder
	repeat := counter / 26 // Calcula cuántas veces se ha completado un ciclo completo a través del alfabeto

	if repeat == 0 {
		return string(rune('A' + counter%26)) // Caso simple para los primeros 26 estados
	}

	// Para estados después de 'Z', agregamos 'A' repetido y luego el siguiente carácter
	for i := 0; i < repeat; i++ {
		name.WriteByte('A')
	}
	name.WriteByte(byte('A' + counter%26))

	return name.String()
}

/**
 * addState agrega un nuevo estado al DFA.
 *
 * Parámetros:
 *  - isFinal: Un booleano que indica si el estado es final.
 *  - stateSet: Un mapa que representa el conjunto de estados del NFA que conforman este estado en el DFA.
 *  - isBuildingDFA: Un booleano que indica si estamos en la fase de construcción del DFA (true) o minimización (false).
 *  - existingNames: Un slice de strings que contiene los nombres de los estados existentes, usado para asegurar la unicidad de los nombres.
 *
 * Retorno:
 *  - *DFAState: Un puntero al nuevo estado del DFA creado.
 */
func (dfa *DFA) addState(isFinal bool, stateSet map[*nfaAutomata.State]bool, isBuildingDFA bool, existingNames []string) *DFAState {
	newState := &DFAState{
		Name: getStateName(isBuildingDFA,
			//stateSet,
			existingNames),
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
