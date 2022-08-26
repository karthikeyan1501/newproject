package main

import (
	"github.com/gin-gonic/gin"
	"newproject/services"
)
func main() {
	router := gin.Default() 
	g := router.Group("")
	{
		g.GET("/state", services.GetDetail)
		g.GET("/states", services.GetDetails)
		g.POST("/states", services.PostDetails)
		g.PUT("/states", services.PutDetails)
		g.DELETE("/states", services.DeleteDetails)
	
	}

	
	router.Run(":8080")
}
