package main

import (
	"fmt"
	"github.com/gin-gonic/gin" // gin is framework for RESTapi
)

var firstName = "Gal"

func getFirstName(context *gin.Context){ // must use context for this framework

}

func main(){

	fmt.Printf("Hello from first name")
	router := gin.Default()
	router.GET("/getFirstName", getFirstName)
	router.Run(":8014")

}
