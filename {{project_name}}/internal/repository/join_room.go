package repository

import (
	"context"
	json2 "encoding/json"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"log"
)

func (c chatRepository) JoinRoom(
	request *pb.JoinRoomRequest,
	server pb.ChatService_JoinRoomServer,
) error {
	messages := c.cache.GetMessages(request.GetRoomId())
	log.Println(messages)
	for _, msg := range messages {
		if err := server.Send(msg); err != nil {
			log.Printf("send error %v", err)
		}
	}

	ctx := context.Background()

	sub := c.cache.Subscribe(ctx, request.GetRoomId())
	defer sub.Close()

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		message := pb.Message{}
		if err := json2.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Println(err)
			panic(err)
		}
		err = server.Send(&message)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
