package repository

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
)

func (c chatRepository) ListRoom(userID string) ([]*pb.Room, error) {
	rooms := c.cache.ListRoom(userID)

	//var ctx = context.Background()
	//
	//sub := s.chatRepository.Subscribe(ctx, "messages")
	//defer sub.Close()
	//for {
	//	msg, err := sub.ReceiveMessage(ctx)
	//	if err != nil {
	//		panic(err)
	//	}
	//	message := chatpb.Message{}
	//	if err := json2.Unmarshal([]byte(msg.Payload), &message); err != nil {
	//		panic(err)
	//	}
	//
	//	room := roomCache.GetRoomWithId(userId, message.GetRoom().GetId())
	//
	//	if err := roomStream.Send(room); err != nil {
	//		log.Printf("send error %v", err)
	//	}
	//}

	return rooms, nil
}
