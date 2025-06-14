package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Creating book: %s\n", vars["title"])
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Reading book: %s\n", vars["title"])
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Updating book: %s\n", vars["title"])
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Deleting book: %s\n", vars["title"])
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Book on host matched: %s\n", vars["tible"])
}

func SecureHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You are using HTTPS.")
}

func InsecureHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You are using HTTP.")
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Showing all books.")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Subrouter: Getting book: %s\n", vars["title"])
}

func main() {
    r := mux.NewRouter()

    // HTTP Methods
    r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
    r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
    r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
    r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

    // Hostname
    r.HandleFunc("/books/{title}", BookHandler).Methods("localhost")

    // Schemes
    r.HandleFunc("/secure", SecureHandler).Schemes("https")
    r.HandleFunc("/insecure", InsecureHandler).Methods("http")

    // Subrouter
    bookrouter := r.PathPrefix("/library").Subrouter()
    bookrouter.HandleFunc("/", AllBooks)
    bookrouter.HandleFunc("/{title}", GetBook)

    fmt.Println("Servidor corriendo en http://localhost:80")
    http.ListenAndServe(":80", r) // Change port here if you want other than 80
}
