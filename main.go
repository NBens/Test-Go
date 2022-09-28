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

var xkcdCacheMap map[int]xkcdJSONStruct

// This function makes a request to an XKCD post with the number "num", and if the number is set to -1, then it gets the latest post
func getXPost(num int) (xkcdJSONStruct, error) {
	var url string
	if num == -1 {
		url = latestPostUrl
	} else {
		url = fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)
	}
	response, err := http.Get(url)
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

// This function makes a request to XKCD's latest post API endpoint, and fetches the data as a xkcdJSONStruct element
func getLatestXPost() (xkcdJSONStruct, error) {
	latestPost, err := getXPost(-1)
	if err != nil {
		return xkcdJSONStruct{}, err
	}
	return latestPost, nil
}

// This function fetches the last 10 posts by XKCD, from the post with the number "startNum" until "startNum - endNum"
func getLatestXPosts(startNum, endNum int) ([]xkcdJSONStruct, error) {
	diff := startNum - endNum
	results := make([]xkcdJSONStruct, 0)
	for i := 0; i < diff+1; i++ {
		temp, err := getXPost(startNum - i)
		if err != nil {
			return nil, err
		}
		results = append(results, temp)
	}
	return results, nil
}

func main() {

	latestXKCDPosts, err := getLatestXPosts(2677, 2673) // Get the last 5 posts
	if err != nil {
		log.Fatal("Couldn't get the latest XKCD posts\n", err)
	}
	fmt.Println(latestXKCDPosts)
}
