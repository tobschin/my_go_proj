package main

import "testing"

func Test_hello_should_be_empty(t *testing.T) {

	if SayHello() != "" {
		t.Fatalf("hello should be empty")
	}

}

func Test_hello_should_not_be_empty(t *testing.T) {

	Hello = "Hello World"

	if SayHello() != "Hello World" {
		t.FailNow()
	}

}
