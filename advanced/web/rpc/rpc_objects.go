// rpc_objects.go
package rpc_objects

import "fmt"

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) error {
	fmt.Println("invoke Multiply", *args)
	*reply = args.N * args.M
	return nil
}
