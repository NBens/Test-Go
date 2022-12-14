package main

import (
	"encoding/xml"
	"time"
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

// PDLElement struct is used to save the data of each element
type PDLElement struct {
	Title          string
	Url            string
	PublishingDate time.Time
	PictureUrl     string
}
