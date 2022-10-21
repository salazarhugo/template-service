package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
)

func (s *Server) TypingEnd(ctx context.Context, req *pb.TypingReq) (*pb.Empty, error) {
	//TODO implement me
	return &pb.Empty{}, nil
}
