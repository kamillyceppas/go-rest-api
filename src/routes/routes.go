package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kamillyceppas/go-rest-api/src/controllers"
	"github.com/kamillyceppas/go-rest-api/src/middleware"
)

// FUNÇÃO PARA MANIPULAR AS REQUISICOES DE API
func HandleResquest() {
	r := mux.NewRouter() // Instancia o roteador do Gorilla Mux
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")         // metodo de criacao
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")    // metodo para deletar
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")           // metodo para editar
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))) // qualquer aplicacao pode consumir a api                                                      // Usa o roteador do mux para lidar com as requisições
}
