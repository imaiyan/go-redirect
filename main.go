package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	url := os.Getenv("REDIRECT_URL")
	port := os.Getenv("LISTENING_PORT")

	mux := http.NewServeMux()

	mux.Handle("/", http.RedirectHandler(url, http.StatusSeeOther))

	fmt.Printf("Listening on port %s and redirecting to %s", port, url)

	http.ListenAndServe(":"+port, mux)
}
