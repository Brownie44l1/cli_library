package storage

import (
	"database/sql"
	"github.com/Brownie44l1/cli_library/library"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStore struct {
    db *sql.DB
}

func NewSQLiteStore(dbPath string) (*SQLiteStore, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS books (
            isbn TEXT PRIMARY KEY,
            title TEXT,
            author TEXT,
            year TEXT,
            available INTEGER
        )
    `)
    if err != nil {
        return nil, err
    }

    return &SQLiteStore{db: db}, nil
}

func (s *SQLiteStore) Add(book library.Book) error {
    _, err := s.db.Exec("INSERT INTO books VALUES (?, ?, ?, ?, ?)",
        book.ISBN, book.Title, book.Author, book.Year, book.Available)
    return err
}

func (s *SQLiteStore) Update(book library.Book) error {
	_, err := s.db.Exec("UPDATE books SET title=?, author=?, year=?, available=? WHERE isbn=?",
		book.Title, book.Author, book.Year, book.Available, book.ISBN)
	return err
}

func (s *SQLiteStore) List() ([]library.Book, error) {
	rows, err := s.db.Query("SELECT isbn, title, author, year, available FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []library.Book
	for rows.Next() {
		var book library.Book
		err := rows.Scan(&book.ISBN, &book.Title, &book.Author, &book.Year, &book.Available)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *SQLiteStore) Get(isbn string) (*library.Book, error) {
	var book library.Book
	err := s.db.QueryRow("SELECT isbn, title, author, year, available FROM books WHERE isbn=?", isbn).
		Scan(&book.ISBN, &book.Title, &book.Author, &book.Year, &book.Available)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *SQLiteStore) Delete(isbn string) error {
	_, err := s.db.Exec("DELETE FROM books WHERE isbn=?", isbn)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLiteStore) Borrow(isbn string) error {
	_, err := s.db.Exec("UPDATE books SET available=0 where isbn=?", isbn)
	return err
}

func (s *SQLiteStore) Return(isbn string) error {
	_, err := s.db.Exec("UPDATE books SET available=1 where isbn=?", isbn)
	return err
}
