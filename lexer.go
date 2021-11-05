package html_parser

import (
	"unicode"
)

func NewLexer(src string) Lexer {
	return Lexer{
		pos:   0,
		runes: []rune(src),
	}
}

type Lexer struct {
	pos   int
	runes []rune
}

func (lx *Lexer) curt() rune {
	return lx.runes[lx.pos]
}
func (lx *Lexer) next() rune {
	return lx.runes[lx.pos+1]
}

func (lx *Lexer) advance() {
	lx.pos++
}

func (lx *Lexer) isEof() bool {
	return lx.pos >= len(lx.runes)
}

func (lx *Lexer) consume(r rune) bool {
	if lx.curt() == r {
		lx.advance()
		return true
	}
	return false
}

func (lx *Lexer) consumeLiteral() Token {
	lit := ""
	s := lx.pos
	for !lx.isEof() {
		c := lx.curt()
		if unicode.IsLetter(c) {
			lit += string(c)
			lx.advance()
		} else {
			break
		}
	}
	e := lx.pos
	return Token{
		kind:     Literal,
		lit:      lit,
		position: [2]int{s, e},
	}
}

func (lx *Lexer) consumeNumber() Token {
	lit := ""
	kind := Int
	s := lx.pos
	for !lx.isEof() {
		c := lx.curt()
		if unicode.IsDigit(c) || c == '-' || c == '.' {
			lit += string(c)
			lx.advance()
			if c == '.' {
				kind = Float
			}
		} else {
			break
		}
	}
	e := lx.pos

	return Token{
		kind:     kind,
		lit:      lit,
		position: [2]int{s, e},
	}
}

func (lx *Lexer) consumeSpace() Token {
	lit := ""
	s := lx.pos
	for !lx.isEof() {
		c := lx.curt()
		if unicode.IsSpace(c) {
			lit += string(c)
			lx.advance()
		} else {
			break
		}
	}
	e := lx.pos
	return Token{
		lit:      lit,
		kind:     Space,
		position: [2]int{s, e},
	}
}

func (lx *Lexer) consumeSymbol() Token {
	lit := ""
	s := lx.pos
	lit += string(lx.curt())
	lx.advance()
	e := lx.pos
	return Token{
		lit:      lit,
		kind:     Symbol,
		position: [2]int{s, e},
	}
}

func (lx *Lexer) Lex() ([]Token, error) {
	var tokens []Token
	for !lx.isEof() {
		c := lx.curt()
		var tok Token
		if unicode.IsLetter(c) {
			tok = lx.consumeLiteral()
		} else if (c == '-' && unicode.IsDigit(lx.next())) || unicode.IsDigit(c) {
			tok = lx.consumeNumber()
		} else if unicode.IsSpace(c) {
			tok = lx.consumeSpace()
		} else if unicode.IsSymbol(c) {
			tok = lx.consumeSymbol()
		} else {
			return nil, UnexpectedRuneError(c)
		}
		tokens = append(tokens, tok)
	}
	return tokens, nil
}
