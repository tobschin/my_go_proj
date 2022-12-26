package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tobschin/my_go_proj/handler"
)

func main() {
	handler.Hello = "Hello World"
	r := gin.Default()
	r.GET("/", handler.GetHello)
	r.POST("/", handler.PostHello)
	r.PUT("/:language", handler.PutHello)
	r.Run()
	log.Printf("Running on default http://localhost:8080")
}

func SayHello() string {
	log.Println(handler.Hello)
	return handler.Hello
}
