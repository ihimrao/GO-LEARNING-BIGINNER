package controller

import (
	"GO-DB/helpers"
	"GO-DB/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://ihimrao:sciJz7anvmj2vDWB@lms.e5ahgmo.mongodb.net/?retryWrites=true&w=majority"))
	collection = client.Database("LMS").Collection("USER")
}

func createUser(user model.User) {
	inserted, err := collection.InsertOne(context.Background(), user)
	helpers.ErrorHandler(err)
	fmt.Println("Inserted user: ", inserted.InsertedID)
}

func deleteUser(userId string) (bson.M, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	helpers.ErrorHandler(err)
	deleted, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	helpers.ErrorHandler(err)
	if deleted.DeletedCount == 1 {
		return bson.M{"_id": id, "deleted": true}, nil
	}
	return bson.M{"_id": id, "deleted": false}, fmt.Errorf("cannot delete user")
}

func updateUser(userId string) {
	id, err := primitive.ObjectIDFromHex(userId)
	helpers.ErrorHandler(err)
	updated, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"firstName": "Sarkar"}})
	helpers.ErrorHandler(err)
	fmt.Println("Updated user: ", updated.ModifiedCount)
}

func getAllUsers() []bson.M {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	helpers.ErrorHandler(err)
	var users []bson.M
	for cursor.Next(context.Background()) {
		var user bson.M
		err := cursor.Decode(&user)
		helpers.ErrorHandler(err)
		users = append(users, user)
	}
	defer cursor.Close(context.Background())
	return users
}

func getUser(id string) bson.M {
	userId, err := primitive.ObjectIDFromHex(id)
	helpers.ErrorHandler(err)
	var result bson.M
	err = collection.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&result)
	helpers.ErrorHandler(err)
	return result
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := getAllUsers()
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	users := getUser(params["id"])
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	helpers.ErrorHandler(err)
	createUser(user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	res, err := deleteUser(params["id"])
	fmt.Println("error: ", err)
	json.NewEncoder(w).Encode(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	updateUser(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
