package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/allasassis/api-rest-go.git/database"
	"github.com/allasassis/api-rest-go.git/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the homepage, welcome!")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ReturnPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p []models.Personality
	database.DB.First(&p, id)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Panic(err)
	}
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	err := json.NewDecoder(r.Body).Decode(&personality)
	if err != nil {
		log.Panic(err)
	}
	database.DB.Create(&personality)
	json.NewEncoder(w).Encode(personality)

}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p []models.Personality
	database.DB.Delete(&p, id)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Panic(err)
	}
}
