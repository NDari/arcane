package arcane

type Token int

const (
	ILLIGAL Token = iota
	EOF
	COMMENT

	litTokenStart
	SYM // main
	I64 // 12345
	F64 // 1.3
	STR // "abc"
	KEY // :thing
	litTokenEnd

	sigilStart
	LPAREN
	RPAREN
	LBRACK
	RBRACK
	LBRACE
	RBRACE
	sigilEnd
)

func (tok Token) IsLiteral() bool {
	return litTokenStart < tok && tok < litTokenEnd
}

func (tok Token) IsSigil() bool {
	return sigilStart < tok && tok < sigilEnd
}
