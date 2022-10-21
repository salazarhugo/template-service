package app

import (
	"context"
	"github.com/go-redis/redis/v9"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/services/chat-service/cache"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/repository"
	"google.golang.org/api/chat/v1"
	"sync"
	"time"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	mu             sync.Mutex
	chatRepository repository.ChatRepository
	channel        map[string][]chan *chat.Message
}

func NewServer() *Server {
	cache := cache.NewCache(
		time.Duration(time.Duration.Hours(1)),
		redis.NewClient(&redis.Options{
			Addr:     "redis-18624.c228.us-central1-1.gce.cloud.redislabs.com:18624",
			Password: "mBiW18GNIgPzQTbBMDEz71UVsAcNDOYF",
			DB:       0,
		}),
	)

	chatRepository := repository.NewChatRepository(cache)

	return &Server{
		UnimplementedChatServiceServer: pb.UnimplementedChatServiceServer{},
		mu:                             sync.Mutex{},
		chatRepository:                 chatRepository,
		channel:                        make(map[string][]chan *chat.Message),
	}
}

func (s *Server) TypingChannel(server pb.ChatService_TypingChannelServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetRoomId(ctx context.Context, req *pb.GetRoomIdReq) (*pb.RoomId, error) {
	//TODO implement me
	return &pb.RoomId{}, nil
}

func (s *Server) AddToken(ctx context.Context, req *pb.AddTokenReq) (*pb.Empty, error) {
	//TODO implement me
	return &pb.Empty{}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.UserIdReq) (*pb.Empty, error) {
	//TODO implement me
	return &pb.Empty{}, nil
}
