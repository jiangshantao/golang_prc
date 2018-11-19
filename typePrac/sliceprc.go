package main

import "fmt"

var Array1 [5]int

var (
	x, y int
	a, s =100, "abc"
)

func firstLoop(){
	var c int
	for i := 0; i<10; i++{
		c++
	}
	println(c)
}

func makeslcie() []int{
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func sliceMakeSlice() []int{
	slice1 := []int{10, 20, 30, 40, 50}
	newSlice := slice1[1:3]
	return newSlice
}

func makeMap() map[string]int{
	m := make(map[string]int)
	m["a"] = 1
	return m
}

func main(){
	fmt.Println(Array1)
	fmt.Println(&a)
	fmt.Println(x)

	a := 100
	fmt.Println(&a)

	firstLoop()
	fmt.Println(makeslcie())
	fmt.Println(makeMap())
	p1 := new(map[string]int)
	fmt.Println(p1)
	fmt.Printf("make slice from slice: %v", sliceMakeSlice())
}