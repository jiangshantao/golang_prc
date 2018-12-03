package main

import "fmt"

type user struct {
	name string
	email string
	ext int
	privileged bool
}

var bill user

type QA struct{
	name string
	age int
	language string
	isMaster bool
}

type admin struct {
	person user
	level int
}


func statementStruct(){
	jst := user{
		name: "jiangshantao",
		email: "jiangshantao@bytedance.com",
		ext: 3,
		privileged: true,
	}
	fmt.Println(jst)

	jst1 := QA{"jiangshantao", 24, "python", true}
	fmt.Println(jst1)

	frank := admin{
		person: user{
			name: "frank",
			email: "1247803229@qq.com",
			ext: 2,
			privileged: false,
		},
		level: 2,
	}
	fmt.Println(frank)
}


func main(){
	fmt.Println(bill)
	statementStruct()
}