package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cluster struct {
	Id     primitive.ObjectID `json:"id"`
	Name   string             `json:"name"`
	UserId []int              `json:"userId"`
}

func getCluster(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params["id"]
	ClusterCollection := ConnectTo_Cluster()
	defer ClusterCollection.Database().Client().Disconnect(context.Background())
	result, err := ClusterCollection.Find(context.Background(), bson.M{"id": id})
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func getAllCluster(w http.ResponseWriter, r *http.Request) {
	ClusterCollection := ConnectTo_Cluster()
	defer ClusterCollection.Database().Client().Disconnect(context.Background())

	res := []Cluster{}
	result, err := ClusterCollection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}

	curErr := result.All(context.TODO(), &res)
	defer result.Close(context.Background())
	if curErr != nil {
		panic(curErr)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
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
}
