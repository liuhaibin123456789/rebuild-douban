package tool

import (
	"douban/proto"
	"google.golang.org/grpc"
	"log"
)

// GrpcClient grpc客户端
var GrpcClient proto.UserCenterClient

func InitGrpcClient() error {
	conn, err := grpc.Dial(":8085", grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return err
	}

	GrpcClient = proto.NewUserCenterClient(conn)
	return nil
}
