package controller

import (
	"context"
	"prac_proj/model"
	pb "prac_proj/proto"
)



type RPCServer struct {
	
	pb.UnimplementedIvServiceServer
}

// var (
// 	ItemService it.IUpdateInventory
// 	InventoryService interfaces.Inventory
// 	Mcoll            *mongo.Collection
// )

func (s *RPCServer) create_IV(ctx context.Context,val *pb.IVStruct)(*pb.Response_IV,error){
	var input *model.IV_model

	if err:= ctx.sho



}