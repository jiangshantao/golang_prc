package main

import (
	"fmt"
	_ "fmt"
)

func pointtype(){
	var i1 int = 5
	var intP *int //定义一个int值的指针变量, var variable_name *type
	intP = &i1 //指针变量赋值, 指针变量的储存地址

	fmt.Printf("A integer: %d, its location in memory: %p\n", i1, &i1)
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	//使用指针变量去访问值, 用*pointer_variable进行表示
}