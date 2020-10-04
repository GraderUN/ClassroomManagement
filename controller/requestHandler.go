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
	//test
	myRouter.HandleFunc("/", apiStatus)

	//basic crud
	myRouter.HandleFunc("/course", createCourse).Methods("POST")
	myRouter.HandleFunc("/course", GetAllCourses).Methods("GET")
	myRouter.HandleFunc("/classroom", createClassroom).Methods("POST")
	myRouter.HandleFunc("/classroom", GetAllClassrooms).Methods("GET")
	myRouter.HandleFunc("/assignations", AssignClassroom).Methods("POST")
	myRouter.HandleFunc("/assignations", GetAllAssignedCourses).Methods("GET")

	//complexrequest
	myRouter.HandleFunc("/assignations/course/{courseid}", GetAssignationsbycourse).Methods("GET")
	myRouter.HandleFunc("/assignations/classroom/{classroomid}", GetAssignationsbyclassroom).Methods("GET")
	myRouter.HandleFunc("/assignations/proffesor/{proffesorid}", GetAssignationsbyproffesor).Methods("GET")

	fmt.Println("Port 8080 is listening")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
