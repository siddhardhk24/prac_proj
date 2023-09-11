package main

import (
	"context"
	"fmt"
	"net"
	"prac_proj/config"
	"prac_proj/controller"
	"prac_proj/service"

	"prac_proj/constants"

	"go.mongodb.org/mongo-driver/mongo"

	pro "prac_proj/proto"

	"google.golang.org/grpc"
)

func initDB(client *mongo.Client) {
	IV_Collection := config.GetCollection(client, "IV_DB", "Cust_Det")
	controller.IV_service = service.InitInventory(IV_Collection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	if err != nil {
		panic(err)
	}
	fmt.Println("1")
	initDB(mongoclient)
	defer mongoclient.Disconnect(context.TODO())

	lis, err := net.Listen("tcp", constants.Port)
	fmt.Println("2")
	if err != nil {
		fmt.Println("failed to listen")
	}

	s := grpc.NewServer()
	fmt.Println("3")
	pro.RegisterIvServiceServer(s, &controller.RPCServer{})
	fmt.Println("server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve")
	}

}
