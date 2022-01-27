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

type Cluster struct {
	Id     primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name   string             `bson:"name" json:"name"`
	UserId []int              `bson:"userId" json:"userId"`
}

func getCluster(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	ClusterCollection := ConnectTo_Cluster()
	result := Cluster{}
	err := ClusterCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}

	defer ClusterCollection.Database().Client().Disconnect(context.Background())
}

func getAllCluster(w http.ResponseWriter, r *http.Request) {
	ClusterCollection := ConnectTo_Cluster()

	res := []Cluster{}
	result, err := ClusterCollection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}

	curErr := result.All(context.Background(), &res)
	if curErr != nil {
		panic(curErr)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}

	defer ClusterCollection.Database().Client().Disconnect(context.Background())
	defer result.Close(context.Background())
}

func insertCluster(c Cluster) {
	ClusterCollection := ConnectTo_Cluster()
	result, err := ClusterCollection.InsertOne(context.Background(), bson.D{
		{"name", c.Name},
		{"userId", c.UserId},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	defer ClusterCollection.Database().Client().Disconnect(context.Background())
}
