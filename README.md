# Simple Rock Paper Scissors game clone using Go

This code repository hosts my first look at using the Go programming language.


## Setup

Assuming that Go is installed from https://golang.org/
Must type the following from command line:

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

