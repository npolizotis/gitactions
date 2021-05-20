package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)


func mainHandler(w http.ResponseWriter,r* http.Request){
	w.Write([]byte("hello world"))
}

func main() {

	r:=chi.NewRouter()
	r.Get("/",mainHandler)
	http.ListenAndServe(":8080", r)
}}