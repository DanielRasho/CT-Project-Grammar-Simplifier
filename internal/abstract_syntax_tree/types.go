/*
El AST se compone de nodos que pueden ser operadores o caracteres. Los operadores pueden tener varios operandos, mientras que los caracteres son nodos terminales sin operandos.
*/

package abstract_syntax_tree

/*
Interfax Node representa un nodo en el AST que tiene un método String con retorno String.
*/
type Node interface {
	String() string
}

/*
CharacterNode representa un carácter en el AST.
*/
type CharacterNode struct {
	Value string
}

/*
OperatorNode representa un operador en el AST.
*/
type OperatorNode struct {
	Value    string
	Operands []Node
}

/*
String devuelve el valor del CharacterNode como una cadena.
*/
func (c *CharacterNode) String() string {
	return c.Value
}

/*
String devuelve el valor del OperatorNode como una cadena.
*/
func (o *OperatorNode) String() string {
	return o.Value
}

/*
GetOperands devuelve los operandos del OperatorNode.
*/
func (o *OperatorNode) GetOperands() []Node {
	return o.Operands
}

/*
NewCharacterNode crea un nuevo CharacterNode.
Parámetros:
  - value: El valor del carácter como una cadena.

Retorno:
  - Un puntero a un nuevo CharacterNode con el valor especificado.
*/
func NewCharacterNode(value string) *CharacterNode {
	return &CharacterNode{Value: value}
}

/*
NewOperatorNode crea un nuevo OperatorNode.
Parámetros:
  - value: El valor del operador como una cadena.
  - operands: Una lista de nodos que son los operandos del operador.

Retorno:
  - Un puntero a un nuevo OperatorNode con el valor y los operandos especificados.
*/
func NewOperatorNode(value string, operands []Node) *OperatorNode {
	return &OperatorNode{Value: value, Operands: operands}
}
