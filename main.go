package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	ip := getIp(r)
	_, err := io.WriteString(w, fmt.Sprintf("Your ip4: %s", ip))
	if err != nil {
		log.Fatal(err)
	}
}

func getIp(r *http.Request) string {
	return strings.Split(r.RemoteAddr, ":")[0]
}
