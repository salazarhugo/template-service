package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListMembers(
	ctx context.Context,
	request *pb.ListMembersRequest,
) (*pb.ListMembersResponse, error) {
	_, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	members, err := s.chatRepository.ListMembers(request)
	if err != nil {
		return nil, err
	}

	return &pb.ListMembersResponse{Users: members}, nil
}
