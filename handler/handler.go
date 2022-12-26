package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Hello string

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func GetHello(c *gin.Context) {
	log.Println("Said Hello")
	c.JSON(http.StatusOK, gin.H{"message": Hello})
}

func PostHello(c *gin.Context) {
	body := Person{}
	c.BindJSON(&body)
	jsonStr, _ := json.Marshal(body)
	log.Printf("sent %s as body", jsonStr)
	c.JSON(http.StatusOK, &body)
}

func PutHello(c *gin.Context) {
	language := c.Param("language")
	if language == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("bad language"))
	}
	body := Person{}
	c.BindJSON(&body)
	jsonStr, _ := json.Marshal(body)
	log.Printf("should say hello in %s to ", language, jsonStr)
	status := http.StatusNotFound
	var message string
	if language == "spanish" {
		message = "Hola "
		status = http.StatusOK
	} else if language == "english" {
		message = "Hi "
		status = http.StatusOK
	} else if language == "german" {
		message = "Servus "
		status = http.StatusOK
	}
	if status == http.StatusOK {
		c.JSON(status, gin.H{"message": message + body.FirstName})
	} else {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("language not supported"))
	}
}
