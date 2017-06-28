package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var mappings map[string]string
	f, err := os.Open("./map.json")
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(&mappings)
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
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
