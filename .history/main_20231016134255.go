package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func (w http.ResponseWriter, r *http.Request)   {
		json.N
		
	})

	log.Println("Server running on port 9000")
	http.ListenAndServe(":9000", nil)
}
