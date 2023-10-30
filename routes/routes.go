package routes

import (
	"github.com/allasassis/api-rest-go.git/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/allasassis/api-rest-go.git/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.ReturnPersonality).Methods("Get")
	r.HandleFunc("/personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/personalities/{id}", controllers.EditPersonality).Methods("Put")
	r.HandleFunc("/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
