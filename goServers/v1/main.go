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
type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

func main() {
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
}

// The JSON record sent from the HTTP client is converted into a Go structure
// using a call to json.Unmarshal()
func insertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
}
