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
	connectingIp := r.Header.Get("CF-Connecting-IP")
	if connectingIp != "" {
		return connectingIp
	}
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}
	parts := strings.Split(r.RemoteAddr, ":")
	if len(parts) == 0 {
		return ""
	}
	end := len(parts) - 1
	return strings.Join(parts[:end], ":")
}
