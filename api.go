package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

/*func apiInit() {
	config := oauth1.NewConfig("i6W7sfCBUV0WOp13PHqMyL5yf", "U8xjm197eif9KQBp9jEo1XC7bdgXwEzEYrY8NhjMW0NrsfhLNU")
	token := oauth1.NewToken("AAAAAAAAAAAAAAAAAAAAAFMEXwEAAAAASyhdewcw31ECxV7cAZTIq16BYPY%3DTY66IW8MlG1XtxAAu8QYA66otPrx1LXBPmAy4AlqjKHsL6Na7a", "U8xjm197eif9KQBp9jEo1XC7bdgXwEzEYrY8NhjMW0NrsfhLNU")
	httpClient := config.Client(oauth1.NoContext, token)
}*/

func insert(data []byte) {
	ApiCollection := ConnectToApiCluster()
	var bdoc []interface{}
	err := bson.UnmarshalExtJSON([]byte(data), true, &bdoc)
	if err != nil {
		panic(err)
	}
	result, err := ApiCollection.InsertMany(context.Background(), bdoc)

	fmt.Println(result)

	if err != nil {
		panic(err)
	}
}
