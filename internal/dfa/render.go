package dfa

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GenerateDOT generates a DOT representation of a DFA as a string.
func GenerateDOT(dfa *DFA) string {
	var sb strings.Builder

	// Write the Graphviz dot header
	sb.WriteString("digraph DFA {\n")
	sb.WriteString("    rankdir=LR;\n") // Left to right orientation

	// Check if the DFA has any states
	if dfa.StartState == nil || len(dfa.States) == 0 {
		panic("DFA has no states defined.")
	}

	// Define the nodes (states)
	for _, state := range dfa.States {
		if state == nil {
			panic("Encountered nil state in DFA.")
		}
		shape := "circle"
		if state.IsFinal {
			shape = "doublecircle"
		}
		sb.WriteString(fmt.Sprintf("    \"%s\" [shape=%s];\n", state.Name, shape))
	}

	// Define the transitions
	for fromState, transitions := range dfa.Transitions {
		if fromState == nil {
			panic("Encountered nil fromState in DFA transitions.")
		}
		for symbol, toState := range transitions {
			if toState == nil {
				panic(fmt.Sprintf("Encountered nil toState for transition on symbol '%s' from state '%s'.", symbol, fromState.Name))
			}
			sb.WriteString(fmt.Sprintf("    \"%s\" -> \"%s\" [label=\"%s\"];\n",
				fromState.Name, toState.Name, symbol))
		}
	}

	// Define the start state
	sb.WriteString(fmt.Sprintf("    \"\" [shape=plaintext,label=\"\"];\n"))
	sb.WriteString(fmt.Sprintf("    \"\" -> \"%s\";\n", dfa.StartState.Name))

	sb.WriteString("}\n")

	return sb.String()
}

// getShape returns the shape for the state node based on whether it's a final state.
func getShape(isFinal bool) string {
	if isFinal {
		return "doublecircle"
	}
	return "circle"
}

// GenerateImage generates an image from the DOT representation using Graphviz
func GenerateImageFromDOT(dot string, outputPath string) error {
	cmd := exec.Command("dot", "-Tpng", "-o", outputPath)
	cmd.Stdin = strings.NewReader(dot)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RenderDFA(dfa *DFA, filename string) error {
	DOT := GenerateDOT(dfa)
	err := GenerateImageFromDOT(DOT, filename)
	return err
}
