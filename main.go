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

	fmt.Printf("Listening on %s:%s", url, port)

	http.ListenAndServe(":"+port, mux)
}
