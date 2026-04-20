package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	url := "https://httpbin.org/post"

	payload := map[string] string{"name": "kenzo tenma", "age": "67", "occupation": "unemployed"}

	//Convert map/struct to JSON bytes

	jsonValues, err := json.Marshal(payload)

	if err!= nil{
		log.Fatal(err)
	}

	//send a post request

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValues))
	
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("%v", resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}