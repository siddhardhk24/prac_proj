package controller

import (
	"net/http"
	grpcclient "prac_proj/client/grpcClient"
	pb "prac_proj/proto"

	"github.com/gin-gonic/gin"
)

func Handle_IV_CreateCust(c *gin.Context) {
	grpcClient, _ := grpcclient.Get_IV_Service()
	var res pb.IVStruct

	if err := c.ShouldBindJSON(&res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	resp, err := grpcClient.Create_IV(c.Request.Context(), &res)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": resp})

}
