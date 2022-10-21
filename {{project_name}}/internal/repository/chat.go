package repository

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	"github.com/salazarhugo/cheers1/services/chat-service/cache"
)

type ChatRepository interface {
	CreateRoom(name string, members []string) (*pb.Room, error)
	JoinRoom(request *pb.JoinRoomRequest, server pb.ChatService_JoinRoomServer) error
	ListRoom(userID string) ([]*pb.Room, error)

	ListMembers(request *pb.ListMembersRequest) ([]*user.UserItem, error)

	SendMessage(msg *pb.Message, server pb.ChatService_SendMessageServer) error
}

type chatRepository struct {
	cache cache.RoomCache
}

func NewChatRepository(cache cache.RoomCache) ChatRepository {
	return &chatRepository{
		cache: cache,
	}
}
