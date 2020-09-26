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
	myRouter.HandleFunc("/create-user", createUserController).Methods("POST")
	myRouter.HandleFunc("/recover-password/{email}", recoverPasswordController).Methods("GET")
	myRouter.HandleFunc("/validate-auth-token/{token}", validateAuthTokenController).Methods("GET")
	myRouter.HandleFunc("/authentication/{email}/{password}", authenticationController).Methods("GET")
	myRouter.HandleFunc("/change-password", chagePasswordController).Methods("PUT")
	myRouter.HandleFunc("/assign-profile/{token}", assignProfileController).Methods("PUT")
	myRouter.HandleFunc("/verify-acount/{email}/{vcode}", verifyAcountController).Methods("PUT")

	fmt.Println("Port 4002 is listening")
	log.Fatal(http.ListenAndServe(":4002", myRouter))
}
