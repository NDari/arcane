package main

import (
	"fmt"
	"strconv"
)

func Read(s Str) ([]*Any, error) {
	var forms []*Any
	l := NewLexer()
	l.SetInput(s.val)
	for form := l.NextLexeme(); form.Type != EOF; form = l.NextLexeme() {
		switch form.Type {
		case SYM, STR, KEY, INT, FLOAT:
			s, err := ReadAtomLiteral(form)
			if err != nil {
				return nil, fmt.Errorf("failed to read atom:\n%v", err)
			}
			forms = append(forms, s)
		case LPAREN:
			s, err := ReadListLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read list:\n%v", err)
			}
			forms = append(forms, s)
		case LBRACE:
			s, err := ReadHashLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read hashmap:\n%v", err)
			}
			forms = append(forms, s)
		case LBRACK:
			s, err := ReadVecLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read vector:\n%v", err)
			}
			forms = append(forms, s)
		default:
			return nil, fmt.Errorf("could not read \"%s\" of type %s", form.Literal, tokenNames[form.Type])
		}
	}
	return forms, nil
}

func ReadAtomLiteral(form *Lexeme) (*Any, error) {
	var s Any
	switch form.Type {
	case SYM:
		s = &Sym{form.Literal}
	case STR:
		s = &Str{form.Literal}
	case KEY:
		s = &Key{form.Literal}
	case INT:
		v, err := strconv.Atoi(form.Literal)
		if err != nil {
			return nil, fmt.Errorf("failed to parse I64 from %s:\n%v", form.Literal, err)
		}
		s = &I64{int64(v)}
	case FLOAT:
		v, err := strconv.ParseFloat(form.Literal, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse F64 from %s:\n%v", form.Literal, err)
		}
		s = &F64{v}
	case EOF:
		return nil, fmt.Errorf("undexpected EOF")
	}
	return &s, nil
}

func ReadListLiteral(l *Lexer) (*Any, error) {
	var lst Any = &List{}
	for form := l.NextLexeme(); form.Type != RPAREN; form = l.NextLexeme() {
		switch form.Type {
		case SYM, STR, KEY, INT, FLOAT:
			s, err := ReadAtomLiteral(form)
			if err != nil {
				return nil, fmt.Errorf("failed to read atom:\n%v", err)
			}
			lst = &List{
				s,
				lst.(*List),
			}
		case LPAREN:
			s, err := ReadListLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read list:\n%v", err)
			}
			lst = &List{
				s,
				lst.(*List),
			}
		case LBRACE:
			s, err := ReadHashLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read hashmap:\n%v", err)
			}
			lst = &List{
				s,
				lst.(*List),
			}
		case LBRACK:
			s, err := ReadVecLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read vector:\n%v", err)
			}
			lst = &List{
				s,
				lst.(*List),
			}
		default:
			return nil, fmt.Errorf("could not read \"%s\" of type %s", form.Literal, tokenNames[form.Type])
		}
	}
	return &lst, nil
}

func ReadVecLiteral(l *Lexer) (*Any, error) {
	var v Any = &Vec{
		make([]*Any, 0),
	}
	for form := l.NextLexeme(); form.Type != RBRACK; form = l.NextLexeme() {
		switch form.Type {
		case SYM, STR, KEY, INT, FLOAT:
			s, err := ReadAtomLiteral(form)
			if err != nil {
				return nil, fmt.Errorf("failed to read atom:\n%v", err)
			}
			v.(*Vec).vals = append(v.(*Vec).vals, s)
		case LPAREN:
			s, err := ReadListLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read list:\n%v", err)
			}
			v.(*Vec).vals = append(v.(*Vec).vals, s)
		case LBRACE:
			s, err := ReadHashLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read hashmap:\n%v", err)
			}
			v.(*Vec).vals = append(v.(*Vec).vals, s)
		case LBRACK:
			s, err := ReadVecLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read vector:\n%v", err)
			}
			v.(*Vec).vals = append(v.(*Vec).vals, s)
		default:
			return nil, fmt.Errorf("could not read \"%s\" of type %s", form.Literal, tokenNames[form.Type])
		}
	}
	return &v, nil
}

func ReadHashLiteral(l *Lexer) (*Any, error) {
	var h Any = &HashMap{
		make(map[*Key]*Any),
	}
	for form := l.NextLexeme(); form.Type != RBRACE; form = l.NextLexeme() {
		if form.Type != KEY {
			return nil, fmt.Errorf("expected a key, but got %s of type %s", form.Literal, tokenNames[form.Type])
		}
		v, err := ReadVal(l)
		if err != nil {
			return nil, fmt.Errorf("failed to read key/val pair:\n%v", err)
		}
		h.(*HashMap).vals[&Key{form.Literal}] = v
	}

	return &h, nil
}

func ReadVal(l *Lexer) (*Any, error) {
	for form := l.NextLexeme(); form.Type != RBRACE; form = l.NextLexeme() {
		switch form.Type {
		case SYM, STR, KEY, INT, FLOAT:
			s, err := ReadAtomLiteral(form)
			if err != nil {
				return nil, fmt.Errorf("failed to read atom:\n%v", err)
			}
			return s, nil
		case LPAREN:
			s, err := ReadListLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read list:\n%v", err)
			}
			return s, nil
		case LBRACE:
			s, err := ReadHashLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read hashmap:\n%v", err)
			}
			return s, nil
		case LBRACK:
			s, err := ReadVecLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read vector:\n%v", err)
			}
			return s, nil
		default:
			return nil, fmt.Errorf("could not read \"%s\" of type %s", form.Literal, tokenNames[form.Type])
		}
	}
	return nil, fmt.Errorf("undexpected EOF when reading value associated with key")
}
