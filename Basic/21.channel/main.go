package main

import (
	"fmt"
	"runtime"
)

func printMessages(message chan string) {
	fmt.Println(<-message)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	for _, each := range []string{"ramdhan", "puji", "januar"} {
		go func(name string) {
			var text = fmt.Sprintf("Hayyy %s", name)
			messages <- text
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessages(messages)
	}
}
