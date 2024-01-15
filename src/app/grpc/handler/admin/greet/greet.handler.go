package greet

import (
	"context"
	pb "gtimekeeper/src/app/grpc/proto"

	"google.golang.org/grpc"
)

type HandlerGreet struct {
	pb.GreetingServiceServer
}

func (h *HandlerGreet) Boot(server *grpc.Server) {
	pb.RegisterGreetingServiceServer(server, h)
}

func (h *HandlerGreet) Greet(ctx context.Context, req *pb.GreetingRequest) (res *pb.GreetingResponse, err error) {
	return &pb.GreetingResponse{
		Message: "Hallo " + req.Greeting.GetName(),
	}, nil
}
