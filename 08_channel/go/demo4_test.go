package _go

import (
	"fmt"
	"testing"
)

func Test4(t *testing.T) {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
