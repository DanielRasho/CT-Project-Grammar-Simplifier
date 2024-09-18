package grammar

import "testing"

func areSlicesEqual(t *testing.T, response []string, expect []string) {
	value := ""
	for _, v := range response {
		value += v + " "
	}

	if len(response) < len(expect) {
		t.Fatalf("Response has less characters than expected. Has %d, %d given. %s", len(response), len(expect), value)
	} else if len(response) > len(expect) {
		t.Fatalf("Response has more characters than expected. Has %d, %d given. %s", len(response), len(expect), value)
	}
	for i, expected := range expect {
		if response[i] != expected {
			t.Fatalf("Characters not match, Given string %s", value)
		}
	}
}

func TestAddProductionToEmptyGrammar(t *testing.T) {
	grammar := make(Grammar)
	grammar.AddProduction("A -> a|a|bc|C")
	expectedBodyItems := []string{"a", "bc", "C"}

	if _, exist := grammar["A"]; !exist {
		t.Fatalf("A production was not appended to the grammar\n")
	}
	areSlicesEqual(t, grammar["A"], expectedBodyItems)
}
func TestAddProductionToNotEmptyGrammar(t *testing.T) {
	grammar := make(Grammar)
	grammar.AddProduction("A -> a|a|bc|C")
	grammar.AddProduction("A -> a|B|J")
	expectedBodyItems := []string{"a", "b", "C", "B", "J"}

	if _, exist := grammar["A"]; !exist {
		t.Fatalf("A production was not appended to the grammar\n")
	}
	areSlicesEqual(t, grammar["A"], expectedBodyItems)
}
