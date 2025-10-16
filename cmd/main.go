package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"github.com/google/shlex"
)

type Library struct {
	books map[string]*Book 
}

func (l *Library) addBook(parts []string) {
	if len(parts) < 5 {
		fmt.Println("USAGE: add {isbn} {title} {author} {year}")
		return
	}

	isbn := parts[1]
	title := parts[2]
	author := parts[3]
	year := parts[4]
	
	l.books[isbn] = &Book{ISBN: isbn, Title: title, Author: author, Year: year, Available: true}
	fmt.Println("Book Added: ", l.books[isbn])
}

func (l *Library) updateBook(parts []string) {
	if len(parts) < 6 {
		fmt.Println("USAGE: update {isbn} {title} {author} {year} {available}")
		return
	}
	isbn := parts[1]
	title := parts[2]
	author := parts[3]
	year := parts[4]
	available, err := strconv.ParseBool(parts[5])

	if err != nil {
		fmt.Println("Invalid availability value (use true/false)")
		return
	}

	if book, ok := l.books[isbn]; ok {
		if title != "" {
			book.Title = title
		}
		if author != "" {
			book.Author = author
		}
		if year != "" {
			book.Year = year
		}
		book.Available = available
		fmt.Println("Book Updated:", book)
	} else {
		fmt.Println("Book not found for ISBN:", isbn)
		return
	}
}

func (l *Library) deleteBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: delete {isbn}")
		return
	}
	isbn := parts[1]

	if _, exists := l.books[isbn]; exists {
		fmt.Println("Book delected: ", l.books[isbn])
		delete(l.books, isbn)
	} else {
		fmt.Println("Book not found")
	}
}

func (l *Library) displayBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: display {isbn}")
		return
	}
	isbn := parts[1]
	if book, exists := l.books[isbn]; exists {
		fmt.Printf("Key: %s -> {Title: %s, Author: %s, Year: %s, Available: %t}\n", book.ISBN, book.Title, book.Author, book.Year, book.Available)
	} else {
		fmt.Println("Book not found")
	}
}

func (l *Library) listBooks() {
	fmt.Println("Displaying all books")
	for isbn, book := range l.books {
		fmt.Printf("Key: %s -> {Title: %s, Author: %s, Year: %s, Available: %t}\n", isbn, book.Title, book.Author, book.Year, book.Available)
	}
}

func (l *Library) borrowBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: borrow {isbn}")
		return
	}

	isbn := parts[1]

	if book, ok := l.books[isbn]; ok {
		book.Available = false
		fmt.Println("Book borrowed: ", book.Title)
	} else {
		fmt.Println("Book does not exist")
	}
}

func (l *Library) returnBook(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: return {isbn}")
		return
	}

	isbn := parts[1]

	if book, ok := l.books[isbn]; ok {
		book.Available = true
		fmt.Println("Book Returned: ", book.Title)
	} else {
		fmt.Println("Book does not exist")
	}
}

func (l *Library) load(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: load {filename}")
		return
	}
	filename := parts[1]
	l.readFromFile(filename)
}

func (l *Library) save(parts []string) {
	if len(parts) < 2 {
		fmt.Println("USAGE: save {filename}")
		return
	}
	filename := parts[1]
	l.saveToFile(filename)
}

func (l *Library) saveToFile(filename string) {
	data, errJson := json.MarshalIndent(l.books, "", "  ")
	if errJson != nil {
		fmt.Println("Error marshaling data:", errJson)
		return
	}

	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}

	fmt.Println("Library saved to", filename)
}

func (l *Library) readFromFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No existing library file found, starting fresh.")
			return
		}
		fmt.Println("Error reading file:", err)
		return
	}

	if l.books == nil {
		l.books = make(map[string]*Book)
	}

	errJson := json.Unmarshal(data, &l.books)
	if errJson != nil {
		fmt.Println("failed to decode json:", errJson)
	}
	fmt.Println("Library loaded from", filename)
}

func main() {
	lib := &Library{books: make(map[string]*Book)}
	lib.readFromFile("library.json")
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Println("Enter Command: ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		text, err := shlex.Split(line)
		if err != nil {
			fmt.Println("parse error:", err)
			continue
		}
		cmd := text[0]

		switch cmd {
		case "add":
			lib.addBook(text)
		case "update":
			lib.updateBook(text)
		case "delete":
			lib.deleteBook(text)
		case "display":
			lib.displayBook(text)
		case "list":
			lib.listBooks()
		case "borrow":
			lib.borrowBook(text)
		case "return":
			lib.returnBook(text)
		case "load":
			lib.load(text)
		case "save":
			lib.save(text)
		case "quit":
			lib.save([]string{"save", "library.json"})
			fmt.Println("Exiting....")
			return
		default:
			fmt.Println("Unknown Command")
		}
		fmt.Println()
	}
}