package main

import (
	"fmt"
	"grpc/grpc"
)

func main() {
	println("gRPC server start")

	s := grpc.NewServer()

	e := s.Run(9000)
	if e != nil {
		fmt.Println(e)
	}
}
