package nfa

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GenerateDOT generates a .dot representation of an NFA as a string.
func GenerateDOT(nfa *NFA) string {
	var sb strings.Builder

	// Write the Graphviz dot header
	sb.WriteString("digraph NFA {\nrankdir=LR;\n")

	// fmt.Printf("\tNUM TRANSITION: %d\n", len(nfa.Transitions))

	// Define the nodes (states)
	for _, transition := range nfa.Transitions {
		sb.WriteString(fmt.Sprintf("    \"%s\" [shape=%s];\n",
			transition.From.Name,
			getShape(transition.From.IsFinal)))

		for _, toState := range transition.To {
			sb.WriteString(fmt.Sprintf("    \"%s\" [shape=%s];\n",
				toState.Name,
				getShape(toState.IsFinal)))
			// fmt.Printf("\t\t [%s] ===> [%s]\n", transition.From.Name, toState.Name)
		}
	}

	// Define the transitions
	for _, transition := range nfa.Transitions {
		for _, toState := range transition.To {
			sb.WriteString(fmt.Sprintf("    \"%s\" -> \"%s\" [label=\"%s\"];\n",
				transition.From.Name, toState.Name, transition.Symbol))
		}
	}

	// Define the start state
	sb.WriteString(fmt.Sprintf("    \"start\" [shape=plaintext,label=\"\"]; \"start\" -> \"%s\";\n",
		nfa.StartState.Name))

	// Optionally define the end state
	if nfa.EndState != nil {
		sb.WriteString(fmt.Sprintf("    \"%s\" [shape=doublecircle];\n", nfa.EndState.Name))
	}

	// Close the dot file
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

func RenderAFN(nfa *NFA, filename string) error {
	DOT := GenerateDOT(nfa)
	err := GenerateImageFromDOT(DOT, filename)
	return err
}
