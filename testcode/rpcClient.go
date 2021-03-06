package main

import (
	"net/rpc"
	"fmt"
	"log"
)
type Args struct {
	A, B int
}
func main()  {
	client, err := rpc.DialHTTP("tcp", "localhost" + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := &Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d \n", args.A, args.B, reply)

}