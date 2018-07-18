package arcane

import "fmt"

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

func Cons(ns *Namespace, args ...*Any) (*Any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("cons needs 2 arguments, but got %d", len(args))
	}

	lst, ok := (*args[1]).(*List)
	if !ok {
		return nil, fmt.Errorf("second argument to cons must be a list, but got %T", args.vals[1])
	}
	lst = &List{
		args.vals[0],
		lst,
	}
	return lst, nil
}
