package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/frealmyr/microfilm/internal/bootstrap"
	"github.com/frealmyr/microfilm/internal/db"
	"github.com/frealmyr/microfilm/pkg/scan"
)

func main() {
	cronJob()
}

func cronJob() {
	images, err := scan.Folder(os.Getenv("WATCH_DIR"))
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(images) // DEBUG

	yugeJson := bootstrap.GenerateJson(images)

	spew.Dump(yugeJson)

	db.Lookup(yugeJson)
}
