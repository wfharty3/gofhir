package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type patient struct {
	Identifier string
	Active     int
	Name       string
	Telecom    string
}

func (p *patient) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(p.Identifier))
}

func encounter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("encounter handler called"))
}

func main() {
	data, _ := json.Marshal(&patient{"123", 1, "John Doe", "6947590"})
	fmt.Println(string(data))
	fmt.Println("")
	/*
		http.Handle("/patient", &patient{Message: "John Doe"})
		http.HandleFunc("/encounter", encounter)
		http.ListenAndServe(":5000", nil)
	*/
}
