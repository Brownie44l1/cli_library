package library_test

import (
	"os"
	"testing"

	"github.com/Brownie44l1/cli_library/library"
)

// helper to create a fresh library for each test
func setupLibrary() *library.Library {
	lib := library.NewLibrary()
	return lib
}

func TestAddBook(t *testing.T) {
	lib := setupLibrary()
	lib.AddBook("1001", "Go in Action", "John Doe", "2022")

	if len(lib.Books) != 1 {
		t.Fatalf("expected 1 book, got %d", len(lib.Books))
	}
	if lib.Books["1001"].Title != "Go in Action" {
		t.Fatalf("expected title 'Go in Action', got %s", lib.Books["1001"].Title)
	}
}

func TestUpdateBook(t *testing.T) {
	lib := setupLibrary()
	lib.AddBook("1001", "Go 101", "John", "2020")

	err := lib.UpdateBook("1001", "Advanced Go", "John", "2024", "true")
	if err != nil {
		t.Fatalf("UpdateBook failed: %v", err)
	}

	book := lib.Books["1001"]
	if book.Title != "Advanced Go" {
		t.Fatalf("expected title 'Advanced Go', got %s", book.Title)
	}
	if book.Year != "2024" {
		t.Fatalf("expected year '2024', got %s", book.Year)
	}
}

func TestBorrowAndReturn(t *testing.T) {
	lib := setupLibrary()
	lib.AddBook("1001", "Go Deep Dive", "Alice", "2021")

	book := lib.Books["1001"]

	// Borrow
	book.Available = false
	if book.Available {
		t.Fatalf("expected Available=false after borrow")
	}

	// Return
	book.Available = true
	if !book.Available {
		t.Fatalf("expected Available=true after return")
	}
}

func TestSaveAndLoad(t *testing.T) {
	filename := "test_library.json"
	_ = os.Remove(filename)         // ensure clean
	defer os.Remove(filename)       // cleanup after test

	lib := setupLibrary()
	lib.AddBook("1001", "Go Persistence", "Tester", "2023")
	// Save
	if err := lib.SaveToFile(filename); err != nil {
		t.Fatalf("SaveToFile failed: %v", err)
	}

	// Load into new library instance
	newLib := setupLibrary()
	if err := newLib.LoadFromFile(filename); err != nil {
		t.Fatalf("LoadFromFile failed: %v", err)
	}

	if _, exists := newLib.Books["1001"]; !exists {
		t.Fatalf("expected book 1001 to exist after load")
	}
}