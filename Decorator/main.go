package main

import (
	"design-patterns-Golang/Decorator/decorator"
	"fmt"
)

func main() {
	// Original Book
	book := decorator.Book{
		ID: 1, Title: "Data book",
		Author: "Alan",
	}

	// Decorated Book (now supports borrowing)
	borrowableBook := decorator.BorrowableBook{Book: book}

	// Use new behavior
	borrowableBook.Borrow()
	borrowableBook.Return()

	fmt.Printf("Price of the book :%d\n", borrowableBook.GetPrice())
}
