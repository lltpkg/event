package main

import (
	"fmt"

	"github.com/lltpkg/event"
)

func main() {
	evName := "Greeting"
	go func() {
		data := "World"
		event.FireEvent(evName, data)
	}()
	evChan, unSub := event.EventChannel(evName)
	defer unSub()
	receivedData := <-evChan
	fmt.Println("Hello,", receivedData)
}
