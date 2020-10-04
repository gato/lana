package main

import (
	"github.com/gato/lana/checkout"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	apiv1 := r.Group("/api/v1/")
	checkout.AddRoutes(apiv1)
	r.Run() // listen and serve on 0.0.0.0:8080
}
