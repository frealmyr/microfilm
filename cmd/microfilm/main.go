package main

import (
	"log"
	"os"

	"github.com/frealmyr/microfilm/internal/db"
	"github.com/frealmyr/microfilm/pkg/scan"
)

func main() {
	images, err := scan.Folder(os.Getenv("WATCH_DIR"))
	if err != nil {
		log.Fatal(err)
	}

	// bootstrap.GenerateJson(images)
	// db.Sqlite()
	db.Lookup(images)
}
