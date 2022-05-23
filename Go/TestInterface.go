package main

import "fmt"

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(Xiaomi)
	phone.call()
}

type Phone interface {
	call()
}
type NokiaPhone struct {
}

func (nokia NokiaPhone) call() {
	fmt.Println("Nokia call")
}

type Xiaomi struct {
}

func (xiaomi Xiaomi) call() {
	fmt.Println("xiaomi call")
}
