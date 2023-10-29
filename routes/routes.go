package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/allasassis/api-rest-go.git/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.ReturnPersonality).Methods("Get")
	r.HandleFunc("/personalities", controllers.CreatePersonality).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", r))
}
