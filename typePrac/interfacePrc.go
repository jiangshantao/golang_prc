package main

import (
	"fmt"
	"strings"
)

type notifier interface {
	notify()
	useEmail()

}

type user2 struct {
	name string
	email string
}

type admin struct {
	name string
	email string
}

func (u user2)useEmail(){
	fmt.Printf("Used email is %s\n", strings.Split(u.email,"@")[1])
}

func (a admin)useEmail(){
	fmt.Printf("Admin email name is %s\n", a.email)
}

func (u *user2)notify(){
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func (a *admin)notify(){
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func sendNotification(n notifier){
	n.notify()
	n.useEmail()
}



func main(){
	bill := user2{"Bill", "bill@163.com"}
	lisa := admin{name:"Lisa", email:"lisa@163.com"}
	var n notifier
	n = &bill
	//sendNotification(&u)
	n.useEmail()
	n.notify()

	fmt.Printf("===============================\n\n")
	sendNotification(&bill)
	sendNotification(&lisa)


}
