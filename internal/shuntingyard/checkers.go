package shuntingyard

func getTokenInfoIfExist(tokens *[]string, i int) (exist bool, c string, isOperator bool, operator Symbol) {
	if i >= len(*tokens) {
		return false, "", false, nil
	}

	exist = true
	c = (*tokens)[i]
	operator, isOperator = OPERATORS[c]

	return
}

func getSymbolInfoIfExist(tokens *[]Symbol, i int) (exist bool, symbol Symbol, isOperator bool, operator *Operator) {
	if i < 0 || i >= len(*tokens) {
		return false, nil, false, nil
	}

	exist = true
	symbol = (*tokens)[i]
	operator, isOperator = (symbol).(*Operator)

	return
}

func getOperatorIfExist(c string) (*Operator, bool) {
	symbol, isOperator := OPERATORS[c]

	if isOperator {
		operator, _ := symbol.(*Operator)
		return operator, true
	}

	return nil, false

}
