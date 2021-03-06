package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GraderUN/ClassroomManagement/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	}
	insertResult, err := collection.InsertOne(ctx, courses)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single course: ", insertResult.InsertedID)
	fmt.Fprint(w, "Inserted a single course: ", insertResult.InsertedID)
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
	}
	insertResult, err := collection.InsertOne(ctx, classroom)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single classroom: ", insertResult.InsertedID)
	fmt.Fprint(w, "Inserted a single classroom: ", insertResult.InsertedID)
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
	}
	insertResult, err := collection.InsertOne(ctx, asignation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Asigned a single classroom to a single course: ", insertResult.InsertedID)
	fmt.Fprint(w, "Asigned a single classroom: ", asignation.Salon, " to a single"+
		"course: ", asignation.Curso, " ", insertResult.InsertedID)
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
	collection := client.Database("GraderDB").Collection("Classroom")
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {

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

func GetAssignationsbycourse(w http.ResponseWriter, r *http.Request) {
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
	vars := mux.Vars(r)
	callID := vars["courseid"]

	filter := bson.M{"curso": callID}
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

func GetAssignationsbyclassroom(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	vars := mux.Vars(r)
	callID := vars["classroomid"]

	filter := bson.M{"salon": callID}
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

func GetAssignationsbyproffesor(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	vars := mux.Vars(r)
	callID := vars["professorid"]

	filter := bson.M{"profesor": callID}
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

func DeleteAssignemet(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	vars := mux.Vars(r)
	callID := vars["classid"]

	idPrimitive, err := primitive.ObjectIDFromHex(callID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	} else {

	}
	filter := bson.M{"_id": idPrimitive}
	collection := client.Database("GraderDB").Collection("AssignedClassroom")
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if res.DeletedCount == 0 {
		fmt.Println("DeleteOne() document not found:", res)
	} else {
		// Print the results of the DeleteOne() method
		fmt.Println("DeleteOne Result:", res)
		fmt.Fprint(w, "successfully deleted")

	}
	w.WriteHeader(http.StatusAccepted)
}

func DeleteClassroom(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	vars := mux.Vars(r)
	callID := vars["classroomid"]

	idPrimitive, err := primitive.ObjectIDFromHex(callID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	} else {

	}
	filter := bson.M{"_id": idPrimitive}
	collection := client.Database("GraderDB").Collection("Classroom")
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if res.DeletedCount == 0 {
		fmt.Println("DeleteOne() document not found:", res)
	} else {
		// Print the results of the DeleteOne() method
		fmt.Println("DeleteOne Result:", res)
		fmt.Fprint(w, "successfully deleted")

	}
	w.WriteHeader(http.StatusAccepted)
}
