package main

import (
	"fmt"
	"strconv"

	"github.com/NDari/arcane/reader"
	"github.com/NDari/arcane/types"
)

// func Eval(ns *Namespace, expr *types.List) (types.Any, error) {
// 	switch head := (*args[0]).(type) {
// 	case *Sym:
// 		v, found := ns.vals[head]
// 		if !found {
// 			return nil, fmt.Errorf("the symbol %s is not defined\n", head.val)
// 		}
// 		return v, nil
// 	case *Key:
// 		panic("key lookup in eval NYI")
// 	case *Str, *F64, *I64, *Vec, *HashMap:
// 		return &head, nil
// 	case *List:
// 		panic("list eval NYI")
// 	default:
// 		panic("strange case in eval")
// 	}
// }

func Read(s string) ([]types.Any, error) {
	var forms []types.Any
	l := reader.NewLexer()
	l.SetInput(s)
	for form := l.NextLexeme(); form.Type != reader.EOF; form = l.NextLexeme() {
		switch form.Type {
		case reader.IDENT, reader.STR, reader.KEY, reader.I64, reader.F64:
			s, err := ReadAtomLiteral(form)
			if err != nil {
				return nil, fmt.Errorf("failed to read atom:\n%v", err)
			}
			forms = append(forms, s)
		case reader.LBRACK:
			s, err := ReadListLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read list:\n%v", err)
			}
			forms = append(forms, s)
		default:
			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
		}
	}
	return forms, nil
}

func ReadAtomLiteral(form *reader.Lexeme) (types.Any, error) {
	switch form.Type {
	case reader.IDENT, reader.KEY:
		return types.Key{Val: form.Literal}, nil
	case reader.STR:
		return types.Str{Val: form.Literal}, nil
	case reader.I64:
		v, err := strconv.Atoi(form.Literal)
		if err != nil {
			return nil, fmt.Errorf("failed to parse I64 from %s:\n%v", form.Literal, err)
		}
		return types.I64{Val: int64(v)}, nil
	case reader.F64:
		v, err := strconv.ParseFloat(form.Literal, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse F64 from %s:\n%v", form.Literal, err)
		}
		return types.F64{Val: v}, nil
	}
	return nil, fmt.Errorf("unknown atom: %s", form.Literal)
}

func ReadListLiteral(l *reader.Lexer) (*types.List, error) {
	lst := types.NewList()
	for form := l.NextLexeme(); form.Type != reader.RBRACK; form = l.NextLexeme() {
		fmt.Println("next is", form)
		switch form.Type {
		case reader.IDENT, reader.STR, reader.KEY, reader.I64, reader.F64:
			s, err := ReadAtomLiteral(form)
			if err != nil {
				return nil, fmt.Errorf("failed to read atom:\n%v", err)
			}
			lst.Append(s)
		case reader.LBRACK:
			s, err := ReadListLiteral(l)
			if err != nil {
				return nil, fmt.Errorf("failed to read list:\n%v", err)
			}
			lst.Append(s)
		default:
			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
		}
	}
	fmt.Println("returning", lst)
	return lst, nil
}

// func ReadVecLiteral(l *reader.Lexer) (types.Any, error) {
// 	var v Any = &Vec{
// 		make([]*Any, 0),
// 	}
// 	for form := l.NextLexeme(); form.Type != RBRACK; form = l.NextLexeme() {
// 		switch form.Type {
// 		case SYM, STR, KEY, INT, FLOAT:
// 			s, err := ReadAtomLiteral(form)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read atom:\n%v", err)
// 			}
// 			v.(*Vec).vals = append(v.(*Vec).vals, s)
// 		case LPAREN:
// 			s, err := ReadListLiteral(l)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read list:\n%v", err)
// 			}
// 			v.(*Vec).vals = append(v.(*Vec).vals, s)
// 		case LBRACE:
// 			s, err := ReadHashLiteral(l)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read hashmap:\n%v", err)
// 			}
// 			v.(*Vec).vals = append(v.(*Vec).vals, s)
// 		case LBRACK:
// 			s, err := ReadVecLiteral(l)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read vector:\n%v", err)
// 			}
// 			v.(*Vec).vals = append(v.(*Vec).vals, s)
// 		default:
// 			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
// 		}
// 	}
// 	return &v, nil
// }

// func ReadHashLiteral(l *Lexer) (*Any, error) {
// 	return nil, nil
// }

// func Cons(ns *Namespace, args ...*Any) (*Any, error) {
// 	if len(args) < 2 {
// 		return nil, fmt.Errorf("cons needs 2 arguments, but got %d", len(args))
// 	}

// 	lst, ok := (*args[1]).(*List)
// 	if !ok {
// 		return nil, fmt.Errorf("second argument to cons must be a list, but got %T", args[1])
// 	}
// 	var l Any = &List{
// 		args[0],
// 		lst,
// 	}
// 	return &l, nil
// }
