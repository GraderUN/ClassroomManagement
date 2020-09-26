package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func apiStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Classroom management ms working")
}

// HandleRequest ..
func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", apiStatus)
	myRouter.HandleFunc("/create-course", createCourse).Methods("POST")
	myRouter.HandleFunc("/courses", GetAllCourses).Methods("GET")
	fmt.Println("Port 8080 is listening")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
