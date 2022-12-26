package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tobischin/my_go_proj/handler"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", handler.GetHello)
	r.POST("/", handler.PostHello)
	return r
}

func SetUpEnv() {
	handler.Hello = "Habe die Ehre"
}

func Test_GetHandler(t *testing.T) {

	SetUpEnv()
	r := SetUpRouter()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	responseStr := string(responseData)
	assert.Equal(t, `{"message":"Habe die Ehre"}`, responseStr)
}

func Test_PostHandler(t *testing.T) {
	postBody := `{"firstname" : "Hans", "lastname" : "Peter"}`
	r := SetUpRouter()
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(postBody)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	responseStr := string(responseData)
	assert.Equal(t, `{"firstname":"Hans","lastname":"Peter"}`, responseStr)
}
