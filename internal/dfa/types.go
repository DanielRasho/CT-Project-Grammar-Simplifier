package dfa

import "github.com/DanielRasho/TC-1-ShuntingYard/internal/nfa"

// State representa un estado en el AFD, que es un conjunto de estados de NFA.
type State struct {
	Name    string       // Nombre del estado en el AFD.
	IsFinal bool         // Indica si es un estado de aceptación.
	States  []*nfa.State // Estados del NFA que componen este estado del AFD.
}

// Transition representa una transición en el AFD.
type Transition struct {
	From   *State // Estado de origen de la transición.
	To     *State // Estado destino de la transición.
	Symbol string // Símbolo que etiqueta la transición.
}

// DFA representa un autómata finito determinista.
type DFA struct {
	States       map[string]*State
	InitialState *State
	Transitions  []Transition
}

// NewState crea un nuevo estado de AFD.
func NewState(name string, isFinal bool, nfaStates []*nfa.State) *State {
	return &State{
		Name:    name,
		IsFinal: isFinal,
		States:  nfaStates,
	}
}

// NewTransition crea una nueva transición en el AFD.
func NewTransition(from *State, to *State, symbol string) Transition {
	return Transition{
		From:   from,
		To:     to,
		Symbol: symbol,
	}
}

// NewDFA crea un nuevo AFD vacío.
func NewDFA() *DFA {
	return &DFA{
		States:      make(map[string]*State),
		Transitions: make([]Transition, 0),
	}
}
