package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age,omitempty"`
}

var users = []User{}

// copied utility function
func generateRandomString(length int) (string, error) {
	// Calculate the number of bytes needed for the given length
	numBytes := (length * 3) / 4
	// Generate random bytes
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	// Encode random bytes to base64
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	// Trim the padding '=' characters and return the result
	return randomString[:length], nil
}
func main() {
	users = append(users, User{"1", "Himanshu", "Yadav", 22})
	users = append(users, User{"2", "Aman", "Yadav", 18})

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hey there, Server is working fine!"))
	}).Methods("GET")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		userData, _ := json.Marshal(users)
		w.Header().Add("Content-Type", "application/json")
		w.Write(userData)
	}).Methods("GET")

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		url := mux.Vars(r)
		userId := url["id"]
		var user User
		found := false
		for _, u := range users {
			if u.Id == userId {
				user = u
				found = true
				break
			}
		}
		if !found {
			w.Write([]byte("NO FOUND"))
			return
		}
		userByte, _ := json.Marshal(user)
		w.Header().Add("Content-Type", "application/json")
		w.Write(userByte)
	}).Methods("GET")

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		id, _ := generateRandomString(10)
		user := User{Id: id}
		json.Unmarshal(body, &user)
		users = append(users, user)
		buffedUser, _ := json.Marshal(user)
		w.Write(buffedUser)
	}).Methods("POST")

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		urlP := mux.Vars(r)
		id := urlP["id"]
		deleted := false
		for index, user := range users {
			if user.Id == id {
				users = append(users[:index], users[index+1:]...)
				deleted = true
			}
		}
		if !deleted {
			w.Write([]byte("Could not found respective user"))
			return
		}
		w.Write([]byte("Deleted Successfully"))
	}).Methods("DELETE")

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {

		body, _ := io.ReadAll(r.Body)
		var userData User
		json.Unmarshal(body, &userData)
		urlP := mux.Vars(r)
		id := urlP["id"]
		for index, user := range users{
		if(id == user.Id){

		}
		}
	}).Methods("PATCH")

	http.ListenAndServe(":8000", router)
}
