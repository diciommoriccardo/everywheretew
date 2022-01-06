package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToCluster() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	uri := "mongodb+srv://cluster-everywheretew.4ulev.mongodb.net/myFirstDatabase?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=bin/mongoCert.pem"
	clientOpts := options.Client().ApplyURI(uri)
	client, ConnectionError := mongo.Connect(ctx, clientOpts)
	err := client.Ping(ctx, readpref.Primary())

	if ConnectionError != nil {
		fmt.Println("ERROR")
		log.Fatal(err)
		log.Fatal(ConnectionError)
	}

	defer client.Disconnect(ctx)
	twitterCollection := client.Database("api").Collection("twitter-data")
	//clusterCollection := client.Database("everywheretew").Collection("cluster")
	//userCollection := client.Database("everywheretew").Collection("users")
	docCount, CollectionErr := twitterCollection.CountDocuments(ctx, bson.D{})

	if CollectionErr != nil {
		fmt.Println("ERROR")
		log.Fatal(CollectionErr)
	}

	fmt.Println(docCount)
}
