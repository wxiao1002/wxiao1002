package main

import "fmt"

func main() {

}

// 定义一个error 结构体
type DivideError struct {
	devidee int
	devider int
}

// 实现接口
func (de *DivideError) Error() string {
	strFormat := "Cannot processed,the divider is zero,devidee:%d"
	return fmt.Sprintln(strFormat, de.devidee)
}

// 定义函数
func Divide(dividee int, divider int) (int, string) {
	if divider == 0 {
		error := DivideError{
			dividee, divider,
		}

		return 0, error.Error()
	}
	return dividee / divider, ""
}
