package grpcserv

import (
	"context"

	bookcenterv1 "github.com/buffalo-big-tech-system/grpc-protos/protos/gen/go/proto/bookcenter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BookcenterServer struct {
	bookcenterv1.UnimplementedBookcenterServer
}

func RegisterGRPC(gRPC *grpc.Server) {
	bookcenterv1.RegisterBookcenterServer(gRPC, &BookcenterServer{})
}

func (BookcenterServer) GetTotalPages(context.Context, *emptypb.Empty) (*GetTotalPagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalPages not implemented")
}
func (UnimplementedBookcenterServer) GetBooksByPage(context.Context, *GetBooksByPageRequest) (*GetBooksByPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooksByPage not implemented")
}
