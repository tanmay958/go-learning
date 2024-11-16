package main

import (
	"fmt"
	"net/http"

	"github.com/tanmay958/dbapi/router"
)

func main() {
	fmt.Println("MONGO API" )
	r := router.Router() 
	fmt.Println("server is starting... ")
	http.ListenAndServe(":4000", r)
	fmt.Println("Listening")
}