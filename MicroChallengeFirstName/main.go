package main

import (
	"fmt"
	"github.com/gin-gonic/gin" // gin is framework for RESTapi
	"net/http"
)

func getFirstName(context *gin.Context) { // must use context for this framework
	firstName := fmt.Sprintf(("Shachar"))
	fmt.Println("[Debug] Returning first name: %s", firstName)
	//We will return the string as JSON and send an http ok message. H will create the JSON
	context.JSON(http.StatusOK, gin.H{"FirstName": firstName})
}

func main(){

	fmt.Printf("Hello from first name") // print for debug
	router := gin.Default() // create a router
	router.GET("/firstName", getFirstName) // create a get method
	router.Run(":8082") // run the router

}
