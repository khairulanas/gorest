package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book is magic
type Book struct {
	ID        string `json:"ID"`
	Title     string `json:"Title"`
	Author    string `json:"Author"`
	Publisher string `json:"Publisher"`
	Price     string `json:"Price"`
}

// Books is magic
var Books []Book

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: return all book")
	json.NewEncoder(w).Encode(Books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: return single book")
	vars := mux.Vars(r)
	key := vars["ID"]

	for _, book := range Books {
		if book.ID == key {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: create book")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: delete book")
	vars := mux.Vars(r)
	key := vars["ID"]

	for index, book := range Books {
		if book.ID == key {
			Books = append(Books[:index], Books[index+1:]...)
			fmt.Fprintf(w, "deleted book with id "+key)
		}
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: update book")
	vars := mux.Vars(r)
	key := vars["ID"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var bookbody Book
	json.Unmarshal(reqBody, &bookbody)

	for index, book := range Books {
		if book.ID == key {
			Books[index] = Book{ID: key, Title: bookbody.Title, Author: bookbody.Author, Publisher: bookbody.Publisher, Price: bookbody.Price}
			json.NewEncoder(w).Encode(Books[index])
		}
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "wellcome to the homepage")
	fmt.Println("endpoint hit: homepage")
}

func handleRequest() {
	myrouter := mux.NewRouter().StrictSlash(true)

	myrouter.HandleFunc("/", homePage)
	myrouter.HandleFunc("/books", returnAllBooks)
	myrouter.HandleFunc("/book", createBook).Methods("POST")
	myrouter.HandleFunc("/book/{ID}", updateBook).Methods("PUT")
	myrouter.HandleFunc("/book/{ID}", deleteBook).Methods("DELETE")
	myrouter.HandleFunc("/book/{ID}", returnSingleBook)

	log.Fatal(http.ListenAndServe(":10000", myrouter))
}

func main() {
	Books = []Book{
		Book{ID: "1", Title: "Origin", Author: "Dan Bwrown", Publisher: "Mizan", Price: "80000"},
		Book{ID: "2", Title: "Supernova", Author: "Dee Lestari", Publisher: "Bentang", Price: "80000"},
	}
	handleRequest()
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// // Article is magic
// type Article struct {
// 	ID      string `json:"ID"`
// 	Title   string `json:"Title"`
// 	Desc    string `json:"Desc"`
// 	Content string `json:"Content"`
// }

// // Articles is magic
// var Articles []Article

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("endpoint hit: return all article")
// 	json.NewEncoder(w).Encode(Articles)
// }

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("endpoint hit: return single article")
// 	vars := mux.Vars(r)
// 	key := vars["ID"]

// 	for _, article := range Articles {
// 		if article.ID == key {
// 			json.NewEncoder(w).Encode(article)
// 		}
// 	}
// }

// func createArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("endpoint hit: create article")
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var article Article
// 	json.Unmarshal(reqBody, &article)
// 	Articles = append(Articles, article)
// 	json.NewEncoder(w).Encode(article)
// }

// func deleteArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("endpoint hit: delete article")
// 	vars := mux.Vars(r)
// 	key := vars["ID"]

// 	for index, article := range Articles {
// 		if article.ID == key {
// 			Articles = append(Articles[:index], Articles[index+1:]...)
// 		}
// 	}
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "wellcome to the homepage")
// 	fmt.Println("endpoint hit: homepage")
// }

// func handleRequest() {
// 	myrouter := mux.NewRouter().StrictSlash(true)

// 	myrouter.HandleFunc("/", homePage)
// 	myrouter.HandleFunc("/articles", returnAllArticles)
// 	myrouter.HandleFunc("/article", createArticle).Methods("POST")
// 	myrouter.HandleFunc("/article/{ID}", deleteArticle).Methods("DELETE")
// 	myrouter.HandleFunc("/article/{ID}", returnSingleArticle)

// 	log.Fatal(http.ListenAndServe(":10000", myrouter))
// }

// func main() {
// 	Articles = []Article{
// 		Article{ID: "1", Title: "hello", Desc: "anu", Content: "anunya"},
// 		Article{ID: "2", Title: "hello", Desc: "anu", Content: "anunya"},
// 		Article{ID: "3", Title: "hello", Desc: "anu", Content: "anunya"},
// 	}
// 	handleRequest()
// }
