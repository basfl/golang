package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_, ok := params["id"]
	if !ok {
		log.Fatal("could not find id ")
	}
	fmt.Println("id==", params["id"])
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			//	return

		}

	}
	//json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	w.Header().Set("Content-Type", "application/json")
	max := 0
	for _, elem := range books {
		if id, _ := strconv.Atoi(elem.ID); id > max {
			max = id
		}
	}
	book.ID = strconv.Itoa(max + 1)
	books = append(books, book)
	writeToFile()
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	params := mux.Vars(r)
	_, ok := params["id"]
	if !ok {
		log.Fatal("could not find id ")
	}
	for index, elem := range books {
		if elem.ID == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&book)
			books = append(books[:index], books[index+1:]...)
			book.ID = params["id"]
			books = append(books, book)
			writeToFile()
			break
		}
	}

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_, ok := params["id"]
	if !ok {
		log.Fatal("could not find id ")
	}
	for index, elem := range books {
		if elem.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			writeToFile()
			break
		}
	}

}

func writeToFile() {

	newContent, _ := json.MarshalIndent(books, "", " ")

	err := ioutil.WriteFile("books.json", newContent, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
