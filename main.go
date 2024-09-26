package main

import (
	"net/http"
	"os"
)

func main() {

	url := os.Getenv("REDIRECT_URL")
	port := os.Getenv("LISTENING_PORT")

	mux := http.NewServeMux()

	mux.Handle("/", http.RedirectHandler(url, http.StatusSeeOther))

	http.ListenAndServe(":"+port, mux)
}
