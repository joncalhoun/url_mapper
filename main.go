package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var mappings map[string]string
	err := json.Unmarshal([]byte(data), &mappings)
	if err != nil {
		panic(err)
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if dest, ok := mappings[r.URL.Path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fmt.Println(r.URL.Path)
		http.Error(w, "Page not found", http.StatusNotFound)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

var data = `
{
  "/2.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-2.1-Building_the_server",
  "/3.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-3.1-Routing_with_if_else_statements",
  "/3.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-2.1-Building_the_server...book-3.1-Routing_with_if_else_statements",
  "/3.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-3.3-Using_the_gorilla_mux_router",
  "/3.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-3.1-Routing_with_if_else_statements...book-3.3-Using_the_gorilla_mux_router"
}
`
