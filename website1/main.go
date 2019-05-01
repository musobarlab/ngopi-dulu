package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8000", nil)
}
