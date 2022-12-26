package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Hello string

type Person struct {
	firstName string
	lastName  string
}

func GetHello(c *gin.Context) {
	log.Println("Said Hello")
	c.JSON(http.StatusOK, gin.H{"message": Hello})
}
