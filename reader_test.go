package main

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	(some? (another-name! :other-key [123
                              12312.1] {
           "stringtest" "another"
         }))
`
	tests := []struct {
		expectedType    Token
		expectedLiteral string
	}{
		{LPAREN, "("},
		{SYM, "some?"},
		{LPAREN, "("},
		{SYM, "another-name!"},
		{KEY, "other-key"},
		{LBRACK, "["},
		{INT, "123"},
		{FLOAT, "12312.1"},
		{RBRACK, "]"},
		{LBRACE, "{"},
		{STR, "stringtest"},
		{STR, "another"},
		{RBRACE, "}"},
		{RPAREN, ")"},
		{RPAREN, ")"},
		{EOF, ""},
	}
	l := NewLexer()
	l.SetInput(input)

	for _, tt := range tests {
		lex := l.NextLexeme()

		if lex.Type != tt.expectedType {
			t.Fatalf("tokentype wrong. expected=%q, got=%q", tt.expectedType, lex.Type)
		}

		if lex.Literal != tt.expectedLiteral {
			t.Fatalf("literal wrong. expected=%s, got=%s", tt.expectedLiteral, lex.Literal)
		}
	}
}
