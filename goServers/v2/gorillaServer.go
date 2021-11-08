package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var DATAFILE = "/tmp/restful.json"

var DATA = []User{}

// Read JSON file from disk
func readDATAFILE() error {
	file, _ := os.Open(DATAFILE)
	defer file.Close()

	decoder := json.NewDecoder(file)
	record := User{}
	decoder.Token()
	for decoder.More() {
		decoder.Decode(&record)
		DATA = append(DATA, record)
	}
	return nil
}

// Write JSON Slice to disk
func writeDATAFILE() error {
	file, err := os.Create(DATAFILE)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	return SliceToJSON(DATA, file)
}

// SliceToJSON encodes a slice with JSON records
func SliceToJSON(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

func DefaultHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)
	rw.WriteHeader(http.StatusNotFound)
	Body := r.URL.Path + " is not supported. Thanks for visiting!\n"
	fmt.Fprintf(rw, "%s", Body)
}

// MethodNotAllowedHandler is executed when the method is incorrect
func MethodNotAllowedHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)
	rw.WriteHeader(http.StatusNotFound)
	Body := "Method not allowed!\n"
	fmt.Fprintf(rw, "%s", Body)
}

// TimeHandler is for handling /time
func TimeHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	rw.WriteHeader(http.StatusOK)
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is: " + t + "\n"
	fmt.Fprintf(rw, "%s", Body)
}

// InsertHandler is for adding a new user
func InsertHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = User{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(user)
	DATA = append(DATA, user)

	err = writeDATAFILE()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func RemoveElement(username string) bool {
	for i, v := range DATA {
		if v.Username == username {
			DATA = append(DATA[:i], DATA[i+1:]...)
			return true
		}
	}
	return false
}

// DeleteHandler is for deleting an existing user + DELETE
func DeleteHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	vars := mux.Vars(r)
	username := vars["username"]

	if len(username) == 0 {
		log.Println("username value not set!")
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// Remove multiple elements if they exist
	// As the DATA slice changes, we need to call it as a function
	for {
		if RemoveElement(username) == false {
			break
		}
	}

	err := writeDATAFILE()
	if err != nil {
		fmt.Println(err)
		return
	}

}

// ListHandler is for getting all data from the user database
func ListHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	err := SliceToJSON(DATA, rw)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

// PORT is where the web server listens to
var PORT = ":1234"

type notAllowedHandler struct{}

func (h notAllowedHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	MethodNotAllowedHandler(rw, r)
}

func main() {
	// Read DATAFILE
	err := readDATAFILE()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(DATA)

	// Create a new ServeMux using Gorilla
	mux := mux.NewRouter()

	s := http.Server{
		Addr:         PORT,
		Handler:      mux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
	mux.NotFoundHandler = http.HandlerFunc(DefaultHandler)

	notAllowed := notAllowedHandler{}
	mux.MethodNotAllowedHandler = notAllowed

	// Register GET
	getMux := mux.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/time", TimeHandler)
	getMux.HandleFunc("/list", ListHandler)

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

	log.Println("Listening to", PORT)
	err = s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
