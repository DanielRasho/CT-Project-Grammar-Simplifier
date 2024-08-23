package shuntingyard

import (
	"strings"
	"testing"
)

func areSlicesEqual(t *testing.T, response []Symbol, expect []string) {
	value := ""
	for _, v := range response {
		value += v.GetValue()
	}

	if len(response) < len(expect) {
		t.Fatalf("Response has less characters than expected. Has %d, %d given. %s", len(response), len(expect), value)
	} else if len(response) > len(expect) {
		t.Fatalf("Response has more characters than expected. Has %d, %d given. %s", len(response), len(expect), value)
	}
	for i, expected := range expect {
		if response[i].GetValue() != expected {
			t.Fatalf("Characters not match, Given string %s", value)
		}
	}
}

func TestFormatRegex(t *testing.T) {
	symbols, _ := convertToSymbols("c(aa|b)*|b·w")
	response, _ := addConcatenationSymbol(&symbols)
	expect := strings.Split("c·(a·a|b)*|b·w", "")

	areSlicesEqual(t, response, expect)
}

func TestFormatLongRegex(t *testing.T) {
	symbols, _ := convertToSymbols("(a|b?c+|d*e|fgh|i|j)")
	response, _ := addConcatenationSymbol(&symbols)
	expect := strings.Split("(a|b?·c+|d*·e|f·g·h|i|j)", "")
	areSlicesEqual(t, response, expect)
}

func TestFormatRegexWithQuestionOperator(t *testing.T) {
	symbols, _ := convertToSymbols("0?(1?)?0*")
	response, _ := addConcatenationSymbol(&symbols)
	expect := strings.Split("0?·(1?)?·0*", "")
	areSlicesEqual(t, response, expect)
}

func TestFormatRegexWithClasses(t *testing.T) {
	symbols, _ := convertToSymbols("0([abc]|b)0*")
	response, _ := addConcatenationSymbol(&symbols)
	expect := strings.Split("0·([abc]|b)·0*", "")
	areSlicesEqual(t, response, expect)
}

func TestSubExpresion(t *testing.T) {
	symbols, _ := convertToSymbols("(a|(c)*)+a")
	v1, _ := addConcatenationSymbol(&symbols)
	start, end := getSubExpresionIndex(&v1, 7, "(", ")")
	expectStart := 0
	expectEnd := 8

	if start != expectStart || end != expectEnd {
		t.Errorf("Incorrect end or start of subexpresion")
	}

}

func TestInterchangeOperatorsPositiveLock(t *testing.T) {
	symbols, _ := convertToSymbols("(a|(c)*)+")
	v1, _ := addConcatenationSymbol(&symbols)
	response := interchangeOperators(&v1)
	expect := strings.Split("((a|(c)*)·(a|(c)*)*)", "")

	areSlicesEqual(t, response, expect)
}
func TestInterchangeOperatorsZeroOrMoreInstance(t *testing.T) {
	symbols, _ := convertToSymbols("(a|(c)*)?")
	v1, _ := addConcatenationSymbol(&symbols)
	response := interchangeOperators(&v1)
	expect := strings.Split("((a|(c)*)|ε)", "")

	value := ""
	for _, v := range response {
		value += v.GetValue()
	}

	if len(response) < len(expect) {
		t.Fatalf("Response has less characters than expected. Has %d, %d given. %s", len(response), len(expect), value)
	} else if len(response) > len(expect) {
		t.Fatalf("Response has more characters than expected. Has %d, %d given. %s", len(response), len(expect), value)
	}
	for i, expected := range expect {
		if response[i].GetValue() != expected {
			t.Fatalf("Characters not match, Given string %s", value)
		}
	}
}

func TestInterchangeOperatorsClass(t *testing.T) {
	symbols, _ := convertToSymbols("a([abc1234])?|j")
	v1, _ := addConcatenationSymbol(&symbols)
	response := interchangeOperators(&v1)
	expect := strings.Split("a·(((a|b|c|1|2|3|4))|ε)|j", "")

	areSlicesEqual(t, response, expect)
}

func TestShuntinYard(t *testing.T) {
	symbols, _ := convertToSymbols("(aa|b)|abb*")
	v1 := interchangeOperators(&symbols)
	v2, _ := addConcatenationSymbol(&v1)
	response := shuntingYard(&v2, false)
	expect := strings.Split("aa·b|ab·b*·|", "")

	areSlicesEqual(t, response, expect)
}
