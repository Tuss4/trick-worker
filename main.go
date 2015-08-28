package main

import (
	"flag"
	"fmt"
	googtransport "google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"
)

var (
	maxResults = flag.Int64("max-results", 25, "Max YouTube Results")
	apiKey     = os.Getenv("GOOGLE_API_KEY")
)

const (
	query = "Martial Arts Tricking"
)

func printIDs(sectionName string, matches map[string]Video) {
	fmt.Printf("%v:\n", sectionName)
	for k, v := range matches {
		fmt.Printf("[%v] %v, %v\n", k, v.title, v.url)
	}
	fmt.Printf("\n\n")
}

func main() {
	flag.Parse()

	client := &http.Client{
		Transport: &googtransport.APIKey{Key: apiKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	call := service.Search.List("id,snippet").Q(query).MaxResults(*maxResults).Order("date")
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	videos := make(map[string]Video)

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = Video{
				item.Id.VideoId, item.Snippet.Title, item.Snippet.Thumbnails.High.Url}
		}
	}
	printIDs("Videos", videos)
}
