package arcane

import "fmt"

func Cons(args Vec, ns *Namespace) (*List, error) {
	if len(args.vals) < 2 {
		return nil, fmt.Errorf("cons needs 2 arguments, but got %d", len(args.vals))
	}

	lst, ok := args.vals[1].(*List)
	if !ok {
		return nil, fmt.Errorf("second argument to cons must be a list, but got %T", args.vals[1])
	}
	return &List{
		args.vals[0],
		lst,
	}, nil
}
