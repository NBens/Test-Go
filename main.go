package main

import (
	"encoding/json"
	"fmt"
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

//var xkcdCacheMap map[int]xkcdJSONStruct

func getLatestXPost() (xkcdJSONStruct, error) {
	response, err := http.Get(latestPostUrl)
	if err != nil {
		return xkcdJSONStruct{}, err
	}

	var latestXKCDPost xkcdJSONStruct

	err = json.NewDecoder(response.Body).Decode(&latestXKCDPost)
	if err != nil {
		return xkcdJSONStruct{}, nil
	}
	return latestXKCDPost, nil
}

func main() {

	latestXKCDPost, err := getLatestXPost()
	if err != nil {
		log.Fatal("Couldn't get the latest XKCD post\n", err)
	}
	fmt.Println(latestXKCDPost.Num)
}
