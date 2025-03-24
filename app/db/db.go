package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// DB é a variável global de conexão com o banco de dados
var DB *sql.DB

// Init inicializa a conexão com o banco de dados
// Murik444@
func Init() {
	connStr := "postgres://postgres:Murik444@@snriwwglaqhtqdyucedt.supabase.co:5432/postgres?sslmode=require"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao pingar o banco de dados:", err)
	}

	log.Println("Conectado ao banco de dados com sucesso!")
}

// Close fecha a conexão com o banco de dados
func Close() {
	DB.Close()
}