package main

import (
	"fmt"
)

func Cons(ns *Namespace, args *Vec) (*Any, error) {
	if len(args.vals) != 2 {
		return nil, fmt.Errorf("cons needs 2 arguments, but got %d", len(args.vals))
	}

	lst, ok := (*args.vals[0]).(*List)
	if !ok {
		return nil, fmt.Errorf("expected first argument to cons to be a list")
	}
	var l Any = &List{
		args.vals[1],
		lst,
	}
	return &l, nil
}

func Eval(ns *Namespace, arg *Any) (*Any, error) {
	switch arg0 := (*arg).(type) {
	case *Sym:
		v, found := ns.vals[*arg0]
		if !found {
			return nil, fmt.Errorf("the symbol %s is not defined\n", arg0.val)
		}
		return v, nil
	case *Str, *F64, *I64, *Vec, *HashMap:
		return &arg0, nil
	case *List:
		switch head := (*arg0.head).(type) {
		case *Sym:
			qoute := Sym{"quote"}
			if *head == qoute {
				var t Any = arg0.tail
				return &t, nil
			}
			cond := Sym{"cond"}
			if *head == cond {
				v, err := evalCond(ns, arg0.tail)
				if err != nil {
					return nil, fmt.Errorf("failed to eval cond statement:\n%v", err)
				}
				return v, nil
			}
			res, err := evList(arg0.tail)
			Apply(ns, head, res)
		}
	default:
		panic("strange case in eval")
	}
}

func evalCond(ns *Namespace, condList *List) (*Any, error) {
	return nil, fmt.Errorf("evalcond is NYI")
}

func Apply(ns *Namespace, args *Any) (*Any, error) {
	return nil, fmt.Errorf("Apply NYI")
}
