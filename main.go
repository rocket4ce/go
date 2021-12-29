package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Custumer struct {
	Name    string
	City    string
	ZipCode string
}

func main() {
	// rutas
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/custumers", getAllCustumers)
	// server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getAllCustumers(w http.ResponseWriter, r *http.Request) {
	custumers := []Custumer{
		{Name: "Dinko", City: "Coquimbo", ZipCode: "123123"},
		{Name: "Dinko2", City: "Coquimbo", ZipCode: "234563"},
	}

	json.NewEncoder(w).Encode(custumers)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hola mundo")
}
