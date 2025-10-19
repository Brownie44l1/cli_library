package library

import (
	"fmt"
	"strconv"
)

type Library struct {
	Books map[string]*Book
}

func NewLibrary() *Library {
	return &Library{Books: make(map[string]*Book)}
}

func (l *Library) AddBook(isbn, title, author, year string) {
	l.Books[isbn] = &Book{ISBN: isbn, Title: title, Author: author, Year: year, Available: true}
	fmt.Println("Book added: ", l.Books[isbn])
}

func (l *Library) UpdateBook(isbn, title, author, year, availableStr string) error {
	book, ok := l.Books[isbn]
	if !ok {
		return fmt.Errorf("Book not found for ISBN: %s", book.ISBN)
	}
	Available, err := strconv.ParseBool(availableStr)
	if err != nil {
		return fmt.Errorf("Invalid availability value (use true/false)")
	}
	book.ISBN = isbn
	book.Title = title
	book.Author = author
	book.Year = year
	book.Available = Available
	fmt.Println("Book updated: ", book)
	return nil
}

func (l *Library) DeleteBook(isbn string) {
	if _, exists := l.Books[isbn]; exists {
		fmt.Println("Book delected: ", l.Books[isbn])
		delete(l.Books, isbn)
	} else {
		fmt.Println("Book not found")
	}
}

func (l *Library) DisplayBook(isbn string) {
	if book, exists := l.Books[isbn]; exists {
		fmt.Printf("ISBN: %s -> {Title: %s, Author: %s, Year: %s, Available: %t}\n", book.ISBN, book.Title, book.Author, book.Year, book.Available)
	} else {
		fmt.Println("Book not found")
	}
}

func (l *Library) ListBooks() {
	for isbn, book := range l.Books {
		fmt.Printf("ISBN: %s -> {Title: %s, Author: %s, Year: %s, Available: %t}\n", isbn, book.Title, book.Author, book.Year, book.Available)
	}
}

func (l *Library) BorrowBook(isbn string) {
	if book, ok := l.Books[isbn]; ok {
		if !book.Available {
			fmt.Println("Book has been borrowed, check back")
		} else {
			book.Available = false
			fmt.Println("Book borrowed: ", book.Title)
		}
	} else {
		fmt.Println("Book does not exist")
	}
}

func (l *Library) ReturnBook(isbn string) {
	if book, ok := l.Books[isbn]; ok {
		if book.Available{
			fmt.Println("Book not borrowed can't be returned")
		} else {
			book.Available = true
			fmt.Println("Book Returned: ", book.Title)
		}
	} else {
		fmt.Println("Book does not exist")
	}
}
