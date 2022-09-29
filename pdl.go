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

// XML Structs for decode
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Content string   `xml:"content,attr"`
	Wfw     string   `xml:"wfw,attr"`
	Dc      string   `xml:"dc,attr"`
	Atom    string   `xml:"atom,attr"`
	Sy      string   `xml:"sy,attr"`
	Slash   string   `xml:"slash,attr"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Text            string `xml:",chardata"`
	Title           string `xml:"title"`
	Link            Link   `xml:"link"`
	Description     string `xml:"description"`
	LastBuildDate   string `xml:"lastBuildDate"`
	Language        string `xml:"language"`
	UpdatePeriod    string `xml:"updatePeriod"`
	UpdateFrequency string `xml:"updateFrequency"`
	Generator       string `xml:"generator"`
	Items           []Item `xml:"item"`
}

type Link struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Item struct {
	Text        string `xml:",chardata"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Creator     string `xml:"creator"`
	PubDate     string `xml:"pubDate"`
	Category    string `xml:"category"`
	Guid        Guid   `xml:"guid"`
	Description string `xml:"description"`
	Encoded     string `xml:"encoded"`
}

type Guid struct {
	Text        string `xml:",chardata"`
	IsPermaLink string `xml:"isPermaLink,attr"`
}

type PDLElement struct {
	Title          string
	Url            string
	PublishingDate time.Time
	PictureUrl     string
}

var PDLElementList []PDLElement
var lastUpdatedHour int

// This function makes a request to PDL's RSS feed
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
