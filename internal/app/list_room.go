package app

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListRoom(
	request *pb.ListRoomRequest,
	server pb.ChatService_ListRoomServer,
) error {
	userID, err := utils.GetUserId(server.Context())
	if err != nil {
		return status.Error(codes.Internal, "failed retrieving userID")
	}

	rooms, err := s.chatRepository.ListRoom(userID)
	if err != nil {
		return err
	}

	for i := range rooms {
		if err := server.Send(rooms[i]); err != nil {
			log.Printf("error while sending to stream %v", err)
		}
	}

	return nil
}
