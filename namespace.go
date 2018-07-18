package arcane

import "fmt"

type Namespace struct {
	upper *Namespace

	HashMap
}

func NewNS(upper *Namespace) *Namespace {
	m := HashMap{
		make(map[Sym]Any),
	}
	return &Namespace{
		upper,
		m,
	}
}

func eval(args *List, ns *Namespace) (Any, error) {
	h, _ := args.Car(), args.Cdr()
	if h == nil {
		return nil, fmt.Errorf("Empty list passed to eval\n")
	}
	switch val := h.(type) {
	case Sym:
		v, found := ns.vals[val]
		if !found {
			return nil, fmt.Errorf("%v is not defined\n", val)
		}
		return v, nil
	case Str, F64, I64:
		return val, nil
	case Key:
		panic("key lookup not yet implemented")
	case *List:
		panic("got a list in eval")
	default:
		panic("strange case in eval")
	}
}

func TopLevel() *Namespace {
	m := make(map[Sym]Any)
	topNameSpace := Sym{
		"arcane",
	}

	eval := Sym{
		"eval",
	}

	atomHuh := Sym{
		"atom?",
	}

	isAtom := Fn{
		func(args *List, ns *Namespace) (Any, error) {
			h := args.Car()
			if h == nil {
				return nil, nil
			}
			if _, ok := h.(Atom); ok {
				return 1, nil
			}
			return nil, nil
		},
	}

	m[atomHuh] = isAtom

	// lambda := Sym{
	// 	"fn",
	// }

	// fun := Fn{
	// 	func
	// }

	def := Sym{
		"def",
	}

	define := Fn{
		func(args *List, ns *Namespace) (Any, error) {
			// h, t := args.Car(), args.Cdr()
			// if lst, ok := h.(*List); ok {
			// switch v := lst.Car() {

			// }
			// }
			return nil, nil
		},
	}

	m[def] = define

	return &Namespace{
		nil,
		m,
	}
}
