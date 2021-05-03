package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// mock data as slice of type book struct
var books []Book

// author struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//book struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// get multiple books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//   get params
	params := mux.Vars(r)

	//   loop through books
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// if no id matches
	json.NewEncoder(w).Encode(Book{})

}

// create book
func cerateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// create book variable
	var book Book

	// decode body and pass to book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	// add content json to header
	w.Header().Set("Content-Type", "content/json")

	// get all params
	params := mux.Vars(r)

	//trvaerse through all the books
	for idx, item := range books {

		// check for item.ID is equal to updation id
		if item.ID == params["id"] {
			//    delete the book
			books := append(books[:idx], books[idx+1:]...)
			var currBook Book
			// decode body and sore the value in currbook
			_ = json.NewDecoder(r.Body).Decode(&currBook)

			// push new id
			currBook.ID = strconv.Itoa(rand.Intn(100000))

			//added new book
			books = append(books, currBook)

			json.NewEncoder(w).Encode(books)
		}
	}
	json.NewEncoder(w).Encode(books)

}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	//set content type
	w.Header().Set("Content-Type", "content/json")

	//get params
	params := mux.Vars(r)

	//iterate through items and delete the given id
	for idx, item := range books {
		if item.ID == params["id"] {
			//    create new book except curr idx
			books := append(books[:idx], books[idx+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}

	//  if book id doesnt exist
	json.NewEncoder(w).Encode(&Book{})
	return
}

func main() {
	//  init router
	router := mux.NewRouter()

	// appending to books
	books = append(books, Book{ID: "1", Isbn: "1234", Title: "book1", Author: &Author{FirstName: "sai", LastName: "duddukuri"}})
	books = append(books, Book{ID: "2", Isbn: "1235", Title: "book2", Author: &Author{FirstName: "sai", LastName: "duddukuri"}})
	books = append(books, Book{ID: "3", Isbn: "1236", Title: "book3", Author: &Author{FirstName: "sai", LastName: "duddukuri"}})

	// router handlers
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", cerateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
