package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	username string
	email    string
}

func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.username, u.email)
}

func main() {
	u := user{
		"mr",
		"1597888126@qq.com",
	}
	// You're trying to create a copy of T,
	// but T doesn't have any methods;
	// sendNotification(u)
	// let's share data
	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}
