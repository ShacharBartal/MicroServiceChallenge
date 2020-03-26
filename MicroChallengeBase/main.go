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


func main(){
	fmt.Println("Hello from base")

	p1 := person{}

	client := &http.Client{} // define MicroChallengeBase as a client

	//will get familyName with following command:
	familyNameMicroServiceResponse , err := infrastructure(client, "http://127.0.0.1:8081/familyName")

	if err != nil { // we didnt success the command above ^
		fmt.Printf("infrastructure error: %v\n", err)
	} else { // successed
		fmt.Printf("raw response: %s\n", string(familyNameMicroServiceResponse))
		err := json.Unmarshal(familyNameMicroServiceResponse, &p1) // convert familyNameMicroServiceResponse to p1
		if err != nil {
			fmt.Printf("unmarshal error: %v\n", err)
		} else {
			fmt.Printf("Family name response: %+v\n", p1)
		}

	}

}

// create http request
// send http request
// convert byte[] response to string
func infrastructure(client *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client req error: " + err.Error())
	} else {
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("body read error: " + err.Error())
		} else {
			return responseBody, err
		}
	}
}
