package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", getIndex)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type IndexTpl struct {
	IP string
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.html").ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	ip := getIp(r)
	indexTpl := IndexTpl{ip}
	err = t.Execute(w, indexTpl)
	if err != nil {
		log.Fatal(err)
	}
}

func getIp(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}
