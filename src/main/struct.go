package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book Book
	book.title = "Web Application with NodeJs AngularJs and MongoDB"
	book.author = "(Us)James Bay"
	book.subject = "WebApplication"
	book.book_id = 1
	printStout(book)
	printPointerStout(&book)
}

func printStout(book Book) {
	fmt.Printf("book title: %s\n", book.title)
	fmt.Printf("book author: %s\n", book.author)
	fmt.Printf("book subject: %s\n", book.subject)
	fmt.Printf("book id: %d\n", book.book_id)
}

func printPointerStout(book *Book) {
	fmt.Printf("book title: %s\n", book.title)
	fmt.Printf("book author: %s\n", book.author)
	fmt.Printf("book subject: %s\n", book.subject)
	fmt.Printf("book id: %d\n", book.book_id)
}
