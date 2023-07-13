package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gustafabayu/go-restGin/controllers/productcontroller"
	"github.com/gustafabayu/go-restGin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.Run()
}
