package shuntingyard

import (
	"fmt"
	"strings"
)

func convertToSymbols(expresion string) ([]Symbol, error) {
	stringTokens := strings.Split(expresion, "")
	symbols := make([]Symbol, 0)

	for i := 0; i < len(stringTokens); {
		_, c1, c1IsOperator, c1Operator := getTokenInfoIfExist(&stringTokens, i)
		c2Exist, c2, c2IsOperator, _ := getTokenInfoIfExist(&stringTokens, i+1)

		// HANDLE ESCAPED SYMBOLS
		if c1 == ESCAPE_SYMBOL {
			if c2Exist {

				// Get info about if C2 is valid escape sequence
				c2Escaped, c2IsEscaped := ESCAPED_CHARACTERS[c2]

				// Escape value
				if c2IsOperator {
					symbols = append(symbols, &Character{value: c2, precedence: 60})
				} else if c2IsEscaped {
					symbols = append(symbols, c2Escaped)
				}
				// fmt.Printf("\t\tStack: %+v \n", (formattedTokens)[0:i+1])
				i += 2
				continue
			} else {
				return nil, fmt.Errorf("invalid escaped symbol")
			}
		}
		// APPEND C1 TOKEN
		if c1IsOperator {
			symbols = append(symbols, c1Operator)
		} else {
			symbols = append(symbols, &Character{value: c1, precedence: 60})
		}
		// fmt.Printf("\t\tStack: %+v \n", (formattedTokens)[0:i+1])
		i++

	}
	return symbols, nil
}

func addConcatenationSymbol(expresion *[]Symbol) ([]Symbol, error) {

	formattedTokens := make([]Symbol, 0)

	for i := 0; i < len(*expresion); {

		_, s1, s1IsOperator, _ := getSymbolInfoIfExist(expresion, i)
		s2Exist, s2, _, _ := getSymbolInfoIfExist(expresion, i+1)

		// SPECIAL CASE, if Class sctructure encontared skip([abc])
		if s1.GetValue() == "[" && s1IsOperator {
			newIndex := i
			// Search for the closing class bracket "]"
			for ; newIndex < len(*expresion); newIndex++ {
				_, step, stepIsOperator, _ := getSymbolInfoIfExist(expresion, newIndex)

				if step.GetValue() == "]" && stepIsOperator {
					break
				}

				formattedTokens = append(formattedTokens, step)

			}
			i = newIndex // To start with the next symbol after the class
			continue
		}

		formattedTokens = append(formattedTokens, s1)

		// ADD CONCAT SYMBOL IF NEXT CHARACTER NEEDS IT
		if s2Exist && ShouldAddConcatenationOperator(s1, s2) {
			formattedTokens = append(formattedTokens, OPERATORS[CONCAT_SYMBOL])
		}

		// fmt.Printf("\t\tStack: %+v \n", (formattedTokens)[0:i+1])
		i++
	}
	// fmt.Printf("\t\tStack: %+v \n", formattedTokens)
	return formattedTokens, nil
}

func ShouldAddConcatenationOperator(s1, s2 Symbol) bool {
	s1Operator, s1IsOperator := s1.(*Operator)
	_, s2IsOperator := s2.(*Operator)

	// fmt.Printf("%s %s\n", c1, c2)

	if s2 == nil {
		return false
	}

	// If both are open or close parenthesis, false
	if (s1IsOperator && s2IsOperator) &&
		((s1.GetValue() == "(" && s2.GetValue() == "(") ||
			(s1.GetValue() == ")" && s2.GetValue() == ")")) {
		return false
	}

	// If the S1 is Operator and
	// 	just and need more that 1 operands,
	// 	is an open parenthesis
	// 	nees less than and operand but the next character is an operator
	if s1IsOperator {
		if s1Operator.GetOperands() > 1 ||
			(s1.GetValue() == "(" && !s2IsOperator) ||
			(s1Operator.GetOperands() < 1 && s2IsOperator) {
			return false
		}
	}
	// 	If S2 is an "(" operator
	if s2IsOperator &&
		((s2.GetValue() == "(") ||
			(s2.GetValue() == "[")) {
		return true
	}
	if s2IsOperator { // If c1 is not operand then
		return false
	}

	return true
}
