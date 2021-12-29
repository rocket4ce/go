package main

import (
	"log"
	"net/http"
)

func Start() {
	// rutas
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/custumers", getAllCustumers)
	// server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
