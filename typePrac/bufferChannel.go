package main

import "fmt"

func main(){
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	x, y := <-c, <-c
	z, ok := <-c
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z, ok)
}