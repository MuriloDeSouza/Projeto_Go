package models

import (	
	"app/db"
	"projeto_go/app/db"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// GetAll busca todos os produtos no banco de dados
func GetAll() ([]Produto, error) {
	rows, err := db.DB.Query("SELECT nome, descricao, preco, quantidade FROM produtos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []Produto
	for rows.Next() {
		var p Produto
		if err := rows.Scan(&p.Nome, &p.Descricao, &p.Preco, &p.Quantidade); err != nil {
			return nil, err
		}
		produtos = append(produtos, p)
	}

	return produtos, nil
}