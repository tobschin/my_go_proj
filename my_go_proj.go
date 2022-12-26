package main

import (
	"log"

	"github.com/my_go_proj/handler/hello_handler"
)

var Hello string

func main() {
	Hello = "Hello World"
	log.Writer(hello_handler.GetHello())

}

func SayHello() string {
	log.Println(Hello)
	return Hello
}
