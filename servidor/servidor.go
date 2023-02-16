package servidor

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"godozero/database"
	"io"
	"log"
	"net/http"
	"strconv"
)

type usermd struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

type userNil struct {
	Msg string `json:"msg"`
}

// Criar usuário - insere usuário
func CreaterUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var user usermd

	if err := json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para json"))
		return
	}

	db, err := database.Conection()

	if err != nil {
		w.Write([]byte("Erro ao tentar conectar com o banco de dados"))
		return
	}

	defer db.Close()

	statment, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")

	if err != nil {
		w.Write([]byte("Erro ao tentar preparar o statment"))
		return
	}

	defer statment.Close()

	insert, err := statment.Exec(user.Nome, user.Email)

	if err != nil {
		w.Write([]byte("Erro ao tentar executar o statment"))
		return
	}

	idInsert, err := insert.LastInsertId()

	if err != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso %d", idInsert)))
}

func FindUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao extrair parâmetros"))
		return
	}

	db, err := database.Conection()

	if err != nil {
		w.Write([]byte("Erro ao tentar conectar com o banco de dados"))
		return
	}

	defer db.Close()

	userDB, err := db.Query("select * from usuarios where id = ?", ID)

	if err != nil {
		w.Write([]byte("Erro ao buscar o usuario"))
		return
	}

	var user usermd
	if userDB.Next() {
		if err := userDB.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
			w.Write([]byte("Erro ao executar scan"))
			return
		}
	}

	var usernil userNil
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		usernil.Msg = "Usuário não se encontra em nosso banco de dados"
		if err := json.NewEncoder(w).Encode(usernil); err != nil {
			w.Write([]byte("Erro ao converter os usuários para JSON"))
			return
		}
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

func FindUsers(w http.ResponseWriter, r *http.Request) {

	db, err := database.Conection()

	if err != nil {
		w.Write([]byte("Erro ao tentar conectar com o banco de dados"))
		return
	}

	defer db.Close()

	usersDB, err := db.Query("select * from usuarios")

	if err != nil {
		w.Write([]byte("Erro ao buscar os usuarios"))
		return
	}

	defer usersDB.Close()

	var users []usermd

	for usersDB.Next() {
		var user usermd
		if err = usersDB.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
			log.Fatal(err)
			w.Write([]byte("Erro ao tentar executar scan"))
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao extrair parâmetros"))
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Erro ao recuperar o body da requisilção"))
		return
	}

	var usermd usermd
	if err := json.Unmarshal(body, &usermd); err != nil {
		w.Write([]byte("Erro ao tentar converter o usuário"))
		return
	}

	db, err := database.Conection()

	if err != nil {
		w.Write([]byte("Erro ao tentar conectar com o banco de dados"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")

	if err != nil {
		w.Write([]byte("Erro no Prepare do statement"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(usermd.Nome, usermd.Email, ID); err != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao extrair parâmetros"))
		return
	}

	db, err := database.Conection()

	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")

	if err != nil {
		w.Write([]byte("Erro no Prepare do statement"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
