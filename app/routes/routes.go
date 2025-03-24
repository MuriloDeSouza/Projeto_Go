package routes

import (
	"app/models"
	"html/template"
	"log"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos, err := models.GetAll()
	if err != nil {
		log.Println("Erro ao buscar produtos:", err)
		http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
		return
	}

	err = temp.ExecuteTemplate(w, "Index", produtos)
	if err != nil {
		log.Println("Erro ao renderizar template:", err)
		http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
	}
}