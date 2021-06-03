package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc:"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"Test Title", Desc: "Test Description", Content: "Hello World"},
	}
	json.NewEncoder(w).Encode(articles);
	fmt.Println("GET endpoint hit: allArticles")
}

func testPostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "POST endpoint hit: testPostArticles")
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", allArticles).Methods("GET")
	router.HandleFunc("/articles", testPostArticles).Methods("POST")
	
    log.Fatal(http.ListenAndServe(":4200", router))
}

func main() {
    handleRequests()
}