package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {

	router := NewRouter()

	fmt.Println("Starting the application...")
	log.Fatal(http.ListenAndServe(":8888", router))
	
	fmt.Println("Closing the application...")

}