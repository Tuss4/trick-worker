package main

import (
	"database/sql"
	"fmt"
)

var (
	QueryVideo = "SELECT video_id FROM videos_video WHERE video_id=$1;"
)

func updateDb(db *sql.DB, videos map[string]Video) {
	for _, v := range videos {
		var vid string
		err := db.QueryRow(QueryVideo, v.id).Scan(&vid)
		switch {
		case err == sql.ErrNoRows:
			fmt.Println("No video with that ID.")
		case err != nil:
			fmt.Println(err)
		default:
			println("Adding video to DB.")
		}
	}
}
