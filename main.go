package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &users)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
