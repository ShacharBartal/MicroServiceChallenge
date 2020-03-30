package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type person struct{
	FirstName string
	FamilyName string
}

func getNameToOurJson(client *http.Client, url string, p *person){
	result , err := infrastructure(client, url)
	if err != nil { // we didnt success the command above ^
		fmt.Printf("infrastructure error: %v\n", err)
	} else { // succeess
		fmt.Printf("raw response: %s\n", string(result))
		err := json.Unmarshal(result, &p) // convert result to p1
		if err != nil {
			fmt.Printf("unmarshal error: %v\n", err)
		} else {
			fmt.Printf("New name response: %+v\n", p)
		}

	}
}


func main(){
	fmt.Println("Hello from base")
	p1 := person{}
	client := &http.Client{} // define MicroChallengeBase as a client
	//will get familyName with following commands

	getNameToOurJson(client, "http://127.0.0.1:8081/familyName", &p1)
	getNameToOurJson(client, "http://127.0.0.1:8082/firstName", &p1)
}

// create http request
// send http request
// convert byte[] response to string
func infrastructure(client *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil) // define req as request to the GET function
	resp, err := client.Do(req) // do the request
	if err != nil {
		return nil, fmt.Errorf("client req error: " + err.Error())
	} else { // if there is an error, return nil and the error message
		responseBody, err := ioutil.ReadAll(resp.Body) // responseBody will get the body of "resp" - the output from the GET request.
		if err != nil {
			return nil, fmt.Errorf("body read error: " + err.Error())
		} else {
			return responseBody, err
		}
	}
}
