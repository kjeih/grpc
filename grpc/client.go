package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// should implement the interface myPkgName.InvoicerServer
type Client struct {
	// type embedded to comply with Google lib
	GrpcClient GRPCClient
}

func NewClient(host string, port int) (*Client, error) {
	hostUrl := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(hostUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	grpcClient := NewGRPCClient(conn)
	return &Client{GrpcClient: grpcClient}, nil
}

func (cli *Client) SubscribeMessage(handler func(string)) error {

	stream, err := cli.GrpcClient.SubscribeMessage(context.Background(), &MessageRequest{Type: "1"})
	if err != nil {
		return err
	}

	for {
		message, e := stream.Recv()
		if e != nil {
			fmt.Println("client error", err.Error())
			continue
		}
		handler(message.Message)
	}
}

func (cli *Client) SendMessage(name string) error {
	reply, err := cli.GrpcClient.SendMessage(context.Background(), &SendRequest{Name: name})
	if err != nil {
		return err
	}
	fmt.Printf("Received %s \n", reply.Message)
	return nil
}
