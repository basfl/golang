package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//initializing dummy values for books
	// books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	//reading data from json file
	content, _ := ioutil.ReadFile("books.json")
	_ = json.Unmarshal([]byte(content), &books)
	//create a router
	r := mux.NewRouter()
	//Router handler endpoint:
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	//	r.HandleFunc("/api/books", getBooks).Methods("GET")
	fmt.Printf("%T", r)
	log.Fatal(http.ListenAndServe(":8000", r))

}
