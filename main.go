package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Starting with XKCD
const (
	latestPostUrl = "https://xkcd.com/info.0.json"
)

// JSON Structure
type xkcdJSONStruct struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func main() {

	response, err := http.Get(latestPostUrl)
	if err != nil {
		log.Fatal("Error getting the request\n", err)
	}

	var latestXKCDPost xkcdJSONStruct
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response\n", err)
	}

	err = json.Unmarshal(data, &latestXKCDPost)
	if err != nil {
		log.Fatal("Error unmarshalling the response\n", err)
	}

	fmt.Println(latestXKCDPost)
}
