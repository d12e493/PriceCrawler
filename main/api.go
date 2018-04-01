package main

import (
	"product-query/controller"
	. "product-query/global"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1/product")
	{
		v1.GET("/", controller.FindProduct)
		v1.GET("/:id", controller.FindProduct)
		v1.POST("/", controller.CreateProduct)
		v1.PUT("/", controller.UpdateProduct)
		v1.DELETE("/:id", controller.DeleteProduct)
	}
	Logger.Info("API server is start")
	router.Run(":" + strconv.Itoa(Config.Api.Port))
}
