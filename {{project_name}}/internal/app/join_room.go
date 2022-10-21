package app

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) JoinRoom(request *pb.JoinRoomRequest, server pb.ChatService_JoinRoomServer) error {
	_, err := utils.GetUserId(server.Context())
	if err != nil {
		return status.Error(codes.Internal, "failed to retrieve userID")
	}

	err = s.chatRepository.JoinRoom(request, server)
	if err != nil {
		return err
	}

	return nil
}
