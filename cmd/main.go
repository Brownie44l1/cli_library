package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Brownie44l1/cli_library/library"
)

func main() {
	lib := library.NewLibrary()
	lib.LoadFromFile("library.json")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter Command: ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
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

		case "save":
			lib.SaveToFile("library.json")

		case "quit":
			lib.SaveToFile("library.json")
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command")
		}
	}
}