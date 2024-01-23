package main

import (
	"GO-DB/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{id}", controller.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	fmt.Println("Server is starting...")
	http.ListenAndServe(":8000", r)
}
