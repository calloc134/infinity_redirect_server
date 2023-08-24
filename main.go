package main

import (
	"log"
	"net/http"
)

// 無限にリダイレクトを繰り返すもの
func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://172.20.6.8:8095/", http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8095", nil))
}
