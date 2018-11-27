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
	SYM   // main
	INT   // 12345
	FLOAT // 1.3
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
	SYM:    "SYM",
	INT:    "INT",   // 12345
	FLOAT:  "FLOAT", // 1.3
	STR:    "STR",   // "abc"
	KEY:    "KEY",   // :thing
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
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func NewLexer() *Lexer {
	l := &Lexer{}
	return l
}

func (l *Lexer) SetInput(input string) {
	l.input = input
	l.readPosition = 0
	l.readChar()
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextLexeme() *Lexeme {
	lex := &Lexeme{}

	l.skipWhitespace()

	switch l.ch {
	case '"':
		return l.readString()
	case 0:
		lex.Literal = ""
		lex.Type = EOF
	case ':':
		if isLetter(l.peekChar()) {
			return l.readKey()
		} else {
			lex.Literal = ":"
			lex.Type = SYM
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
			return l.readNumber()
		}
		return l.readSym()
	}

	l.readChar()
	return lex
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readSym() *Lexeme {
	pos := l.position
	l.readChar()
	for isAlpha(l.ch) {
		l.readChar()
	}

	return &Lexeme{
		Type:    SYM,
		Literal: l.input[pos:l.position],
	}
}

func (l *Lexer) readKey() *Lexeme {
	l.readChar() // skip the :
	pos := l.position
	for isAlpha(l.ch) {
		l.readChar()
	}

	return &Lexeme{
		Type:    KEY,
		Literal: l.input[pos:l.position],
	}
}

func (l *Lexer) readNumber() *Lexeme {
	pos := l.position
	for isDigit(l.ch) || (l.ch == '.' && (isDigit(l.peekChar()) || isWhitespace(l.peekChar()))) {
		l.readChar()
	}
	lit := l.input[pos:l.position]
	var t Token
	if strings.Contains(lit, ".") {
		t = FLOAT
	} else {
		t = INT
	}

	return &Lexeme{
		Type:    t,
		Literal: lit,
	}
}

func (l *Lexer) readString() *Lexeme {
	pos := l.position + 1 // skip current char which is "
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}

	l.readChar() // skip over the " we end on
	return &Lexeme{
		Type:    STR,
		Literal: l.input[pos : l.position-1],
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
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
