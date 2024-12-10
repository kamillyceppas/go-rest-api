package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kamillyceppas/go-rest-api/database"
	"github.com/kamillyceppas/go-rest-api/models"
)

// FUNÇÃO PARA EXIBIR A PAGINA HOME
func Home(w http.ResponseWriter, r *http.Request) { // receber requisão e responder
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // qnd receber requisao a resposta sera do tipo json
	var p []models.Personalidade // array de personalidades
	database.DB.Find(&p)         // FIND encontra todas personalidades
	json.NewEncoder(w).Encode(p)
}

// FUNÇÃO PARA RETORNAR APENAS A PERSONALIDADE ESCOLHIDA
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id) // FIRST retorna apenas a personalidade e o id escolhido
	json.NewEncoder(w).Encode(personalidade)
}

// caso nao encontre retorna tudo vazio

// FUNÇÃO PARA CRIAR UMA NOVA PERSONALIDADE
func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade) // recebe um json e quer decodificar
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade) // pega uma informacao para exibir encodando
}

// FUNÇÃO PARA DELETAR UMA PERSONALIDADE
func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

// FUNCÃO PARA EDITAR UMA PERSONALIDADE
func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade

	// Verifica se o registro existe
	result := database.DB.First(&personalidade, id)
	if result.Error != nil {
		http.Error(w, "Personalidade não encontrada", http.StatusNotFound)
		return
	}

	// Atualiza os campos específicos
	var dadosAtualizados models.Personalidade
	json.NewDecoder(r.Body).Decode(&dadosAtualizados)

	personalidade.Nome = dadosAtualizados.Nome
	personalidade.Historia = dadosAtualizados.Historia

	// Salva as alterações
	database.DB.Save(&personalidade)

	// Retorna a personalidade atualizada
	json.NewEncoder(w).Encode(personalidade)
}
