package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	QueryVideo = "SELECT video_id FROM videos_video WHERE video_id=$1;"
)

type NewVideo struct {
	Type         string `json:"video_type"`
	Title        string `json:"title"`
	VideoID      string `json:"video_id"`
	ThumbnailURL string `json:"thumbnail_url"`
}

func postVideo(v Video) {
	d := NewVideo{YT, v.title, v.id, v.url}
	jd, err := json.Marshal(d)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest("POST", getApiUrl(), bytes.NewBuffer(jd))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", getToken())
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	println("Video:", v.id, "status:", resp.Status)
	defer resp.Body.Close()
}

func updateDb(db *sql.DB, videos map[string]Video) {
	for _, v := range videos {
		var vid string
		err := db.QueryRow(QueryVideo, v.id).Scan(&vid)
		switch {
		case err == sql.ErrNoRows:
			fmt.Println("No video with that ID.")
			fmt.Printf("Adding video %v to database.\n", v.id)
			postVideo(v)
		case err != nil:
			fmt.Println(err)
		}
	}
}
