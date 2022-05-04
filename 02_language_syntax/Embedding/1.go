package main

import "fmt"

type user struct {
	name string
	age  int
}

type admin struct {
	user  // embedding
	level string
}

func (u *user) notify() {
	fmt.Printf("user name is %s", u.name)
}

func main() {
	ad := admin{
		user:  user{age: 11, name: "john smith"},
		level: "slow",
	}
	ad.notify()
	// the same
	ad.user.notify()
}
