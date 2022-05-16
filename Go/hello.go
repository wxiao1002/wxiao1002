package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")
	fmt.Println("Hello" + " world!")
	var age int
	age = 10
	fmt.Println(age)
	var stockcode = 123
	var enddate = "2020-12-12"
	var url = "Code=%d,endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
}
