package routes

import (
	"log"
	"net/http"

	"github.com/RenanLourenco/go-rest-api/controllers"
	"github.com/RenanLourenco/go-rest-api/middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.ListPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetOnePersonality).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("PUT")

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	
	log.Fatal(http.ListenAndServe(":8000",handlers.CORS(headers,methods,origins)(r)))
}
