package main

import (
	"fmt"
	"strings"
	"time"
)

func shout(ping chan string, pong chan string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	time.Sleep(10 * time.Second)
	fmt.Println("Type something and press ENTER (enter Q to quit")
	for {
		fmt.Print("-> ")

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}
		ping <- userInput
		// wait for a response
		response := <-pong
		fmt.Println("Response: ", response)
	}

	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}
