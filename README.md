# Simple Rock Paper Scissors game clone using Go

This code repository hosts the examples used in the following [The New Stack](http://thenewstack.io/make-a-restful-json-api-go/) article.


## Setup

must type the following:

go get
go get -u github.com/tidwall/gjson


## Execution

To run the example game ...

``` The webpage 
# Run webpage
go run ./webPage/webpage.go
```

``` The Rest Service
# Run Rest Service
# note cannot get go run ./*.go to work
go run ./main.go ./router.go ./logger.go ./routes.go ./handlers.go
```

## Resources

Go Documentation - https://golang.org/pkg/
Serve web pages - https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
gjson - https://github.com/tidwall/gjson 

