package shuntingyard

import (
	"fmt"
	"slices"

	"github.com/DanielRasho/Computation-Theory/internal/balancer"
	"github.com/golang-collections/collections/stack"
)

func RegexToPostfix(expresion string, showLogs bool) (string, []Symbol, error) {

	// Check if expresion is balanced
	if isBalanced, _ := balancer.IsBalanced(expresion); !isBalanced {
		return "", nil, fmt.Errorf("expresion is not balanced")
	}
	// Convert tokens to Symbols interface, taking into a account escaped symbols should be Characters
	expresionEscaped, err := convertToSymbols(expresion)
	if err != nil {
		return "", nil, err
	}
	// Try to format the expresion
	expresionFormatted, err := addConcatenationSymbol(&expresionEscaped)
	if err != nil {
		return "", nil, err
	}
	// Interchange Especial operators (?, []) to its equivalents
	expresionPrepared := interchangeOperators(&expresionFormatted)
	// Executing shunting yard
	postFixSymbols := shuntingYard(&expresionPrepared, showLogs)

	finalPostFixExpression := ""
	for _, token := range postFixSymbols {
		finalPostFixExpression += token.GetValue()
	}

	return finalPostFixExpression, postFixSymbols, nil
}

func shuntingYard(tokens *[]Symbol, showLogs bool) []Symbol {
	postfix := make([]Symbol, 0)
	stack := stack.New()

	for i, token := range *tokens {
		_, isOperator := token.(*Operator)
		if token.GetValue() == "(" && isOperator {
			stack.Push(token)
		} else if token.GetValue() == ")" && isOperator {
			for {
				tokenValue, _ := stack.Peek().(Symbol)

				if tokenValue.GetValue() == "(" && isOperator {
					break
				}

				postfix = append(postfix, stack.Pop().(Symbol))
			}
			stack.Pop()
		} else {
			for stack.Len() > 0 {
				peekedChar := stack.Peek().(Symbol)

				if peekedChar.GetPrecedence() >= token.GetPrecedence() {
					postfix = append(postfix, stack.Pop().(Symbol))
				} else {
					break
				}
			}
			stack.Push(token)
		}

		if showLogs {
			fmt.Printf("\t=== STEP %d ====\n", i)
			fmt.Printf("\t\tStack: %+v \n", (*tokens)[0:i+1])
			fmt.Printf("\t\tResponse: %+v \n", postfix)
		}
	}

	for stack.Len() > 0 {
		postfix = append(postfix, stack.Pop().(Symbol))
	}

	return postfix
}

func interchangeOperators(tokens *[]Symbol) []Symbol {

	formattedSymbols := make([]Symbol, 0)

	for i := len(*tokens) - 1; i >= 0; {
		_, s1, s1IsOperator, _ := getSymbolInfoIfExist(tokens, i)
		s2Exist, _, _, _ := getSymbolInfoIfExist(tokens, i-1)

		if !s2Exist {
			formattedSymbols = append(formattedSymbols, s1)
			break
		} else {
			if s1IsOperator && (s1).GetValue() == "?" {

				start, end := getSubExpresionIndex(tokens, i-1, "(", ")")
				// fmt.Printf("star: %d %d\n", start, end)
				subExpresion := (*tokens)[start:end]
				if end-start > 3 { // This means subexpression has parenthesis
					subExpresion = (*tokens)[start+1 : end-1]
					subExpresion = interchangeOperators(&subExpresion)
				}

				slices.Reverse(subExpresion)
				// fmt.Println(subExpresion)

				formattedSymbols = append(formattedSymbols, OPERATORS[")"])
				formattedSymbols = append(formattedSymbols, &Character{value: "ε", precedence: 60})
				formattedSymbols = append(formattedSymbols, OPERATORS["|"])
				formattedSymbols = append(formattedSymbols, OPERATORS[")"])
				formattedSymbols = append(formattedSymbols, subExpresion...)
				formattedSymbols = append(formattedSymbols, OPERATORS["("])
				formattedSymbols = append(formattedSymbols, OPERATORS["("])

				i -= (end - start) + 1
				continue
			} else if s1IsOperator && (s1).GetValue() == "+" {
				start, end := getSubExpresionIndex(tokens, i-1, "(", ")")
				subExpresion := (*tokens)[start:end]
				if end-start > 3 { // This means subexpression has parenthesis
					// fmt.Println("LARGE")
					subExpresion = (*tokens)[start+1 : end-1]
					subExpresion = interchangeOperators(&subExpresion)
				}

				slices.Reverse(subExpresion)

				formattedSymbols = append(formattedSymbols,
					OPERATORS[")"],
					OPERATORS["*"],
					OPERATORS[")"],
				)
				formattedSymbols = append(formattedSymbols, subExpresion...)
				formattedSymbols = append(formattedSymbols,
					OPERATORS["("],
					OPERATORS["·"],
					OPERATORS[")"],
				)
				formattedSymbols = append(formattedSymbols, subExpresion...)
				formattedSymbols = append(formattedSymbols,
					OPERATORS["("],
					OPERATORS["("],
				)
				i -= (end - start) + 1
				continue
			} else if s1IsOperator && (s1).GetValue() == "]" {
				start, end := getSubExpresionIndex(tokens, i, "[", "]")

				// fmt.Println(start)
				// fmt.Println(end)

				if end-start > 3 {
					subExpresion := (*tokens)[start+1 : end-1]

					formattedSymbols = append(formattedSymbols, OPERATORS[")"])
					for j := len(subExpresion) - 1; j > 0; j-- {
						formattedSymbols = append(formattedSymbols, subExpresion[j])
						formattedSymbols = append(formattedSymbols, OPERATORS["|"])
					}
					formattedSymbols = append(formattedSymbols, subExpresion[0])
					formattedSymbols = append(formattedSymbols, OPERATORS["("])
				}

				i -= (end - start)
				continue
			}
		}
		// If any condition raises, just append the caracter
		formattedSymbols = append(formattedSymbols, s1)
		i--
	}
	slices.Reverse(formattedSymbols)
	// fmt.Println(formattedSymbols)

	return formattedSymbols
}

// Given an expresion (a+b)ab, and the index for ")" it will return the
// index of the start and end of the expresion. Ex: start = 0, end = 4
// Start is inclusive
// End is exlusive
func getSubExpresionIndex(tokens *[]Symbol, index int, startSymbol string, endSymbol string) (start, end int) {
	parenthesis := 0
	endExpressionIndex := index + 1
	i := index

	for ; i >= 0; i-- {
		symbol := (*tokens)[i]
		_, isOperator := symbol.(*Operator)

		if isOperator {
			if symbol.GetValue() == endSymbol {
				// fmt.Println("ADD PARENTHESIS")
				parenthesis++
			} else if symbol.GetValue() == startSymbol {
				// fmt.Println("CLOSE PARENTHESIS")
				parenthesis--
			}
		}

		if parenthesis == 0 {
			break
		}
		// fmt.Printf("PASS index = %d value= %s\n", endExpressionIndex, symbol.GetValue())
	}

	return i, endExpressionIndex
}
