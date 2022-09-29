// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetWorkout(writer http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	//var id string
	//var record int
	id := queryParams.Get("_id")
	log.Printf("ID is %s", id)
	//} else {
	//	id = "6331d8f63212da02bf3419a6"
	//	log.Printf("ID is %s", id)
	//}
	//if queryParams.Has("record") {
	record, _ := strconv.Atoi(queryParams.Get("record"))
	log.Printf("record is %d", record)
	//} else {
	//	record = 76
	//	log.Printf("record is %d", record)
	//}

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
	err = coll.FindOne(context.TODO(), bson.D{{"record", record}}).Decode(&workout)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(workout)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the id %s\n", id)
		log.Panicln(err)
	}

	jsonData, err2 := json.MarshalIndent(workout, "", "    ")
	if err2 != nil {
		log.Panicln(err)
	}

	fmt.Fprint(writer, string(jsonData))
}
