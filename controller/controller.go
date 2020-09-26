package controller

import (
	"ClassroomManagements/models"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
)

var collection *mongo.Collection
var ctx = context.TODO()
var uri = "mongodb+srv://Sanhernandezmon:Shm1qazz@realmcluster.xofvc.mongodb.net/GraderDB?retryWrites=true&w=majority"

func createCourse(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	var courses models.Courses
	collection := client.Database("GraderDB").Collection("Courses")
	reqBody, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(reqBody, &courses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insertResult, err := collection.InsertOne(ctx, courses)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	w.WriteHeader(http.StatusCreated)
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	var courses []*models.Courses
	collection := client.Database("GraderDB").Collection("Courses")
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s models.Courses
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		courses = append(courses, &s)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(courses)
}
