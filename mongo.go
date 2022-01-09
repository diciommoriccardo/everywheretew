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

func ConnectToApiCluster() *mongo.Collection {
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

	ApiCollection := client.Database("api").Collection("twitter-data")
	docCount, CollectionErr := ApiCollection.CountDocuments(ctx, bson.D{})

	if CollectionErr != nil {
		fmt.Println("ERROR")
		log.Fatal(CollectionErr)
	}

	fmt.Println(docCount)
	return ApiCollection
}

func ConnectTo_Cluster() *mongo.Collection {
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

	ClusterCollection := client.Database("everywheretew").Collection("cluster")
	docCount, CollectionErr := ClusterCollection.CountDocuments(ctx, bson.D{})

	if CollectionErr != nil {
		fmt.Println("ERROR")
		log.Fatal(CollectionErr)
	}

	fmt.Println(docCount)
	return ClusterCollection
}

func ConnectToUserCluster() *mongo.Collection {
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

	UserCollection := client.Database("everywheretew").Collection("users")
	docCount, CollectionErr := UserCollection.CountDocuments(ctx, bson.D{})

	if CollectionErr != nil {
		fmt.Println("ERROR")
		log.Fatal(CollectionErr)
	}

	fmt.Println(docCount)
	return UserCollection
}
