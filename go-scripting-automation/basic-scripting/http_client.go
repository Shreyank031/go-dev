package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	baseUrl := "https://httpbin.org"

	//Http GET request
	getReq, err := http.Get(baseUrl)
	if err != nil {
		fmt.Printf("Error on get request:", err)
		return
	}
	defer getReq.Body.Close()
	getBody, _ := io.ReadAll(getReq.Body)
	fmt.Println("Get Reqponse:", string(getBody))

	//Post request
	postBody := []byte(`{"key" : "value"}`)
	postReq, err := http.Post(baseUrl+"/post", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalf("Error in Post request:", err)
	}
	defer postReq.Body.Close()
	postBodyReq, err := io.ReadAll(postReq.Body)
	if err != nil {
		fmt.Printf("Error while reading the post request:", err)
	}
	fmt.Println("Post Request:", string(postBodyReq))

}
