package repository

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
)

func (c chatRepository) CreateRoom(
	name string,
	members []string,
) (*pb.Room, error) {
	// Direct Chat
	if len(members) == 2 {
		room, err := c.cache.GetOrCreateDirectRoom(members[0], members[1])
		if err != nil {
			return nil, err
		}
		return room, nil
	}

	room := c.cache.CreateGroup(name, members)

	return room, nil
}
