package main

import (

	"fmt"
	"net/http"
	"io/ioutil"
	"math/rand"
	"encoding/json"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	
)

type RandomNumber struct {
	randomnumber     int   `json:"random_number"`

}

type CurrentResult struct {
	results     string   `json:"results"`
	player     	string   `json:"player"`
	computer    string   `json:"computer"`
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

func convertIntegerToAction(num int) (string) {

	if num < 0 || num > 100 {
        return "Invalid Input" 
    }
	
	if( num >= 80 ) {
		return "spock"
	} else if( num > 60 ) {
		return "lizard"
	} else if( num > 40 ) {
		return "scissors"
	} else if( num > 20 ) {
		return "paper"
	} else {
		return "rock"
	}

}


func getRandomNumberAsString(w http.ResponseWriter) (string, error) {
    response, err := http.Get("https://codechallenge.boohma.com/random")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    }
	
	data, _ := ioutil.ReadAll(response.Body)
	value := gjson.Get(string(data), "random_number")
	//println(value.String())

    return value.String(), err
}

func play(w http.ResponseWriter, r *http.Request) {

	var playerChoice string
	//fmt.Println( r.URL.Query() )
	for k, v := range r.URL.Query() {
	
		//fmt.Println( k )
		//fmt.Println( v )
		
		pc := strings.Join(v," ")
		playerChoice = pc
        fmt.Printf("%s: %s\n", k, playerChoice)
    }
	
	numForComputer, err := getRandomNumberAsString(w)
	if err != nil {
        panic(err.Error())
    }
	
	iNumForComputer, err := strconv.Atoi( numForComputer )
	fmt.Println( iNumForComputer )
	
	computerPick := convertIntegerToAction( iNumForComputer )
	fmt.Println( computerPick )
	
	var finalResult string
	if( playerChoice == computerPick ) {
		finalResult = "tie"
	} else if( playerChoice == "rock" ) {
		if( computerPick == "scissors" || computerPick == "lizard" ) {
			finalResult = "win"
		} else {
			finalResult = "lose"
		}
	} else if( playerChoice == "paper" ) {
		if( computerPick == "rock" || computerPick == "spock" ) {
			finalResult = "win"
		} else {
			finalResult = "lose"
		}
	} else if( playerChoice == "scissors" ) {
		if( computerPick == "paper" || computerPick == "lizard" ) {
			finalResult = "win"
		} else {
			finalResult = "lose"
		}
	} else if( playerChoice == "lizard" ) {
		if( computerPick == "spock" || computerPick == "paper" ) {
			finalResult = "win"
		} else {
			finalResult = "lose"
		}
	} else {
		if( computerPick == "scissors" || computerPick == "rock" ) {
			finalResult = "win"
		} else {
			finalResult = "lose"
		}
	}
	
	fmt.Println( finalResult )
	
	result := CurrentResult {
	
		results: finalResult,
		player: playerChoice,
		computer: computerPick,
		
	}
	resJson, _ := json.MarshalIndent(result, "", " ")
		
	fmt.Println( resJson )	
}








