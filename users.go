package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type user struct {
	name  string
	id    primitive.ObjectID
	email string
}
