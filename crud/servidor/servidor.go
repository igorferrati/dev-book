package servidor

import (
	"crud/database"
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

// CriarUsuario insere usuários no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	// Prepare statement
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter id inserido"))
		return
	}

	// Status code
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInserido)))

}

// BuscarUsuarios trás todos usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro buscar usuários"))
		return
	}
	defer linhas.Close()

	// slice de usuários
	var usuarios []usuario
	for linhas.Next() {
		// variável do tipo struct usuario
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear linhas usuário"))
			return
		}
		// popula slice
		usuarios = append(usuarios, usuario)

	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}

}

// BuscarUsuario um usuário específico salvos no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuário por ID"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear linha de usuário"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao convertr usuário para JSON"))
		return
	}
}
