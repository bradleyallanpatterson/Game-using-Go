package main

import (
	"log"
	"net/http"

)

func main() {



	fmt.Println("Starting the application...")
	http.Handle("/", http.FileServer(http.Dir("./webPage")))
	
	log.Fatal(http.ListenAndServe(":8880", nil))
	

}