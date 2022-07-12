package main

import (
	"learning_go/idiomatic_go/pkg/display"
	"learning_go/idiomatic_go/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("an exciting message")
}
