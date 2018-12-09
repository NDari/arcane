package main

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/NDari/arcane/reader"
// 	"github.com/NDari/arcane/types"
// )

// TODO make the reader take a lexer so that it can be reused over and over.

// func Read(s string) ([]types.Any, error) {
// 	var forms []types.Any
// 	l := reader.NewLexer()
// 	l.SetInput(s)
// 	for form := l.NextLexeme(); form.Type != reader.EOF; form = l.NextLexeme() {
// 		switch form.Type {
// 		case reader.IDENT, reader.STR, reader.KEY, reader.I64, reader.F64:
// 			s, err := ReadAtomLiteral(form)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read atom:\n%v", err)
// 			}
// 			forms = append(forms, s)
// 		case reader.LBRACK:
// 			s, err := ReadListLiteral(l)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read list:\n%v", err)
// 			}
// 			forms = append(forms, s)
// 		case reader.LBRACE:
// 			s, err := ReadHashLiteral(l)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read hash:\n%v", err)
// 			}
// 			forms = append(forms, s)
// 		default:
// 			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
// 		}
// 	}
// 	return forms, nil
// }

// func ReadAtomLiteral(form *reader.Lexeme) (types.Any, error) {
// 	switch form.Type {
// 	case reader.IDENT, reader.KEY:
// 		return types.Key{Val: form.Literal}, nil
// 	case reader.STR:
// 		return types.Str{Val: form.Literal}, nil
// 	case reader.I64:
// 		v, err := strconv.Atoi(form.Literal)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to parse I64 from %s:\n%v", form.Literal, err)
// 		}
// 		return types.I64{Val: int64(v)}, nil
// 	case reader.F64:
// 		v, err := strconv.ParseFloat(form.Literal, 64)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to parse F64 from %s:\n%v", form.Literal, err)
// 		}
// 		return types.F64{Val: v}, nil
// 	}
// 	return nil, fmt.Errorf("unknown atom: %s", form.Literal)
// }

// func ReadListLiteral(l *reader.Lexer) (*types.List, error) {
// 	lst := types.NewList()
// 	for form := l.NextLexeme(); form.Type != reader.RBRACK; form = l.NextLexeme() {
// 		fmt.Println("next is", form)
// 		switch form.Type {
// 		case reader.IDENT, reader.STR, reader.KEY, reader.I64, reader.F64:
// 			s, err := ReadAtomLiteral(form)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read atom:\n%v", err)
// 			}
// 			lst.Append(s)
// 		case reader.LBRACK:
// 			s, err := ReadListLiteral(l)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read list:\n%v", err)
// 			}
// 			lst.Append(s)
// 		case reader.LBRACE:
// 			s, err := ReadHashLiteral(l)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read hash:\n%v", err)
// 			}
// 			lst.Append(s)
// 		default:
// 			return nil, fmt.Errorf("could not read \"%s\"", form.Literal)
// 		}
// 	}
// 	fmt.Println("returning", lst)
// 	return lst, nil
// }

// func ReadHashLiteral(l *reader.Lexer) (*types.Map, error) {
// 	return nil, nil
// }

// func ReadKvPair(l *reader.Lexer) (*types.List, error) {
// 	maybeKey := l.NextLexeme()
// 	if maybeKey.Type == reader.RBRACE {
// 		return nil, nil
// 	}
// 	if maybeKey.Type != reader.KEY {
// 		return nil, fmt.Errorf("When parsing hash literal, expected Key, found %s: %s", maybeKey.String(), maybeKey.Literal)
// 	}
// 	maybeVal := l.NextLexeme()
// }
