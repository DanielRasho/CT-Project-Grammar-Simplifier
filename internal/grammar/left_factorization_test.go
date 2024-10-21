package grammar

import (
	"fmt"
	"testing"
)

func TestFindLongestPrefix(t *testing.T) {
	// Step 2: Define the grammar

	bodies := [][]Symbol{
		{
			Symbol{value: "a", isTerminal: true},
			Symbol{value: "b", isTerminal: true},
			Symbol{value: "M", isTerminal: false, id: 0},
		},
		{
			Symbol{value: "a", isTerminal: true},
			Symbol{value: "b", isTerminal: true},
			Symbol{value: "M", isTerminal: false, id: 0},
		},
		{
			Symbol{value: "j", isTerminal: true},
			Symbol{value: "K", isTerminal: false, id: 0},
		},
		{
			Symbol{value: "a", isTerminal: true},
			Symbol{value: "b", isTerminal: true},
			Symbol{value: "b", isTerminal: true},
		},
		{
			Symbol{value: "j", isTerminal: true},
			Symbol{value: "l", isTerminal: true},
		},
	}

	// fmt.Println(grammar.String(false))
	prefix, remains, _ := findLongestCommonPrefix(bodies)

	fmt.Println(prefix)
	fmt.Println("==========")

	for _, v := range remains {
		fmt.Println(v)
	}
}

func TestFindLongestPrefixNothingInCommon(t *testing.T) {
	// Step 2: Define the grammar

	bodies := [][]Symbol{
		{
			Symbol{value: "a", isTerminal: true},
		},
		{
			Symbol{value: "b", isTerminal: true},
		},
		{
			Symbol{value: "c", isTerminal: true},
		},
	}

	// fmt.Println(grammar.String(false))
	prefix, _, _ := findLongestCommonPrefix(bodies)

	if prefix == nil {
		fmt.Println("SUCCESS")
	}
	/*
		fmt.Println(prefix)
		fmt.Println("==========")

		for _, v := range remains {
			fmt.Println(v)
		}*/
}

func TestApplyFactorization(t *testing.T) {
	// Step 2: Define the grammar
	grammar := &Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}
	grammar.AddProductionFromString("A -> {B}x|jk|{B}b|jl")
	A := Symbol{value: "A", isTerminal: false, id: 0}

	result := leftFactor(grammar, &A, grammar.Productions[A])

	grammar.SetProductionBodies(A, result)

	// Expect to receive:
	// {A_0} -> j{A_1}|{B_0}{A_2}
	// {A_1} -> k|l
	// {A_2} -> x|b
	fmt.Println(grammar.String(false))
}
func TestApplyFactorization2(t *testing.T) {
	// Step 2: Define the grammar
	grammar := &Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}
	grammar.AddProductionFromString("A -> {B}x|{B}")
	A := Symbol{value: "A", isTerminal: false, id: 0}

	result := leftFactor(grammar, &A, grammar.Productions[A])

	grammar.SetProductionBodies(A, result)

	// Expect to receive:
	// {A_0} -> j{A_1}|{B_0}{A_2}
	// {A_1} -> k|l
	// {A_2} -> x|b
	fmt.Println(grammar.String(false))
}

func TestFactorizeGrammar(t *testing.T) {
	// Step 2: Define the grammar
	grammar := &Grammar{
		Productions: make(map[Symbol][][]Symbol),
	}
	grammar.AddProductionFromString("A -> {B}x|jk|{B}b|jl")
	grammar.AddProductionFromString("B -> mm|mb|m")

	a := factorizeGrammar(grammar)

	// Expect to receive:
	// {A_0} -> j{A_1}|{B_0}{A_2}
	// {A_1} -> k|l
	// {A_2} -> x|b
	fmt.Println(a.String(true))
}
