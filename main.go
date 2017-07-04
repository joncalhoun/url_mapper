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
  "/2.1": "https://github.com/joncalhoun/lenslocked.com/tree/book-2.1-Building_the_server",
  "/3.1": "https://github.com/joncalhoun/lenslocked.com/tree/book-3.1-Routing_with_if_else_statements",
  "/3.1-diff": "https://github.com/joncalhoun/lenslocked.com/compare/book-2.1-Building_the_server...book-3.1-Routing_with_if_else_statements",
  "/3.3": "https://github.com/joncalhoun/lenslocked.com/tree/book-3.3-Using_the_gorilla_mux_router",
  "/3.3-diff": "https://github.com/joncalhoun/lenslocked.com/compare/book-3.1-Routing_with_if_else_statements...book-3.3-Using_the_gorilla_mux_router",
  "/4.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-4.4-Creating_a_template",
  "/4.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-3.3-Using_the_gorilla_mux_router...book-4.4-Creating_a_template",
  "/4.5": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-4.5-Contextual_encoding",
  "/4.5-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-4.4-Creating_a_template...book-4.5-Contextual_encoding",
  "/6.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.1-The_home_template",
  "/6.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-4.5-Contextual_encoding...book-6.1-The_home_template",
  "/6.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.2-The_contact_template",
  "/6.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.1-The_home_template...book-6.2-The_contact_template",
  "/6.3.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.3.1-Footer_template",
  "/6.3.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.2-The_contact_template...book-6.3.1-Footer_template",
  "/6.3.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.3.2-Creating_the_view_type",
  "/6.3.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.3.1-Footer_template...book-6.3.2-Creating_the_view_type",
  "/6.3.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.3.3-Bootstrap_layout",
  "/6.3.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.3.2-Creating_the_view_type...book-6.3.3-Bootstrap_layout",
  "/6.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.4-Navbar",
  "/6.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.3.3-Bootstrap_layout...book-6.4-Navbar",
  "/6.5.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.5.2-Globbing",
  "/6.5.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.4-Navbar...book-6.5.2-Globbing",
  "/6.5.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.5.3-View_rendering",
  "/6.5.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.5.2-Globbing...book-6.5.3-View_rendering",
  "/6.5.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.5.4-Move_footer",
  "/6.5.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.5.3-View_rendering...book-6.5.4-Move_footer",
  "/7.1.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.1.1-Sign_up_form",
  "/7.1.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.5.4-Move_footer...book-7.1.1-Sign_up_form",
  "/7.1.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.1.2-Form_panel",
  "/7.1.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.1.1-Sign_up_form...book-7.1.2-Form_panel",
  "/7.1.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.1.3-Sign_up_link",
  "/7.1.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.1.2-Form_panel...book-7.1.3-Sign_up_link"
}
`
