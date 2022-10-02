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
	id := queryParams.Get("_id")
	log.Printf("ID is %s", id)
	date := queryParams.Get("wDate")
	log.Printf("wDate is %s", date)
	record, _ := strconv.Atoi(queryParams.Get("record"))
	log.Printf("record query is %d", record)

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

	// client, collection
	coll := client.Database("workouts").Collection("workouts")
	var workout Workout

	// by date
	if len(date) > 0 || date != "" {
		err = coll.FindOne(context.TODO(), bson.D{{"workout_date", date}}).Decode(&workout)
		if err != nil {
			log.Panicln(err)
		}
		log.Println(workout)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the date %s\n", date)
			log.Panicln(err)
		}

		jsonData, err2 := json.MarshalIndent(workout, "", "    ")
		if err2 != nil {
			log.Panicln(err)
		}

		fmt.Fprint(writer, string(jsonData))
	}
	// by Record
	if record > 0 || record != 0 {
		err = coll.FindOne(context.TODO(), bson.D{{"record", record}}).Decode(&workout)
		if err != nil {
			log.Panicln(err)
		}
		log.Println(workout)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the id %d\n", record)
			log.Panicln(err)
		}

		jsonData, err2 := json.MarshalIndent(workout, "", "    ")
		if err2 != nil {
			log.Panicln(err)
		}
		fmt.Fprint(writer, string(jsonData))
	}
}
