package main

import (
	"fmt"
)

func sum(a []int, o int ,c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Printf("order: %d, sum: %d\n",o,sum)
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	fmt.Println(a[:len(a)/2])

	c := make(chan int)
	go sum(a[:len(a)/2], 1, c)
	go sum(a[len(a)/2:], 2, c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
