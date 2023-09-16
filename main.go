package main

import (
	"askida-tatli/db"
	"askida-tatli/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)

	testDbOperations()

	router.Run("localhost:8080")
}

func testDbOperations() {
	// Replace the connection string with your own
	client, err := db.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	db := client.Database("test_db")

	// Create a new user
	user := model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	err = user.Create(context.Background(), db, "users", &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User created: %v\n", user)

	// Read a user by ID
	var readUser model.User
	err = readUser.Read(context.Background(), db, "users", bson.M{"_id": user.ID}, &readUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User read: %v\n", readUser)

	// Update a user's email
	update := bson.M{"$set": bson.M{"email": "john.doe_updated@example.com", "updated_at": primitive.NewDateTimeFromTime(user.UpdatedAt)}}
	err = user.Update(context.Background(), db, "users", bson.M{"_id": user.ID}, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User updated: %v\n", user)

	// Delete a user by ID
	//err = user.Delete(context.Background(), db, "users", bson.M{"_id": user.ID})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("User deleted")
}

func getUsers(context *gin.Context) {
	u := user{"5", "Murtaza"}
	context.IndentedJSON(http.StatusOK, u)
}
