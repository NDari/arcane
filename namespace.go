package arcane

type Namespace struct {
	upper *Namespace

	vals map[*Sym]*Any
}

func NewNS(upper *Namespace) *Namespace {
	m := make(map[*Sym]*Any)
	return &Namespace{
		upper,
		m,
	}
}

// func TopLevel() *Namespace {
// 	m := make(map[Sym]Any)
// 	topNameSpace := Sym{
// 		"arcane",
// 	}

// 	eval := Sym{
// 		"eval",
// 	}

// 	atomHuh := Sym{
// 		"atom?",
// 	}

// isAtom := Fn{
// 	func(args *List, ns *Namespace) (Any, error) {
// 		h := args.Car()
// 		if h == nil {
// 			return nil, nil
// 		}
// 		if _, ok := h.(Atom); ok {
// 			return 1, nil
// 		}
// 		return nil, nil
// 	},
// }

// m[atomHuh] = isAtom

// lambda := Sym{
// 	"fn",
// }

// fun := Fn{
// 	func
// }

// def := Sym{
// 	"def",
// }

// define := Fn{
// 	func(args *List, ns *Namespace) (Any, error) {
// 		// h, t := args.Car(), args.Cdr()
// 		// if lst, ok := h.(*List); ok {
// 		// switch v := lst.Car() {

// 		// }
// 		// }
// 		return nil, nil
// 	},
// }

// m[def] = define

// return &Namespace{
// 	nil,
// 	m,
// }
// }
