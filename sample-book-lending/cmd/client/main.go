package main

//todo
import (
	// ...
	"context"
	"google.golang.org/grpc/metadata"
	pb "sample-book-lending/pkg/grpc"
)

func main() {

}

func callGetUser(client pb.AccountServiceClient, userID, token string) (*pb.GetUserInfoResponse, error) {
	md := metadata.Pairs("authorization", "Bearer "+token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	req := &pb.GetUserInfoRequest{UserId: userID}
	return client.GetUserInfo(ctx, req)
}
