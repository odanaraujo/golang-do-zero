package main

import (
	"fmt"
	"godozero/servidor"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/pessoas", servidor.CriarUsuario).Methods(http.MethodPost)
	r.HandleFunc("/pessoas", servidor.BuscarPessoas).Methods(http.MethodGet)
	r.HandleFunc("/pessoa/{id}", servidor.BuscarPessoa).Methods(http.MethodGet)
	r.HandleFunc("/pessoa/{id}", servidor.AtualizarPessoa).Methods(http.MethodPut)
	r.HandleFunc("/pessoa/{id}", servidor.DeletarPessoa).Methods(http.MethodDelete)
	fmt.Println("Escutando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
