package controller

import (
	"context"
	"prac_proj/interfaces"
	"prac_proj/model"
	pb "prac_proj/proto"
)



type RPCServer struct {
	
	pb.UnimplementedIvServiceServer
}
var (IV_service interfaces.IV_int)

func (s *RPCServer) Create_IV(ctx context.Context,val *pb.IVStruct)(*pb.Response_IV,error){
input :=&model.IV_model{
	Name:     val.Name,
	Quantity: val.Quantity,
}

_,err:=IV_service.Create__IV(input)
if err!= nil{
	return nil,err
}else{
	iv:=&pb.Response_IV{
		Ctd: "success",
	}
	return iv,nil
}



	




}