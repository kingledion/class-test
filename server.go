package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {

	log.Printf("Method %s\n", req.Method)

	for name, headers := range req.Header {
		log.Printf("Header %s: %s\n", name, headers)
	}

	log.Printf("Body: %s\n", req.Body)

	fmt.Fprintf(w, "Hello, this is Dan Hartig's machine\n")
}

type addReq struct {
	Var1 float64 `json:"var1"`
	Var2 float64 `json:"var2"`
}

func add(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		log.Printf("Header %s: %s\n", name, headers)
	}

	log.Printf("Body: %s\n", req.Body)

	var request addReq
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		fmt.Fprintf(w, "Unable to bind the request %s\n", req.Body)
		return
	}

	fmt.Fprintf(w, "%f + %f = %f\n", request.Var1, request.Var2, request.Var1+request.Var2)
}

func main() {

	fmt.Println("getenv: ", os.Getenv("TEST_DB_DSN"))

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/add", add)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Print(err)
	}
}
