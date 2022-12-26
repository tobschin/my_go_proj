package main

import (
	"testing"

	"github.com/tobischin/my_go_proj/handler"
)

func Test_hello_should_be_empty(t *testing.T) {

	if SayHello() != "" {
		t.Fatalf("hello should be empty")
	}

}

func Test_hello_should_not_be_empty(t *testing.T) {

	handler.Hello = "Hello World"

	if SayHello() != "Hello World" {
		t.FailNow()
	}

}
