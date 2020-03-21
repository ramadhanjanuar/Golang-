package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Application running")
	time.Sleep(time.Second * 4)
	fmt.Println("Continue..")

	var timer = time.NewTimer(time.Second * 4)
	<-timer.C
	fmt.Println("continue..")

	var ch = make(chan bool)

	time.AfterFunc(time.Second*2, func() {
		fmt.Println("expired")
		ch <- true
	})

	<-ch

	fmt.Println("continue")

	<-time.After(time.Second * 2)

	fmt.Println("end")
	/* Print Hello every 1 second
	for true {
		fmt.Println("Hello..")
		time.Sleep(time.Second * 1)
	}
	*/

	done := make(chan bool)
	ticker := time.NewTicker(time.Second * 2)

	go func() {
		time.Sleep(time.Second * 3)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hello!!", t)
		}
	}

	var timeout = 5
	var chNew = make(chan bool)

	go timerMethod(timeout, chNew)
	go watcher(timeout, chNew)

	var input string
	fmt.Print("what is 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}
}

func timerMethod(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}
