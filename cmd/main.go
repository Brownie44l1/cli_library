package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Brownie44l1/cli_library/library"
	"github.com/Brownie44l1/cli_library/storage"
	"github.com/google/shlex"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	store, err := storage.NewSQLiteStore("library.db")
	if err != nil {
		log.Fatal(err)
	}
	lib := library.NewLibrary(store)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter Command: ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		parts, err := shlex.Split(line)
		if err != nil {
			fmt.Println("Error parsing command:", err)
			continue
		}

		switch parts[0] {
		case "add":
			if len(parts) < 5 {
				fmt.Println("Usage: add {isbn} {title} {author} {year}")
				continue
			}
			lib.AddBook(parts[1], parts[2], parts[3], parts[4])

		case "update":
			if len(parts) < 6 {
				fmt.Println("Usage: update {isbn} {title} {author} {year} {available}")
				continue
			}
			err := lib.UpdateBook(parts[1], parts[2], parts[3], parts[4], parts[5])
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "list":
			lib.ListBooks()

		case "get":
			if len(parts) < 2 {
				fmt.Println("Usage: get {isbn}")
				continue
			}
			lib.GetBook(parts[1])

		case "search":
			if len(parts) < 2 {
				fmt.Println("Usage: search {author}")
				continue
			}
			lib.SearchBooks(parts[1])

		case "borrow":
			lib.BorrowBook(parts[1])

		case "return":
			lib.ReturnBook(parts[1])

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command")
		}
	}
}
