package servidor

import (
	"encoding/json"
	"fmt"
	"godozero/banco"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Pessoa struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var pessoa Pessoa

	if erro = json.Unmarshal(corpoRequisicao, &pessoa); erro != nil {
		w.Write([]byte("Erro ao converter  o usuário para o struct"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into pessoas (nome, email) values (?, ?)")

	if erro != nil {
		w.Write([]byte("Erro ao tentar preparar o banco de dados"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(pessoa.Nome, pessoa.Email)

	if erro != nil {
		fmt.Println(erro)
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	_, erro = insercao.LastInsertId()

	if erro != nil {
		w.Write([]byte("Erro ao recuperar o id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário %s inserido com sucesso", pessoa.Nome)))
}

func BuscarPessoas(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao tentar conectar ao banco"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("select * from pessoas")

	if erro != nil {
		w.Write([]byte("Erro ao executar a Query"))
		return
	}

	defer linhas.Close()

	var pessoas []Pessoa
	for linhas.Next() {

		var pessoa Pessoa

		if erro := linhas.Scan(&pessoa.ID, &pessoa.Nome, &pessoa.Email); erro != nil {
			w.Write([]byte("Erro no Scan"))
			return
		}

		pessoas = append(pessoas, pessoa)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(pessoas); erro != nil {
		w.Write([]byte("Erro ao converter pessoas para json"))
		return
	}
}

func BuscarPessoa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o ID para uint"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao abrir o banco de daods"))
		return
	}

	linha, erro := db.Query("select * from pessoas where id = ?", ID)

	if erro != nil {
		w.Write([]byte("Erro ao realizara consulta no banco de dados"))
		return
	}

	var pessoa Pessoa
	if linha.Next() {
		if erro := linha.Scan(&pessoa.ID, &pessoa.Nome, &pessoa.Email); erro != nil {
			w.Write([]byte("Erro ao escanear pessoa"))
			return
		}
	}

	if erro = json.NewEncoder(w).Encode(pessoa); erro != nil {
		w.Write([]byte("Erro ao converter pessoas para json"))
		return
	}
}

func AtualizarPessoa(w http.ResponseWriter, r *http.Request) {

	fmt.Println("chegou aqui")
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o ID para uint"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var pessoa Pessoa
	if erro := json.Unmarshal(corpoRequisicao, &pessoa); erro != nil {
		w.Write([]byte("Erro ao converter  o usuário para o struct"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao abrir o banco de dados"))
		return
	}

	defer db.Close()

	statemant, erro := db.Prepare("update pessoas set nome = ?, email = ? where id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	defer statemant.Close()

	_, erro = statemant.Exec(pessoa.Nome, pessoa.Email, ID)

	if erro != nil {
		w.Write([]byte("Erro ao atualizar pessoa"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarPessoa(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o ID para uint"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao abrir o banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("delete from pessoas where id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	_, erro = statement.Exec(ID)

	if erro != nil {
		w.Write([]byte("Erro ao deletar pessoa"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
