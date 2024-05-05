package main
// Quando nós vamos usar um pacote mas de forma implicita, nao vai ser nesse arquivo exatamente que vai precisar dele nós importamos o pacote com _ antes
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// String de conexão para o banco de dados "usuario:senha@/bancodedados"
	stringConexao := "valve:valve@/valve?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil{
		log.Fatal(erro)
	}
	fmt.Println("A conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil{
		log.Fatal(erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)
}
