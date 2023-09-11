package main

import (
	grpcclient"prac_proj/client/grpcClient"
	customer"prac_proj/client/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	_,conn:=grpcclient.Get_IV_Service()
	defer conn.Close()
	r := gin.Default()
	customer.IV_routes(r)
	r.Run(":8000")
}