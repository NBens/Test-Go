package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"
)

type Data struct {
	Title          string
	Url            string
	PublishingDate time.Time
	PictureUrl     string
}

func consolidateData() []Data {

	data := []Data{}
	// Update the XKCD Cache Map
	updateXContent(&xkcdCacheMap)

	// Since the RSS data is updated hourly, we need to update our PDL elements list if the hour changes
	// This can be replaced later with a background job that updates the data in a Redis db
	if time.Now().Hour() != lastUpdatedHour {
		updatePDLContent(&PDLElementList, &lastUpdatedHour)
	}

	// Saving PDL list to the data element
	for _, value := range PDLElementList {
		data = append(data, Data(value))
	}

	timeLayout := "2-1-2006"
	for _, value := range xkcdCacheMap {
		t, err := time.Parse(timeLayout, (value.Day + "-" + value.Month + "-" + value.Year))
		if err != nil {
			log.Fatal("Couldn't parse the time\n", err)
		}
		data = append(data, Data{
			Title:          value.Title,
			Url:            value.Link,
			PictureUrl:     value.Img,
			PublishingDate: t,
		})
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].PublishingDate.After(data[j].PublishingDate)
	})

	return data
}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(404)
			return
		}
		data := consolidateData()
		jsonStr, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(500)
			log.Print("Couldn't parse JSON\n", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonStr)
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Could not serve at port 8080\n", err)
	}
}
