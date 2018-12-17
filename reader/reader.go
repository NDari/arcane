package reader

import (
	"fmt"
	"strconv"

	"github.com/NDari/arcane/types"
)

type Reader struct {
	*Lexer
}

func NewReader(input string) *Reader {
	l := NewLexer()
	l.SetInput(input)

	return &Reader{
		l,
	}
}

func (l *Reader) ReadAll() (*types.List, error) {
	forms := types.NewList()
	for {
		form, err := l.ReadAny()
		if err != nil {
			return nil, fmt.Errorf("failed to ReadAll:\n%v", err)
		}
		if form == nil {
			return forms, nil
		}

		forms.Append(form)
	}
}

func (l *Reader) ReadAny() (types.Any, error) {
	form := l.NextLexeme()
	switch form.Type {
	case EOF:
		return nil, nil
	case IDENT, STR, SYM, I64, F64:
		s, err := l.ReadAtomLiteral(form)
		if err != nil {
			return nil, fmt.Errorf("failed to read atom:\n%v", err)
		}
		return s, nil
	case LBRACK:
		s, err := l.ReadListLiteral()
		if err != nil {
			return nil, fmt.Errorf("failed to read list:\n%v", err)
		}
		return s, nil
	case LBRACE:
		s, err := l.ReadHashLiteral()
		if err != nil {
			return nil, fmt.Errorf("failed to read hash:\n%v", err)
		}
		return s, nil
	case LPAREN:
		s, err := l.ReadExpr()
		if err != nil {
			return nil, fmt.Errorf("failed to read hash:\n%v", err)
		}
		return s, nil
	default:
		return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
	}
}

func (l *Reader) ReadAtomLiteral(form *Lexeme) (types.Any, error) {
	switch form.Type {
	case IDENT:
		return types.Ident{Val: form.Literal}, nil
	case SYM:
		return types.Sym{Val: form.Literal}, nil
	case STR:
		return types.Str{Val: form.Literal}, nil
	case I64:
		v, err := strconv.Atoi(form.Literal)
		if err != nil {
			return nil, fmt.Errorf("failed to parse I64 from %s:\n%v", form.Literal, err)
		}
		return types.I64{Val: int64(v)}, nil
	case F64:
		v, err := strconv.ParseFloat(form.Literal, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse F64 from %s:\n%v", form.Literal, err)
		}
		return types.F64{Val: v}, nil
	}
	return nil, fmt.Errorf("unknown atom: %s", form.Literal)
}

func (l *Reader) ReadListLiteral() (*types.List, error) {
	lst := types.NewList()
	for form := l.NextLexeme(); form.Type != RBRACK; form = l.NextLexeme() {
		switch form.Type {
		case IDENT, STR, SYM, I64, F64:
			s, err := l.ReadAtomLiteral(form)
			if err != nil {
				return nil, fmt.Errorf("failed to read atom:\n%v", err)
			}
			lst.Append(s)
		case LBRACK:
			s, err := l.ReadListLiteral()
			if err != nil {
				return nil, fmt.Errorf("failed to read list:\n%v", err)
			}
			lst.Append(s)
		case LBRACE:
			s, err := l.ReadHashLiteral()
			if err != nil {
				return nil, fmt.Errorf("failed to read hash:\n%v", err)
			}
			lst.Append(s)
		case LPAREN:
			e, err := l.ReadExpr()
			if err != nil {
				return nil, fmt.Errorf("failed to read expr:\n%v", err)
			}
			lst.Append(e)
		default:
			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
		}
	}
	return lst, nil
}

func (l *Reader) ReadHashLiteral() (*types.Map, error) {
	m := types.NewMap()
	for {
		k, v, err := l.ReadKvPair()
		if err != nil {
			return nil, fmt.Errorf("failed to parse hash literal:\n%v", err)
		}
		if k == "" || v == nil {
			break
		}

		m.Set(k, v)
	}
	return m, nil
}

func (l *Reader) ReadKvPair() (string, types.Any, error) {
	maybeSym := l.NextLexeme()
	if maybeSym.Type == RBRACE {
		return "", nil, nil
	}
	if maybeSym.Type != IDENT {
		return "", nil, fmt.Errorf("When parsing hash literal, expected IDENT, found %s: %s", maybeSym.String(), maybeSym.Literal)
	}
	k := maybeSym.Literal

	v, err := l.ReadAny()
	if err != nil {
		return "", nil, fmt.Errorf("failed to read value associated with key %s:\n%v", k, err)
	}
	return k, v, nil
}

func (l *Reader) ReadExpr() (*types.Map, error) {
	m := types.NewMap()
	funcName := l.NextLexeme()
	if funcName.Type != IDENT {
		return nil, fmt.Errorf("Expected IDENT at the start of an expression, but got %s with value: %s", funcName.String(), funcName.Literal)
	}
	f, err := l.ReadAtomLiteral(funcName)
	if err != nil {
		return nil, fmt.Errorf("failed to read function name in expression:\n%v", err)
	}
	m.Set("$fn", f)
	lst := types.NewList()
	for form := l.NextLexeme(); form.Type != RPAREN; form = l.NextLexeme() {
		switch form.Type {
		case IDENT, STR, SYM, I64, F64:
			s, err := l.ReadAtomLiteral(form)
			if err != nil {
				return nil, fmt.Errorf("failed to read atom:\n%v", err)
			}
			lst.Append(s)
		case LBRACK:
			s, err := l.ReadListLiteral()
			if err != nil {
				return nil, fmt.Errorf("failed to read list:\n%v", err)
			}
			lst.Append(s)
		case LBRACE:
			s, err := l.ReadHashLiteral()
			if err != nil {
				return nil, fmt.Errorf("failed to read hash:\n%v", err)
			}
			lst.Append(s)
		case LPAREN:
			e, err := l.ReadExpr()
			if err != nil {
				return nil, fmt.Errorf("failed to read expr:\n%v", err)
			}
			lst.Append(e)
		default:
			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
		}
	}
	m.Set("$exprs", lst)
	return m, nil
}
