package html_parser

import (
	"fmt"
	"testing"
)

func TestLexer_Lex(t *testing.T) {
	var tests = []struct {
		title string
		in    string
	}{
		{
			"a",
			"a",
		},
		{
			"img",
			"<img src='https://example.com' />",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			lexer := NewLexer(test.in)
			tokens, err := lexer.Lex()
			if err != nil {
				t.Fatal(err)
			}
			for _, tok := range tokens {
				fmt.Printf("%v\n", tok.String())
			}
		})
	}
}
