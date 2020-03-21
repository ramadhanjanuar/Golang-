package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Printf("receive data %d \n", data)
		case <-time.After(time.Second * 8):
			fmt.Println("Time out no activiteis")
			break loop
		}

	}
}

func main() {
	var messages = make(chan int)
	go sendData(messages)
	retrieveData(messages)
}
