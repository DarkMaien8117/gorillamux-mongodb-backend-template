package main

import (
	"fmt"
	"go-mongodb/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router() //get the instance of "router/router.go"
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r)) //Waiting request what has < r > at 8080 port
}
