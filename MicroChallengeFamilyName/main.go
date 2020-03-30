package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//declare the function we are getting from [address]/familyName

func getFamilyName(context *gin.Context){
	familyName := fmt.Sprintf(("Bartal"))
	fmt.Println("[Debug] Returning family name: %s", familyName)

	//We will return the string as JSON and send an http ok message. H will create the JSON
	context.JSON(http.StatusOK, gin.H{"FamilyName" : familyName})
}

func main(){
	fmt.Printf("Hello from family name ")
	router := gin.Default() //creating the remote
	router.GET("/familyName", getFamilyName) //declare - if we go to address/FamilyName - we will go to getFamilyName() function
	router.Run(":8081") //running our router on 8081 port
}

