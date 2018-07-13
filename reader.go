package arcane

type Lexeme struct {
	Type    Token
	Literal string
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

func (l *Lexer) NextLexeme() Lexeme {
	var lex Lexeme

	l.skipWhitespace()

	switch l.ch {
	case '"':
		lex.Literal = l.readString()
		lex.Type = STR
	case 0:
		lex.Literal = ""
		lex.Type = EOF
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
			lex.Type = NUM
			lex.Literal = l.readNumber()
			return lex
		} 
		lex.Literal = l.readSym()
		lex.Type = SYM
		return lex
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

func (l *Lexer) readSym() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}

	return l.input[position:l.position]
}

func newLexeme(token Token, ch byte) Lexeme {
	return Lexeme{
		Type:    token,
		Literal: string(ch),
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
