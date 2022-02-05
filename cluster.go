package main

import (
	"context"
	"encoding/json"
	"everywheretew.it/main/common"
	"everywheretew.it/main/models"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type Cluster struct {
	Id     primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name   string             `bson:"name" json:"name"`
	UserId []int              `bson:"userId" json:"userId"`
	// stub data following
	Women       int
	Men         int
	NotDeclared int
	Facebook    float32
	Twitter     float32
	Referral    float32
	Products    []models.Product
}

func getCluster(w http.ResponseWriter, r *http.Request) {
	common.SetupCORS(&w, r)
	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	ClusterCollection := ConnectTo_Cluster()
	result := Cluster{}
	err := ClusterCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		panic(err)
	}

	//insert stub data
	result = inMemoryClusterDb[id]
	// end stubb
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}

	defer ClusterCollection.Database().Client().Disconnect(context.Background())
}

func getAllCluster(w http.ResponseWriter, r *http.Request) {
	common.SetupCORS(&w, r)
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

var inMemoryClusterDb = map[string]Cluster{
	"61d9c7f193247cd87083247c": {
		Name:        "test1",
		Women:       13,
		Men:         24,
		NotDeclared: 1,
		Facebook:    30,
		Twitter:     50,
		Referral:    20,
		Products:    models.InMemoryProductsDb["61d9c7f193247cd87083247c"],
	},
	"61d9d046b0556f25834f5c7a": {
		Name:        "test2",
		Women:       40,
		Men:         30,
		NotDeclared: 2,
		Facebook:    50,
		Twitter:     45,
		Referral:    5,
		Products:    models.InMemoryProductsDb["61d9d046b0556f25834f5c7a"],
	},
	"61d9d07b63564d1ec3798083": {
		Name:        "test3",
		Women:       25,
		Men:         35,
		NotDeclared: 1,
		Facebook:    16,
		Twitter:     16,
		Referral:    68,
		Products:    models.InMemoryProductsDb["61d9d07b63564d1ec3798083"],
	},
}
