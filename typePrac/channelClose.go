package main

import "fmt"

func fibonacci(n int, c chan int){
	x, y := 0, 1
	for i := 0; i < n; i++{
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main(){
	c := make(chan int, 10)
	fmt.Printf("Channel Capacity: %d\n", cap(c))

	go fibonacci(cap(c), c)

	x := <- c
	fmt.Println("Channel C",x)

	for i := range c {
		fmt.Println(i)
	}
}