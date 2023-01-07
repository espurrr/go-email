package main

import (
	"encoding/json"
	"fmt"
	"goemailclient/mailer"
	"goemailclient/models"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func postMail(w http.ResponseWriter, r *http.Request) {

	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)

	// unmarshal this into a new Mail struct
	var mail models.Mail
	json.Unmarshal(reqBody, &mail)

	// send email
	mailer.SendMail(mail)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/sendMail", postMail).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {

	handleRequests()

}
