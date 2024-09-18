package grammar

import (
	"strings"
)

/*
{
	"A" : ["ab", "bb", "Ba"],
	"B" : ["ab", "bb", "Ba"],
	"C" : ["ab", "bb", "epsilon"],
}
*/
// Definition of Grammar. Where the key is the production head, and
// the value is the production's body.
type Grammar map[string][]string

// PrintMap prints a map[string][]string in a readable format.
func (g *Grammar) String() string {
	var sb strings.Builder

	for key, values := range *g {
		sb.WriteString("Key: ")
		sb.WriteString(key)
		sb.WriteString("\n")
		sb.WriteString("\tValues: [")
		sb.WriteString(strings.Join(values, ", "))
		sb.WriteString("]\n")
	}

	return sb.String()
}
