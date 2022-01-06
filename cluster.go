package main

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type cluster struct {
	id     primitive.ObjectID
	name   string
	userid []int
}

func getCluster(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Cluster")
}
