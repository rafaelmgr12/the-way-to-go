package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	fmt.Println("sending", 10)
	go block(c)
	c <- 10 // putting 10 on the channel
	fmt.Println("sent", 10)

}

func block(c chan int) {
	time.Sleep(5 * 1e9)
	fmt.Println("received", <-c) // value recieved from channel
}
