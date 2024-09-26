package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.RedirectHandler("https://www.google.com/", http.StatusSeeOther))

	http.ListenAndServe(":8080", mux)
}
