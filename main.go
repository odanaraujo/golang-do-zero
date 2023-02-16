package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"godozero/servidor"
	"log"
	"net/http"
)

func main() {
	//contem todas as rotas para fazer a interação com o banco de dados
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CreaterUser).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.FindUsers).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", servidor.FindUser).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", servidor.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/usuario/{id}", servidor.DeleteUser).Methods(http.MethodDelete)
	fmt.Println("listening at the door 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
