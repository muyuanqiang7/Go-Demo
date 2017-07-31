package main

import "fmt"

type Phone interface {
	call()
}

type nokiaPhone struct {
}

func (phone nokiaPhone) call() {
	fmt.Println("i am nokia, i am call you")
}

type iPhone struct {
}

func (phone iPhone) call() {
	fmt.Println("i am iPhone, i am call you")
}

func main() {
	var phone Phone
	phone = new(nokiaPhone)
	phone.call()
	phone = new(iPhone)
	phone.call()
}