package main

import (
	"fmt"
	//"unsafe"
)

const (
	_ = iota //iota=0
	KB int64 = 1 << (10 * iota)
	MB
	GB
)

const (
	a = "hello world"
	b = len(a)
)

func main(){
	fmt.Println(a, b, KB, MB, GB)
}