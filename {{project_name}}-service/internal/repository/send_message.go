package repository

import (
	json2 "encoding/json"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"log"
)

func (c chatRepository) SendMessage(
	msg *pb.Message,
	server pb.ChatService_SendMessageServer,
) error {
	c.cache.SetMessage(msg)

	json, err := json2.Marshal(msg)
	if err != nil {
		panic(err)
	}

	err = c.cache.Publish(server.Context(), msg.GetRoom().GetId(), json).Err()

	if err != nil {
		log.Println("could not publish to channel", err)
	}

	ack := pb.SendMessageResponse{Status: pb.Message_SENT}
	err = server.SendAndClose(&ack)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
