package servidor

import (
	"crud-basico/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar converter o corpo da requisicao"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar converter o corpo da requisicao"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar conectar ao banco de dados"))
		return
	}
	defer db.Close()

	//prepare statement
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar obter o id inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar conectar ao banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar executar o statement"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Ocorreu um erro ao tentar escanear o usuario"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro = json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar converter o usuario para json"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar converter o id"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar conectar ao banco de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar executar o statement"))
		return
	}
	defer linha.Close()

	var usuario usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Ocorreu um erro ao tentar escanear o usuario"))
			return
		}
	}

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario não encontrado"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro = json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar converter o usuario para json"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar converter o id"))
		return
	}

	//corpo da requisicao
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar ler o corpo da requisicao"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar converter o json para usuario"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar executar o statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar converter o id"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		w.Write([]byte("Ocorreu um erro ao tentar executar o statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
