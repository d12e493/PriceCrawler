package main

import (
	"io"
	"log"
	"os"
	v1Controller "product-query/controller/v1"
	. "product-query/global"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// change logger to file
	gin.DisableConsoleColor()
	if logFile, err := CreateLogFile("access"); err == nil {
		log.SetOutput(logFile)
		gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	} else {
		Logger.Error("Only Log in standard out")
	}

	router := gin.Default()

	v1 := router.Group("/v1/product")
	{
		v1.GET("/", v1Controller.FindProduct)
		v1.GET("/:id", v1Controller.FindProduct)
		v1.POST("/", v1Controller.CreateProduct)
		v1.PUT("/", v1Controller.UpdateProduct)
		v1.PATCH("/", v1Controller.SyncProduct)
		v1.DELETE("/:id", v1Controller.DeleteProduct)
	}

	port := strconv.Itoa(Config.Api.Port)

	Logger.Info("API server is start")
	Logger.Info("Host : " + Config.Api.Host)
	Logger.Info("Port : " + port)
	Logger.Info("Start at " + time.Now().String())
	router.Run(":" + port)
}
