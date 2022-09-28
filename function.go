// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
)

func GetWorkout(writer http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	id := queryParams.Get("_id")

	uri := os.Getenv("DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("workouts").Collection("workouts")
	var workout Workout
	err = coll.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&workout)
	if err != nil {
		return
	}
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", "title")
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(workout, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Fprint(writer, string(jsonData))
}
