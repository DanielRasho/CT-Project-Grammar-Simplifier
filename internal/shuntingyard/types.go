package shuntingyard

type Symbol interface {
	GetValue() string
	GetPrecedence() int
	String() string
}

type Character struct {
	value      string
	precedence int
}

func (c *Character) GetValue() string {
	return c.value
}

func (c *Character) GetPrecedence() int {
	return c.precedence
}

func (c *Character) String() string {
	return c.value
}

type Operator struct {
	value      string
	precedence int
	operands   int
}

func (c *Operator) GetValue() string {
	return c.value
}

func (c *Operator) GetPrecedence() int {
	return c.precedence
}

func (c *Operator) GetOperands() int {
	return c.operands
}

func (c *Operator) String() string {
	return c.value
}

const ESCAPE_SYMBOL string = "\\"
const CONCAT_SYMBOL string = "·"

var OPERATORS = map[string]Symbol{
	")": &Operator{value: ")", precedence: 10, operands: 1},
	"(": &Operator{value: "(", precedence: 10, operands: 0},
	"]": &Operator{value: "]", precedence: 10, operands: 1},
	"[": &Operator{value: "[", precedence: 10, operands: 0},
	"|": &Operator{value: "|", precedence: 20, operands: 2},
	"·": &Operator{value: "·", precedence: 30, operands: 2},
	"?": &Operator{value: "?", precedence: 40, operands: 1},
	"*": &Operator{value: "*", precedence: 40, operands: 1},
	"+": &Operator{value: "+", precedence: 40, operands: 1},
	"^": &Operator{value: "^", precedence: 50, operands: 2},
}

var ESCAPED_CHARACTERS = map[string]Symbol{
	"ε":  &Character{value: "", precedence: 60},
	"\\": &Character{value: "\\", precedence: 60},
	"n":  &Character{value: "\n", precedence: 60},
	"{":  &Character{value: "{", precedence: 60},
	"}":  &Character{value: "}", precedence: 60},
}
