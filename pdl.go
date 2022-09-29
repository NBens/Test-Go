package main

import (
	"encoding/xml"
	"log"
	"net/http"
	"regexp"
	"time"
)

// Declaring PoorlyDrawnLines RSS Link
const (
	PDLRSSFeed = "https://feeds.feedburner.com/PoorlyDrawnLines"
)

var PDLElementList []PDLElement
var lastUpdatedHour int // This will keep track of the time of our last update of the RSS feed

// This function makes a request to PDL's RSS feed and unmarshalls the XML into our structs
func getRSSFeed() (Rss, error) {
	response, err := http.Get(PDLRSSFeed)
	if err != nil {
		return Rss{}, err
	}

	var RSSFeed Rss

	err = xml.NewDecoder(response.Body).Decode(&RSSFeed)
	if err != nil {
		return Rss{}, nil
	}
	return RSSFeed, nil
}

// This function updates the content of our global list of elements ([]PDLElement)
func updatePDLContent(elemList *[]PDLElement, updateHour *int) {
	// Regex code to match URLs (Encoded element contains the image URL)
	re := regexp.MustCompile(`(http|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)

	lastFeed, err := getRSSFeed()
	if err != nil {
		log.Fatal("Couldn't get the latest PDL RSS feed\n", err)
	}

	timeLayout := "Mon, 02 Jan 2006 15:04:05 +0000"
	for key, value := range lastFeed.Channel.Items {
		t, err := time.Parse(timeLayout, value.PubDate)

		if err != nil {
			log.Fatal("Couldn't parse the time\n", err)
		}

		(*elemList)[key].Title = value.Title
		(*elemList)[key].Url = value.Link
		(*elemList)[key].PublishingDate = t
		(*elemList)[key].PictureUrl = re.FindString(value.Encoded)
	}
	*updateHour = time.Now().Hour()
}

func init() {
	// Initialize the slice of PDL RSS elements
	PDLElementList = make([]PDLElement, 10)
	updatePDLContent(&PDLElementList, &lastUpdatedHour)
}
