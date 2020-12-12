package main

import "fmt"

func main() {
	c := make(chan int)
	s := []int{1, 25, 10, 3}
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)
	go sum([]int{35, 2}, c)
	go sum([]int{55, 2}, c)
	x, y, z, a := <-c, <-c, <-c, <-c
	fmt.Println(x, y, z, a)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
