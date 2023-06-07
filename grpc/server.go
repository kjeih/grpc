package grpc

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

// should implement the interface myPkgName.InvoicerServer
type Server struct {
	// type embedded to comply with Google lib
	UnimplementedGRPCServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterGRPCServer(grpcServer, s)
	if e := grpcServer.Serve(listener); e != nil {
		return errors.New(fmt.Sprintf("failed to serve: %v", e))
	}

	return errors.New("server downed")
}

func (Server) SubscribeMessage(request *MessageRequest, stream GRPC_SubscribeMessageServer) error {
	for {
		err := stream.Send(&MessageReply{
			Message: "Hello Clients!!!",
		})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

func (Server) SendMessage(context context.Context, request *SendRequest) (*SendReply, error) {
	fmt.Printf("Send Request %s \n", request.Name)
	return &SendReply{Message: "Completed"}, nil
}
