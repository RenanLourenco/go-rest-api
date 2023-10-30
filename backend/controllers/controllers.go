package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RenanLourenco/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func ListPersonalities(w http.ResponseWriter, r *http.Request){
	personalities := models.Find()

	json.NewEncoder(w).Encode(personalities)
}

func GetOnePersonality(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	personality := models.FindOne(id)

	json.NewEncoder(w).Encode(personality)

}

func CreatePersonality(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)

	personality := models.Create(newPersonality)

	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deletedPersonality := models.Delete(id)

	json.NewEncoder(w).Encode(deletedPersonality)

}

func EditPersonality(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality 
	json.NewDecoder(r.Body).Decode(&personality)
	updatedPersonality := models.Edit(id, personality)
	json.NewEncoder(w).Encode(updatedPersonality)

}
