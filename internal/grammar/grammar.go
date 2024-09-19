package grammar

import (
	"fmt"
	"strings"
)

const Epsilon = "ε"

// Adds a production to a grammar, removing repeated body values.
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

// Given a grammar it removes all epsilon productions
func SimplifyGrammar(grammar *Grammar, printSteps bool) *Grammar {
	if printSteps {
		fmt.Println("1️⃣  Grammar BEFORE simplification")
		fmt.Println(grammar)
		fmt.Println("=================================")
	}
	directNullables := identifyDirectNullables(grammar)
	if printSteps {
		fmt.Println("2️⃣  Direct Nullables found ")
		fmt.Printf("\t%v\n", *directNullables)
		fmt.Println("=================================")
	}
	allNullables := identifyIndirectNullables(grammar, *directNullables)
	if printSteps {
		fmt.Println("3️⃣  All Nullables found ")
		fmt.Printf("\t%v\n", *allNullables)
		fmt.Println("=================================")
	}

	// TODO: make the simplification with all the identified nullables
	return nil
}

// Return a list of all the Direct nullables on the grammar
// Ex: A -> ε
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

// Identify indirect nullables
//
// Returns: List of all nullables (direct an indirect)
func identifyIndirectNullables(grammar *Grammar, nullabes []string) *[]string {

	pastNullables := make([]string, len(nullabes))
	newNullables := make([]string, len(nullabes))
	copy(pastNullables, nullabes)
	copy(newNullables, nullabes)

	// The algorithm works by having a list of PAST_NULLABLES and a list of NEW_NULLABLES
	// 1. On each step the new nullable productions found are added to NEW_NULLABLES
	// 2. If NEW_NULLABLES and PAST_NULLABLES remain equal after one step, it means that
	//    no new production where found, hence ALL NULLABLE PRODUCTION WHERE FOUND
	// 3. Else, repeat step 1 and 2.
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

// ReplaceNullables reemplaza las producciones que contienen símbolos anulables.
func ReplaceNullables(grammar *Grammar, nullables []string) *Grammar {
	newGrammar := make(Grammar)

	// Paso 1. Leer cada body de la gramática por cada head
	for head, productions := range *grammar {
		productionSet := make(map[string]struct{}) // Mapa para rastrear producciones únicas
		var newProductions []string

		// Cola para procesar producciones pendientes
		queue := append([]string{}, productions...)

		// Procesar todas las producciones en la cola
		for len(queue) > 0 {
			production := queue[0]
			queue = queue[1:]

			// Caso 1: No hay símbolos anulables, se añade la producción tal cual si no se ha añadido antes
			if _, exists := productionSet[production]; !exists {
				newProductions = append(newProductions, production)
				productionSet[production] = struct{}{}
			}

			// Caso 2: Existen símbolos anulables en la producción
			for _, nullable := range nullables {
				if strings.Contains(production, nullable) {
					// Generar todas las combinaciones posibles reemplazando el símbolo nullable
					combinations := CombinationNullables(nullable, production)
					for _, newProd := range combinations {
						// Evitar duplicados y procesar nuevas combinaciones
						if _, exists := productionSet[newProd]; !exists {
							newProductions = append(newProductions, newProd)
							productionSet[newProd] = struct{}{}
							queue = append(queue, newProd)
						}
					}
				}
			}
		}

		newGrammar[head] = newProductions
	}

	return &newGrammar
}

// CombinationNullables genera todas las combinaciones posibles al reemplazar el símbolo nullable.
func CombinationNullables(nullable string, baseProduction string) []string {
	var newProductions []string
	chars := []rune(baseProduction)

	// Recorrer la producción y reemplazar el símbolo nullable por epsilon
	for i := 0; i < len(chars); i++ {
		if string(chars[i]) == nullable {
			// Crear una nueva producción con el símbolo reemplazado por epsilon
			newProd := []rune(baseProduction)
			newProd[i] = []rune(Epsilon)[0]
			newProductions = append(newProductions, string(newProd))
		}
	}

	return newProductions
}

// CalculateRemoveSize cuenta cuántas veces aparece un no terminal en una producción.
func CalculateRemoveSize(production string, nonTerminals map[string]struct{}) int {
	count := 0
	for _, char := range production {
		symbol := string(char)
		if _, exists := nonTerminals[symbol]; exists {
			count++
		}
	}
	return count
}

// Revoves duplicates on a slice.
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

// slice: sliceof single character strings,
//
// item: string to check
//
// Returns: true if item is make only by items of slice
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

// Checks if a string exists in a slice
func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
