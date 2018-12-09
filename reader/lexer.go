package reader

import (
	"fmt"
	"strings"
)

type Token int

const (
	ILLIGAL Token = iota
	EOF

	litTokenStart
	IDENT // main
	I64   // 12345
	F64   // 1.3
	STR   // "abc"
	KEY   // :thing
	FN    // #{  }
	litTokenEnd

	groupingStart
	LPAREN
	RPAREN
	LBRACK
	RBRACK
	LBRACE
	RBRACE
	groupingEnd
)

var tokenNames = [...]string{
	EOF:    "EOF",
	I64:    "I64", // 12345
	F64:    "F64", // 1.3
	STR:    "STR", // "abc"
	KEY:    "KEY", // :thing
	LPAREN: "LPAREN",
	RPAREN: "RPAREN",
	LBRACK: "LBRACK",
	RBRACK: "RBRACK",
	LBRACE: "LBRACE",
	RBRACE: "RBRACE",
}

type Lexeme struct {
	Type    Token
	Literal string
}

func (l *Lexeme) String() string {
	tname := tokenNames[l.Type]
	return fmt.Sprintf("%s:\t\t%s", tname, l.Literal)
}

type Lexer struct {
	input       string
	position    int  // current position in input (points to current char)
	lexPosition int  // next char to be lexed (after current char)
	ch          byte // current char under examination
}

func NewLexer() *Lexer {
	l := &Lexer{}
	return l
}

func (l *Lexer) SetInput(input string) {
	l.input = input
	l.lexPosition = 0
	l.lexChar()
}

func (l *Lexer) lexChar() {
	if l.lexPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.lexPosition]
	}

	l.position = l.lexPosition
	l.lexPosition++
}

func (l *Lexer) NextLexeme() *Lexeme {
	lex := &Lexeme{}

	l.skipWhitespace()

	switch l.ch {
	case '"':
		return l.lexString()
	case 0:
		lex.Literal = ""
		lex.Type = EOF
	case ':':
		if isLetter(l.peekChar()) {
			return l.lexKey()
		} else {
			lex.Literal = ":"
			lex.Type = IDENT
		}
	case '(':
		lex.Literal = "("
		lex.Type = LPAREN
	case ')':
		lex.Literal = ")"
		lex.Type = RPAREN
	case '[':
		lex.Literal = "["
		lex.Type = LBRACK
	case ']':
		lex.Literal = "]"
		lex.Type = RBRACK
	case '{':
		lex.Literal = "{"
		lex.Type = LBRACE
	case '}':
		lex.Literal = "}"
		lex.Type = RBRACE
	default:
		if isDigit(l.ch) {
			return l.lexNumber()
		}
		return l.lexIdent()
	}

	l.lexChar()
	return lex
}

func (l *Lexer) peekChar() byte {
	if l.lexPosition >= len(l.input) {
		return 0
	}

	return l.input[l.lexPosition]
}

func (l *Lexer) lexIdent() *Lexeme {
	pos := l.position
	l.lexChar()
	for isAlpha(l.ch) {
		l.lexChar()
	}

	return &Lexeme{
		Type:    IDENT,
		Literal: l.input[pos:l.position],
	}
}

func (l *Lexer) lexKey() *Lexeme {
	l.lexChar() // skip the :
	pos := l.position
	for isAlpha(l.ch) {
		l.lexChar()
	}

	return &Lexeme{
		Type:    KEY,
		Literal: l.input[pos:l.position],
	}
}

func (l *Lexer) lexNumber() *Lexeme {
	pos := l.position
	for isDigit(l.ch) || (l.ch == '.' && (isDigit(l.peekChar()) || isWhitespace(l.peekChar()))) {
		l.lexChar()
	}
	lit := l.input[pos:l.position]
	var t Token
	if strings.Contains(lit, ".") {
		t = F64
	} else {
		t = I64
	}

	return &Lexeme{
		Type:    t,
		Literal: lit,
	}
}

func (l *Lexer) lexString() *Lexeme {
	pos := l.position + 1 // skip current char which is "
	for {
		l.lexChar()
		if l.ch == '"' {
			break
		}
	}

	l.lexChar() // skip over the " we end on
	return &Lexeme{
		Type:    STR,
		Literal: l.input[pos : l.position-1],
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.lexChar()
	}
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isAlpha(ch byte) bool {
	return isLetter(ch) || ch == '-' || ch == '?' || ch == '!'
}
