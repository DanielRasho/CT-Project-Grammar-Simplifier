/*
El AFN (Autómata Finito No Determinista) se compone de estados, transiciones, un estado inicial y estados de aceptación.
Los estados representan los nodos en el autómata, mientras que las transiciones definen cómo se mueve el autómata entre estos estados basado en los símbolos del alfabeto.

Los estados pueden ser de aceptación o no, y las transiciones pueden estar etiquetadas con un símbolo del alfabeto o con una transición ε (epsilon), que permite movimientos entre estados sin consumir un símbolo.
*/

package nfa

// State representa un estado en el autómata.
type State struct {
	Name    string // Nombre del estado.
	IsFinal bool   // Indica si el estado es un estado de aceptación.
}

// Transition representa una transición entre varios estados y un solo estado con un símbolo dado.
type Transition struct {
	From   *State   // Estado de origen de la transición.
	To     []*State // Lista de estados destino.
	Symbol string   // Símbolo que etiqueta la transición.
}

// NFA representa un autómata finito no determinista.
type NFA struct {
	StartState  *State       // Estado inicial del AFN.
	EndState    *State       // Estado final del AFN.
	Transitions []Transition // Lista de transiciones del AFN.
}

// NewState crea un nuevo estado con el nombre proporcionado y un flag de estado final.
func NewState(name string, isFinal bool) *State {
	return &State{
		Name:    name,
		IsFinal: isFinal,
	}
}

// NewTransition crea una nueva transición desde un estado hacia una lista de estados con el símbolo dado.
func NewTransition(from *State, to []*State, symbol string) Transition {
	return Transition{
		From:   from,
		To:     to,
		Symbol: symbol,
	}
}

// NewNFA crea un nuevo AFN con el estado inicial, el estado final y las transiciones proporcionadas.
func NewNFA(startState *State, endState *State, transitions []Transition) *NFA {
	return &NFA{
		StartState:  startState,
		EndState:    endState,
		Transitions: transitions,
	}
}

// ExtractSymbols extrae todos los símbolos únicos de las transiciones en un NFA.
func ExtractSymbols(nfa *NFA) []string {
	symbolSet := make(map[string]bool)
	for _, t := range nfa.Transitions {
		if t.Symbol != "ε" { // Ignorar las transiciones epsilon para la construcción de DFA.
			symbolSet[t.Symbol] = true
		}
	}

	// Convertir el mapa a una lista para su uso.
	var symbols []string
	for symbol := range symbolSet {
		symbols = append(symbols, symbol)
	}
	return symbols
}
