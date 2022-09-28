package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
)

// Starting with XKCD
const (
	latestXKCDPostUrl = "https://xkcd.com/info.0.json"
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
		url = latestXKCDPostUrl
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

// This function fetches the last n posts by XKCD, from the post with the number "startNum" until "startNum - endNum"
func getLatestXPosts(startNum, endNum int) (map[int]xkcdJSONStruct, error) {
	diff := startNum - endNum
	results := make(map[int]xkcdJSONStruct, 0)
	for i := 0; i < diff+1; i++ {
		temp, err := getXPost(startNum - i)
		if err != nil {
			return nil, err
		}
		results[temp.Num] = temp
	}
	return results, nil
}

func isXPostInCache(num int) bool {
	_, ok := xkcdCacheMap[num]
	return ok
}

func max(cache map[int]xkcdJSONStruct) (maxNumber int) {
	maxNumber = math.MinInt32
	for n := range cache {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

func main() {
	xkcdCacheMap = make(map[int]xkcdJSONStruct)
	lastPost, err := getLatestXPost()
	if err != nil {
		log.Fatal("Couldn't get the latest XKCD post\n", err)
	}

	// Check if the lates post is already cached, if it is then the other elements will also be already loaded
	if isXPostInCache(lastPost.Num) {
		fmt.Println(xkcdCacheMap) // If it is then print it for now
	} else {
		// If not, check the missing number posts (latest.Num - largest number in the map)
		// And fetch them and add them to the map while also deleting the unused ones

		currentMaxKey := 9
		if len(xkcdCacheMap) > 0 {
			currentMaxKey = max(xkcdCacheMap)
		}
		latestMissingPosts, err := getLatestXPosts(lastPost.Num-1, lastPost.Num-currentMaxKey)
		if err != nil {
			log.Fatal("Couldn't get the latest XKCD posts\n", err)
		}
		for key, value := range latestMissingPosts {
			xkcdCacheMap[key] = value
		}
		// Then add the latest post which we've already fetched in line 105
		xkcdCacheMap[lastPost.Num] = lastPost

		keys := make([]int, 0)
		for k := range xkcdCacheMap {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		keys = keys[:len(keys)-10] // Keys to remove from the map
		for _, element := range keys {
			delete(xkcdCacheMap, element)
		}
		fmt.Println(xkcdCacheMap)
	}

}
