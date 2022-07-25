package service

import (
	"context"

	pb "github.com/wenealves10/microservices-golang-grpc-k8s/pkg/pb/greeting/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterService struct {
}

func (gs *GreeterService) Greeting(ctx context.Context, req *pb.GreetRequest) (res *pb.GreetResponse, err error) {
	if req.Msg == nil {
		err = status.New(codes.InvalidArgument, "Message cannot be empty").Err()
		return
	}

	helloMsg := req.Msg.Greeter.String() + " " + req.Msg.Name

	res = &pb.GreetResponse{
		Resp: helloMsg,
	}

	return

}
