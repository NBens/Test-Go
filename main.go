package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println(xkcdCacheMap)

	fmt.Println(PDLElementList)

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(404)
			return
		}
		jsonStr, err := json.Marshal(xkcdCacheMap)
		if err != nil {
			w.WriteHeader(500)
			log.Print("Couldn't parse JSON\n", err)
			return
		}

		w.WriteHeader(200)
		w.Write(jsonStr)
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Could not serve at port 8080\n", err)
	}
}
