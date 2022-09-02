package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Seja bem vindo ao nosso site!"))
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
