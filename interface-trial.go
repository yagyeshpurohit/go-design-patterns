package main

import "fmt"

type Animal interface {
	eat() bool
	sleep() bool
}

type Dog struct {
	name  string
	breed string
}

func (d Dog) eat() bool {
	return true
}
func (d Dog) sleep() bool {
	return false
}

func main() {
	var dog1 Animal
	dog1 = Dog{
		name:  "Tiger",
		breed: "Doberman",
	}
	fmt.Println(dog1.sleep())
}
