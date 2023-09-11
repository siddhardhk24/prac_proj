package routes

import (
	"prac_proj/client/controller"

	"github.com/gin-gonic/gin"
)

func IV_routes(r *gin.Engine){
 r.POST("/IV_customer",controller.Handle_IV_CreateCust)
}