package server

import (
	"CRUD/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// This file is related with the Methods (POST - Creat, GET - Read, PUT - Update and DELETE - Delete)
// Create an user - Insert an user on the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser chamado")

	bodyReaquest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var user user

	if erro = json.Unmarshal(bodyReaquest, &user); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
	}

	// Using the function Connection() from the database.go to create the connection with the database
	db, erro := database.Connection()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao Banco de dados"))
		return
	}

	defer db.Close()

	// Insert the data on the database
	// Prepare statement
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(user.Nome, user.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInsert, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso - ID: %d", idInsert)))

}

// This function is responsible for search for all users created on the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Connection()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco!"))
		return
	}
	defer db.Close()

	// SELECT * FROM USER
	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar informações de usuários no banco!"))
		return
	}
	defer linhas.Close()

	var users []user
	for linhas.Next() {
		var usuario user

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao scanear o usuário"))
			return
		}

		users = append(users, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}

}

// Search for a especif user
func SearchUser(w http.ResponseWriter, r *http.Request) {

}
