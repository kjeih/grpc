package main

import (
	"fmt"
	"grpc/grpc"
)

func main() {
	println("gRPC client start")
	client, e := grpc.NewClient("localhost", 9000)
	if e != nil {
		fmt.Println("Error ", e.Error())
		return
	}

	messageHandler := func(message string) {
		fmt.Printf("Message - %s\n", message)
		client.SendMessage("Hi server")
	}
	err := client.SubscribeMessage(messageHandler)
	if err != nil {
		fmt.Println("Error ", err.Error())
		return
	}
}
