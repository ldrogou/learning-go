package main

import (
	"fmt"
	"time"
)

func ping(c chan<- string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("ping %v", i)
	}
}

func print(c <-chan string) {
	for {
		lec := <-c
		fmt.Println(lec)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	c := make(chan string)
	go ping(c)
	print(c)
}
