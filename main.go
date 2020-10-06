package main

import (
	"flag"
	"fmt"
	"github.com/gato/lana/checkout"
	"github.com/gin-gonic/gin"
)

var (
	port = flag.Int64("port", 8080, "port to listen to")
)

func main() {
	flag.Parse()
	r := gin.Default()
	apiv1 := r.Group("/api/v1/")
	checkout.AddRoutes(apiv1)
	runPort := fmt.Sprintf(":%d", *port)
	r.Run(runPort)
}
