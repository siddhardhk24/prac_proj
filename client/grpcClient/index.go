package grpcClient

import (
	"log"
	"prac_proj/constants"
	pb "prac_proj/proto"
	"sync"

	"google.golang.org/grpc"
)

var IV_Once sync.Once

type Grpc_IV_Client pb.IvServiceClient

var (
	IV_Instance Grpc_IV_Client
)

func Get_IV_Service() (Grpc_IV_Client, *grpc.ClientConn) {
	var conn *grpc.ClientConn
	IV_Once.Do(func() {
		conn, err := grpc.Dial("localhost"+constants.Port, grpc.WithInsecure())

		if err != nil {
			log.Fatalf("failed to connect: %v", err)
		}

		IV_Instance = pb.NewIvServiceClient(conn)
	})

	return IV_Instance, conn
}
