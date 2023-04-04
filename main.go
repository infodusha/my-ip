package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	_, err := io.WriteString(w, fmt.Sprintf("Your ip: %s", ip))
	if err != nil {
		log.Fatal(err)
	}
}
