package main

import (
	"fmt"
	"strconv"
)

func Eval(ns *Namespace, args ...*Any) (*Any, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("invalid argument vector sent to eval: %v\n", args)
	}

	switch head := (*args[0]).(type) {
	case *Sym:
		v, found := ns.vals[head]
		if !found {
			return nil, fmt.Errorf("the symbol %s is not defined\n", head.val)
		}
		return v, nil
	case *Key:
		panic("key lookup in eval NYI")
	case *Str, *F64, *I64, *Vec, *HashMap:
		return &head, nil
	case *List:
		panic("list eval NYI")
	default:
		panic("strange case in eval")
	}
}

func Read(s string) ([]*Any, error) {
	var forms []*Any
	l := NewLexer()
	l.SetInput(s)
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
			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
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
			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
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
			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
		}
	}
	return &v, nil
}

func ReadHashLiteral(l *Lexer) (*Any, error) {
	return nil, nil
}

func Cons(ns *Namespace, args ...*Any) (*Any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("cons needs 2 arguments, but got %d", len(args))
	}

	lst, ok := (*args[1]).(*List)
	if !ok {
		return nil, fmt.Errorf("second argument to cons must be a list, but got %T", args[1])
	}
	var l Any = &List{
		args[0],
		lst,
	}
	return &l, nil
}
