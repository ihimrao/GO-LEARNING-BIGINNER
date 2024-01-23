package helpers

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func ErrorHandler(err error) {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No matching document found.")
			return
		}
		if err == fmt.Errorf("cannot delete user") {
			fmt.Println("cannot delete user")
			return
		}
		log.Fatal(err)
	}
}
