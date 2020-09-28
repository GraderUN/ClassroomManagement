package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GraderUN/ClassroomManagement/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var collection *mongo.Collection
var ctx = context.TODO()
var uri = os.Getenv("DB_CONN")

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

	fmt.Println("Inserted a single course: ", insertResult.InsertedID)
	w.WriteHeader(http.StatusCreated)
}

func createClassroom(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	var classroom models.Classroom
	collection := client.Database("GraderDB").Collection("Classroom")
	reqBody, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(reqBody, &classroom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insertResult, err := collection.InsertOne(ctx, classroom)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single classroom: ", insertResult.InsertedID)
	w.WriteHeader(http.StatusCreated)
}

func AssignClassroom(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	var asignation models.AssignedClassroom
	collection := client.Database("GraderDB").Collection("AssignedClassroom")
	reqBody, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(reqBody, &asignation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insertResult, err := collection.InsertOne(ctx, asignation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Asigned a single classroom to a single course: ", insertResult.InsertedID)
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

func GetAllClassrooms(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	var classrooms []*models.Classroom
	collection := client.Database("GraderDB").Collection("Clasroom")
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s models.Classroom
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		classrooms = append(classrooms, &s)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(classrooms)
}

func GetAllAssignedCourses(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	var assignations []*models.AssignedClassroom
	collection := client.Database("GraderDB").Collection("AssignedClassroom")
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s models.AssignedClassroom
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		assignations = append(assignations, &s)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(assignations)
}

func GetAllAssignedCoursesof(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	//define filter
	var id models.IDType
	reqBody, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(reqBody, &id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.M{"curso": id.Key}
	var assignations []*models.AssignedClassroom
	collection := client.Database("GraderDB").Collection("AssignedClassroom")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s models.AssignedClassroom
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		assignations = append(assignations, &s)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(assignations)
}

func GetAllAssignedClassroomof(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	//define filter
	var id models.IDType
	reqBody, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(reqBody, &id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.M{"salon": id.Key}
	var assignations []*models.AssignedClassroom
	collection := client.Database("GraderDB").Collection("AssignedClassroom")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s models.AssignedClassroom
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		assignations = append(assignations, &s)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(assignations)
}
