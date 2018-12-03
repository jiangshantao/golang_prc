package main

import "fmt"

type user struct {
	name string
	email string
}

func (u user)notify(num int) int{
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
	a := num + 1
	return a
}

func (u *user)changeEmail(email string){
	fmt.Print(u.email + "\n")
	u.email = email
	fmt.Println("user Value", u)
}

func main(){

	frank := user{
		name: "jiangshantao",
		email: "1247803229@qq.com",
	}

	lisa := &user{"Lisa", "lisa@bytedance.com"}
	//lisa.notify(3) //Go在背后执行了(*lisa).notify()的操作
	//
	//frank.notify(2)

	lisa.changeEmail("hehe")
	frank.changeEmail("jia")
}