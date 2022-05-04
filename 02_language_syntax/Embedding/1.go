package main

import "fmt"

type user struct {
	name string
	age  int
}

type admin struct {
	person user // NOT Embedding
	level  string
}

func (u *user) notify() {
	fmt.Printf("user name is %s", u.name)
}

func main() {
	ad := admin{
		person: user{age: 11, name: "john smith"},
		level:  "slow",
	}
	ad.person.notify()
}
