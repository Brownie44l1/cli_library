package library

import (
	"fmt"
	"strconv"
)

type Library struct {
	store BookStore
}

func NewLibrary(store BookStore) *Library {
	return &Library{store: store}
}

func (l *Library) AddBook(isbn, title, author, year string) error {
	if len(isbn) == 0 || len(isbn) > 4 {
		return fmt.Errorf("ISBN must be 1-4 integers")
	}

	yearInt, err := strconv.Atoi(year)
	if err != nil || yearInt < 1800 || yearInt > 2025 {
		return fmt.Errorf("Invalid year: must be between 1800-2025")
	}

	_, err = l.store.Get(isbn) 
	if err == nil {
		return fmt.Errorf("Book with ISBN %s already exists", isbn)
	}

	if isbn == "" || title == "" || author == "" || year == "" {
		return fmt.Errorf("All fields (ISBN, Title, Author, Year) are required")
	}

	book := Book{ISBN: isbn, Title: title, Author: author, Year: year, Available: true}
    return l.store.Add(book)
}

func (l *Library) UpdateBook(isbn, title, author, year, availableStr string) error {
	_, err := l.store.Get(isbn)
	if err != nil {
		return fmt.Errorf("Book not found for ISBN: %s", isbn)
	}
	available, err := strconv.ParseBool(availableStr)
	if err != nil {
		return fmt.Errorf("Invalid availability value (use true/false)")
	}
	
	book := Book {ISBN: isbn, Title: title, Author: author, Year: year, Available: available,}

	err = l.store.Update(book)
	if err != nil {
		return err
	}

	fmt.Println("Book updated: ", book)
	return nil
}

func (l *Library) DeleteBook(isbn string) error {
	book, err := l.store.Get(isbn)
	if err != nil {
		fmt.Println("Book not found")
		return err
	}

	err = l.store.Delete(isbn)
	if err != nil {
		return err
	}

	fmt.Println("Book Delected: ", book)
	return nil
}

func (l *Library) GetBook(isbn string) error {
	book, err := l.store.Get(isbn)
	if err != nil {
		fmt.Println("Book not found")
		return err
	} else {
		fmt.Printf("- %s by %s (%s) [Available: %t]\n", book.Title, book.Author, book.Year, book.Available)
	}
	return nil
}

func (l *Library) ListBooks() error {
	books, err := l.store.List()
	if err != nil {
		return nil
	}

	if len(books) == 0 {
		fmt.Println("No books in library")
		return nil
	}
	
	for _, book := range books {  
		fmt.Printf("- %s by %s (%s) [Available: %t]\n", book.Title, book.Author, book.Year, book.Available)
	}
	return nil
}

func (l *Library) BorrowBook(isbn string) error {
	book, err := l.store.Get(isbn) 
	if err != nil {
		return err
	}

	if !book.Available {
		fmt.Println("Book is already borrowed")
		return fmt.Errorf("book unavailable")
	}
	
	err = l.store.Borrow(isbn)
	if err != nil {
		return err
	}
	fmt.Println("Book borrowed successfully: ", book.Title)
	return nil
}

func (l *Library) ReturnBook(isbn string) error {
	book, err := l.store.Get(isbn) 
	if err != nil {
		return err
	}

	if book.Available {
		fmt.Println("Book is not borrowed, can't return")
		return fmt.Errorf("book available")
	}
	
	err = l.store.Return(isbn)
	if err != nil {
		return err
	}
	fmt.Println("Book returned successfully: ", book.Title)
	return nil
}

func (l *Library) SearchBooks(author string) error {
	book, err := l.store.Search(author)
	if err != nil {
		fmt.Printf("Author %s not found\n", author)
		return err
	} else {
		fmt.Printf("- %s by %s (%s) [Available: %t]\n", book.Title, book.Author, book.Year, book.Available)
	}
	return nil
}
