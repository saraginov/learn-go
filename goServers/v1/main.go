package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// the `json:"user"` and `json:"password"` are called structure tags.
// these tags help when working with JSON records that contain the user and
// password fields
// The main bug when converting JSON records into Go structs and vice versa is
// not making the Go struct fields public
type UserV1 struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

func mainV1() {
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	// Create a new router using `http.NewServeMux()`
	mux := http.NewServeMux()
	// define parameters of HTTP server using http.Server
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	// each mux.Handle() call associates an endpoint to a handler function
	// multiple endpoints can be served by the same handler function when
	// appropriate
	mux.Handle("/time", http.HandlerFunc(timeHandler))
	mux.Handle("/insert", http.HandlerFunc(insertHandler))
	mux.Handle("/list", http.HandlerFunc(listHandler))

	// todo [] I don't know where code below belongs yet
	// Register GET
	getMux := mux.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/time", TimeHandler)
	getMux.HandleFunc("/list", TimeHandler)

	// Register DELETE
	// Delete User
	deleteMux := mux.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/delete", DeleteHandler)
	deleteMux.HandleFunc("/delete/", DeleteHandler)
	deleteMux.HandleFunc("/delete/{username}", DeleteHandler)

	// Register POST
	// Add User + Login + Logout
	postMux := mux.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/insert", InsertHandler)
}

// The JSON record sent from the HTTP client is converted into a Go structure
// using a call to json.Unmarshal()
func insertHandlerV1(w http.ResponseWriter, r *http.Request) {
	// check if HTTP method used by client (r.Method) matches desired
	// http.MethodPost
	if r.Method != http.MethodPost {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	// read client data (r.Body) and umarshal it into a user structure
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	// Source code doesn't explain where &user is suppose to come from, maybe it's
	// a global variable for this file, hopefully they clarify
	// unmarshal data into a user structure.
	// If the value of the Username field isn't empty, we insert that structure
	// into the Data map
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
}

// ignores the data sent from the client, constructs the date&time str and the
// Body of the response and then sends it to the client using fmt.Fprintf
func timeHandlerV1(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:" + t + "\n"
	fmt.Fprintf(w, "%s", Body)
	// Fprintf formats according to a format specifier and writes to w.
	// It returns the number of bytes written and any write error encountered.
	// numBits, err := fmt.Fprintf(w, "%s", Body)
}

// I don't see this code anywhere other than when I downloaded the source file
func listHandlerV1(w http.ResponseWriter, r *http.Request) {

}
