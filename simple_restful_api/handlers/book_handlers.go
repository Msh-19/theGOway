package handlers

import (
	"encoding/json"
	"net/http"
	"simple_restful_api/models"
	"strconv"
)

var books = []models.Book{
	{ID: 1, Title: "1984", Author:"George Orwell", PublishedYear: 1949},
	{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", PublishedYear: 1960},	
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
    for _, book := range books {
        if book.ID == id {
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    var newBook models.Book
    json.NewDecoder(r.Body).Decode(&newBook)
    newBook.ID = len(books) + 1
    books = append(books, newBook)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newBook)
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
    var updatedBook models.Book
    json.NewDecoder(r.Body).Decode(&updatedBook)
    for i, book := range books {
        if book.ID == id {
            books[i] = updatedBook
            books[i].ID = id
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(books[i])
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}


func DeleteBook(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}
