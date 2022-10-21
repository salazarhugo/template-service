package repository

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	userpb "github.com/salazarhugo/cheers1/genproto/cheers/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c chatRepository) ListMembers(
	request *pb.ListMembersRequest,
) ([]*user.UserItem, error) {

	membersIDs := c.cache.GetRoomMembers(request.RoomId)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("android-gateway-clzdlli7.nw.gateway.dev:443", opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)
	response, err := client.GetUserItemsIn(context.Background(), &userpb.GetUserItemsInRequest{UserIds: membersIDs})
	if err != nil {
		return nil, err
	}

	return response.Users, nil
}
