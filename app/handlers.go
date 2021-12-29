package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Custumer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func getAllCustumers(w http.ResponseWriter, r *http.Request) {
	custumers := []Custumer{
		{Name: "Dinko", City: "Coquimbo", ZipCode: "123123"},
		{Name: "Dinko2", City: "Coquimbo", ZipCode: "234563"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(custumers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(custumers)

	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hola mundo")
}
