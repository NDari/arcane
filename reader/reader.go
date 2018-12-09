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

func (l *Reader) ReadAll() ([]types.Any, error) {
	var forms []types.Any
	for {
		form, err := l.ReadAny()
		if err != nil {
			return nil, fmt.Errorf("failed to ReadAll: %v", err)
		}
		if form == nil {
			return forms, nil
		}

		forms = append(forms, form)
	}
}

func (l *Reader) ReadAny() (types.Any, error) {
	form := l.NextLexeme()
	switch form.Type {
	case EOF:
		return nil, nil
	case IDENT, STR, KEY, I64, F64:
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
	default:
		return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
	}
}

func (l *Reader) ReadAtomLiteral(form *Lexeme) (types.Any, error) {
	switch form.Type {
	case IDENT, KEY:
		return types.Key{Val: form.Literal}, nil
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
		case IDENT, STR, KEY, I64, F64:
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
		default:
			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
		}
	}
	return lst, nil
}

func (l *Reader) ReadHashLiteral() (*types.Map, error) {
	m := types.NewMap()
	for {
		pair, err := l.ReadKvPair()
		if err != nil {
			return nil, fmt.Errorf("failed to parse hash literal: %v", err)
		}
		if pair.IsEmpty() {
			break
		}

		m.Set(pair)
	}
	return m, nil
}

func (l *Reader) ReadKvPair() (*types.List, error) {
	maybeKey := l.NextLexeme()
	if maybeKey.Type == RBRACE {
		return types.NewList(), nil
	}
	if maybeKey.Type != KEY {
		return nil, fmt.Errorf("When parsing hash literal, expected Key, found %s: %s", maybeKey.String(), maybeKey.Literal)
	}
	atm, err := l.ReadAtomLiteral(maybeKey)
	if err != nil {
		return nil, fmt.Errorf("could not parse key %s: %v", maybeKey.Literal, err)
	}

	k, ok := atm.(types.Key)
	if !ok {
		return nil, fmt.Errorf("could not convert %v to key", atm)
	}

	v, err := l.ReadAny()
	if err != nil {
		return nil, fmt.Errorf("failed to read value associated with key %v: %v", k.Repr(), err)
	}
	lst := types.NewList(k, v)
	return lst, nil
}
