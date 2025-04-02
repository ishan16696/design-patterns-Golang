package decorator

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
}

func (b *Book) GetPrice() int {
	if b.Title == "Data book" {
		return 150
	}
	return 100
}

// Decorator which wraps the original base object
type BorrowableBook struct {
	Book       // Embed the original object: Book
	IsBorrowed bool
}

// Adds new behavior using decorator object
func (b *BorrowableBook) Borrow() error {
	if b.IsBorrowed {
		return fmt.Errorf("book '%s' is already borrowed", b.Title)
	}
	b.IsBorrowed = true
	fmt.Printf("Book '%s' borrowed\n", b.Title)
	return nil
}

// Adds new behavior using decorator object
func (b *BorrowableBook) Return() error {
	if !b.IsBorrowed {
		return fmt.Errorf("book '%s' is not borrowed", b.Title)
	}
	b.IsBorrowed = false
	fmt.Printf("Book '%s' returned\n", b.Title)
	return nil
}
