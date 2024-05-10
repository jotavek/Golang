package servidor

import (
	"crud/banco"
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

// CriarUsuario é um handlehttp que insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// Lê todo corpo da requisição
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	// Converte os dados JSON do corpo da requisição para a estrutura de usuário
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
	}
	// Enviando o usuario para o banco de dados
	// Primeiro conectamos ao banco com a função que criamos banco.Conectar
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com banco de dados"))
		return
	}
	defer db.Close()

	//Para enviar o usuario pelo terminal do mysql seria dessa forma: INSERT INTO USUARIOS (nome, email) values ("nome", "email")

	// Iremos usar PREPARE STATEMENT para evitar ataques.
	// Prepara uma instrução SQL para inserir um novo usuário no banco de dados
	// Utilizamos os ? em values para evitar ataques do tipo SQL INJECTION
	statement, erro := db.Prepare("insert into usuarios(nome, email) values(?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()
	// Executa a instrução SQL para inserir o novo usuário no banco de dados
	// Essa função irá substituir os ? em values
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	// Obtém o ID do usuário inserido
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	// STATUS CODE

	// Define o status de resposta como 201 (Created) e retorna uma mensagem com o ID do usuário inserido
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))

}

// BuscarUsuarios traz todos os usuarios salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
	}
	defer db.Close()

	// SELECT * FROM USUARIOS
	// Quando fazemos consulta de dados utilizamos o Query que vai retornar linhas da tabela para gente

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuarios"))
	}
	defer linhas.Close()

	var usuarios []usuario
	// Pra cada linha que for retornado ele vai executar uma interação pra gente
	// linhas.Next() é utilizado para iterar sobre todas as linhas retornadas pela consulta ao banco de dados
	for linhas.Next() {
		var usuario usuario

		// linhas.Scan() é utilizado para escanear os valores das colunas dessa linha para as variáveis correspondentes na estrutura usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)

	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON"))
		return
	}

}

// BuscarUsuario traz um usuários específico salvo no banco dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	// strconv.ParseUint converte uma string em inteiro, o segundo parametro é a base do numero no caso está em base decimal e o terceiro parâmetro é o tamanho dos bits, uint32, uint64..
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, erro := db.Query("SELECT * FROM USUARIOS WHERE ID = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro buscar o usuário!"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuário!"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário para JSON!"))
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro para inteiro!"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	// O Statement é usado para qualquer operação que não seja de consulta.

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario remove um usuário dentro do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
