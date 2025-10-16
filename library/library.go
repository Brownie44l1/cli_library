package library

type Library struct {
	Books map[string]*Book 
}

func NewLibrary() *Library {
	return &Library{Books: make(map[string]*Book)}
}

func (l *Library) AddBook(isbn, title, author, year string) {

}