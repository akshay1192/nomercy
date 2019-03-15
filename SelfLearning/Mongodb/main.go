package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/exp/errors/fmt"
	"time"
)

type doc struct {
	Name  string `bson:"name"`
	Value int    `bson:"value"`
}

func main() {

	// creating context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// creating mongo client and then connecting the client to server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("error in connecting to mongo : ", err)
		return
	}
	defer client.Disconnect(ctx)

	// checking the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("error in pinging to mongo : ", err)
		return
	}

	// creates db and collection if not present and returns handle to collection
	collection := client.Database("testing").Collection("numbers")
	if collection == nil {
		fmt.Println("error in getting collection handle")
		return
	}

	// inserts doc into mongo
	if result, err := collection.InsertOne(ctx, bson.M{"name": "gup", "value": 3}); err != nil {
		fmt.Println("error in inserting doc : ", err)
		return
	} else {
		// mongo document id upon insertion
		fmt.Println("Object id from single insertion : ", result.InsertedID, "\n")
	}

	// inserting multiple docs
	if results, err := collection.InsertMany(ctx, []interface{}{bson.M{"name": "pi", "value": 3}, bson.M{"name": "aks", "value": 5}}); err != nil {
		fmt.Println("error in inserting doc : ", err)
		return
	} else {
		// mongo document id upon insertion
		fmt.Println("Object ids from multiple insertions : ", results.InsertedIDs, "\n")
	}

	// reading docs
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result doc
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Result in find loop : ", result)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("")

	// reading single doc
	docObj := doc{}
	collection.FindOne(ctx, bson.M{"name": "aks"}).Decode(&docObj)

	fmt.Println("Result from findOne : ", docObj.Name, " ", docObj.Value, "\n")

	// deletes doc from mongo
	if res, err := collection.DeleteOne(ctx, bson.M{"name": "pi"}); err != nil {
		fmt.Println("error in removing doc : ", err)
		return
	} else {
		fmt.Println("Delete one count : ", res.DeletedCount, "\n")
	}

	// deletes multiple doc from mongo
	if res, err := collection.DeleteMany(ctx, bson.M{}); err != nil {
		fmt.Println("error in removing doc : ", err)
		return
	} else {
		fmt.Println("Delete many count : ", res.DeletedCount, "\n")
	}

}
