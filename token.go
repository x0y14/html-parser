package html_parser

import "fmt"

type TokenKind int

const (
	_ TokenKind = iota
	String
	Int
	Float

	Literal
	Symbol

	Space
)

var tokenKind = [...]string{
	String:  "String",
	Int:     "Int",
	Float:   "Float",
	Literal: "Literal",
	Symbol:  "Symbol",
	Space:   "Space",
}

type Token struct {
	lit      string
	kind     TokenKind
	position [2]int
}

func (tok *Token) String() string {
	return fmt.Sprintf("Token{ kind: %v, lit: %v, pos: %v-%v }", tokenKind[tok.kind], tok.lit, tok.position[0], tok.position[1])
}
