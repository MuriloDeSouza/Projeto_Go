package main

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"

    _ "github.com/lib/pq"
)

type Produto struct {
    Nome       string
    Descricao  string
    Preco      float64
    Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
    http.HandleFunc("/", index)

    log.Println("Servidor rodando em http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Erro ao iniciar o servidor:", err)
    }
}

// Função para colocar os produtos na página inicial
func index(w http.ResponseWriter, r *http.Request) {
    log.Println("Acessaram a rota /")

    // Conectando ao banco de dados Supabase
	// Murik444@
    connStr := "postgresql://postgres:"senha"@db.snriwwglaqhtqdyucedt.supabase.co:5432/postgres"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Erro ao conectar ao banco de dados:", err)
    }
    defer db.Close()
    // Buscando produtos do banco de dados
    rows, err := db.Query("SELECT nome, descricao, preco, quantidade FROM produtos")
    if err != nil {
        log.Fatal("Erro ao buscar produtos:", err)
    }
    defer rows.Close()

    var produtos []Produto
    for rows.Next() {
        var p Produto
        if err := rows.Scan(&p.Nome, &p.Descricao, &p.Preco, &p.Quantidade); err != nil {
            log.Fatal("Erro ao escanear produto:", err)
        }
        produtos = append(produtos, p)
    }

    // Renderizando o template com os produtos
    err = temp.ExecuteTemplate(w, "Index", produtos)
    if err != nil {
        log.Println("Erro ao renderizar template:", err)
        http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
    }
}