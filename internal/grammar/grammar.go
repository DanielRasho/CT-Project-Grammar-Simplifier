package grammar

import (
	"strings"
)

func (g *Grammar) AddProduction(production string) {
	// Since a production has the shape A -> a|bC
	// There are 2 divisions between the Head, Arrow, And Body.
	division1 := strings.Index(production, " ")                               // Find first space index
	division2 := division1 + 1 + strings.Index(production[division1+1:], " ") // Find second space index

	head := production[:division1]
	body := production[division2+1:]
	bodyItems := strings.Split(body, "|")

	// If production is not registered create it
	if _, exist := (*g)[head]; !exist {
		(*g)[head] = removeDuplicates(bodyItems)
	} else {
		// Else append the body new items with the old ones
		existentBodyItems := (*g)[head]
		(*g)[head] = removeDuplicates(append(existentBodyItems, bodyItems...))
	}
}

func SimplifyGrammar(grammar *Grammar) *Grammar {
	// directNullables := identifyDirectNullables(grammar)
	// allNullables := identifyIndirectNullables(grammar, *directNullables)

	// TODO: make the simplification with all the identified nullables
	return nil
}

func identifyDirectNullables(grammar *Grammar) *[]string {

	directNullables := make([]string, 0, 3)

	for head, body := range *grammar {
		for _, v := range body {
			if v == "ε" {
				directNullables = append(directNullables, head)
				continue
			}
		}
	}
	return &directNullables
}

func identifyIndirectNullables(grammar *Grammar, nullabes []string) *[]string {

	pastNullables := make([]string, len(nullabes))
	newNullables := make([]string, len(nullabes))
	copy(pastNullables, nullabes)
	copy(newNullables, nullabes)

	for {
		for head, bodies := range *grammar {
			// If the production is already nullable dont analize it.
			if contains(pastNullables, head) {
				continue
			}
			// else check if it is nullable and add it to the NewNullables list
			for _, body := range bodies {
				isNullable := isComposedOf(pastNullables, body)
				if isNullable {
					newNullables = append(newNullables, head)
					break
				}
			}
		}

		if len(pastNullables) == len(newNullables) {
			break
		}

		pastNullables = newNullables
	}

	return &newNullables
}

func replaceNullables(grammar *Grammar, nullables []string) *Grammar {

	return nil
}

func removeEpsilons(grammar *Grammar) *Grammar {

	return nil
}

func removeDuplicates(slice []string) []string {
	uniqueMap := make(map[string]bool)
	var result []string

	for _, item := range slice {
		if _, exists := uniqueMap[item]; !exists {
			uniqueMap[item] = true
			result = append(result, item)
		}
	}

	return result
}

func isComposedOf(slice []string, item string) bool {
	// Create a set for quick lookup
	set := make(map[string]struct{}, len(slice))
	for _, char := range slice {
		set[char] = struct{}{}
	}

	// Check if every character in the item is in the set
	for _, char := range item {
		if _, exists := set[string(char)]; !exists {
			return false
		}
	}
	return true
}

// Function to check if a string exists in a slice
func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
