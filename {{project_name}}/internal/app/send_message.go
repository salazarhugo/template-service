package app

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) SendMessage(server pb.ChatService_SendMessageServer) error {
	_, err := utils.GetUserId(server.Context())
	if err != nil {
		return status.Error(codes.Internal, "failed to retrieve userID")
	}

	msg, err := server.Recv()
	if err != nil {
		return err
	}

	err = CheckPermissions()
	if err != nil {
		return err
	}

	err = s.chatRepository.SendMessage(msg, server)

	//go func() {
	//	streams := s.channel[msg.Room.GetId()]
	//	for _, msgChan := range streams {
	//		msgChan <- msg
	//	}
	//}()

	return nil
}

func CheckPermissions() error {
	//isMember := roomCache.IsMember(userId, msg.Room.Id)
	isMember := true
	if isMember == false {
		log.Println("Can't access rooms you are not member of")
		return status.Error(codes.PermissionDenied, "you are not member of this group chat")
	}
	return nil
}
