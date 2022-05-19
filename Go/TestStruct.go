package main

import "fmt"

type Book struct {
	title   string
	author  string
	book_id int32
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Book{
		"go",
		"wangxiao",
		123,
	})
	// key : value 形式赋值
	new_book := Book{
		title:   "new",
		author:  "wangxiap",
		book_id: 1 << 2,
	}
	fmt.Println(new_book)
	// 忽略字段
	old_book := Book{
		title:  "new",
		author: "wangxiap",
	}
	fmt.Println(old_book)

	var test_book Book
	test_book.book_id = 1
	test_book.title = "test"
	test_book.author = "hello"

	printBook(&test_book)
}

func printBook(book *Book) {
	fmt.Println(book.book_id)
	fmt.Println(book.title)
	fmt.Println(book.author)
}
