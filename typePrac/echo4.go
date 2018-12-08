package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s","", "separator")

func main(){
	flag.Parse()
	av := flag.Args()
	fmt.Println("All params:",av)
	fmt.Println("n var value:", n)
	fmt.Println("n value:", *n)
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n{
		fmt.Println("Error!")
	}
}

/*
./echo4 -n=true -s / a c v,执行结果
All params: [a c v]
n var value: 0xc00007a002
n value: true
a/c/v
 */