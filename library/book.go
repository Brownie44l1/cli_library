package library

type Book struct {
    ISBN      string `json:"isbn"`
    Title     string `json:"title"`
    Author    string `json:"author"`
    Year      string `json:"year"`
    Available bool   `json:"available"`
}