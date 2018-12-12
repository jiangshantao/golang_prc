package main

import (
	"fmt"
)

type notifier2 interface {
	notify2()
}

type user2 struct {
	name string
	email string
}

type admin2 struct {
	user2
	level int
}

func (u *user2)notify2(){
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func (a *admin2)notify2(){
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func sendNotification2(n notifier2){
	n.notify2()
}

func main(){
	ad := admin2{
		user2: user2{
			name:"jst",
			email:"jst@bytedance.com",
		},
		level: 1,
	}

	ad.user2.notify2()

	ad.notify2()

	sendNotification2(&ad)

}
