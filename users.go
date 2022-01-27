package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name  string
	Id    primitive.ObjectID
	Email string
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	UserCollection := ConnectToUserCluster()
	result := User{}
	err := UserCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}

	defer UserCollection.Database().Client().Disconnect(context.Background())
}

func AddUser(user User) {
	UserCollection := ConnectToUserCluster()
	result, err := UserCollection.InsertOne(context.Background(), bson.D{
		{"name", user.Name},
		{"email", user.Email},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	defer UserCollection.Database().Client().Disconnect(context.Background())
}
