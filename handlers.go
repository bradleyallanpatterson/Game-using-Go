package main

import (

	"fmt"
	"net/http"
	"io/ioutil"
	"math/rand"

)

type Client struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`

}

type Info struct {
	Key Client `json:"client"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
	
}

func Billups(w http.ResponseWriter, r *http.Request) {
    response, err := http.Get("https://codechallenge.boohma.com/choice")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
		fmt.Fprint(w, string(data) + "\n")
	}
}

func choice(w http.ResponseWriter, r *http.Request) {
	
	var randNumber = rand.Intn(4)+1
    fmt.Println(randNumber)
	
	if randNumber == 1 {
       fmt.Fprint(w, "{\"id\":1, \"name\": \"rock\"}")
    } else if randNumber == 2 {
       fmt.Fprint(w, "{\"id\":2, \"name\": \"paper\"}")
    } else if randNumber == 3 {
       fmt.Fprint(w, "{\"id\":3, \"name\": \"scissors\"}")
    } else if randNumber == 4 {
       fmt.Fprint(w, "{\"id\":4, \"name\": \"lizard\"}")
    } else {
       fmt.Fprint(w, "{\"id\":5, \"name\": \"spock\"}")
    }

	
}

func choices(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "[{\"id\":1, \"name\": \"rock\"}, {\"id\":2, \"name\": \"paper\"}, {\"id\":3, \"name\": \"scissors\"}, {\"id\":4, \"name\": \"lizard\"}, {\"id\":5, \"name\": \"spock\"} ]")
	
}