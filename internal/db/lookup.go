package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/frealmyr/microfilm/internal/bootstrap"
)

type pathChecksum struct {
	filepath string
	checksum string
}

type JsonImage = bootstrap.JsonImage
type JsonExif = bootstrap.JsonExif

func Lookup(images []jsonImage) {
	// var allRecords []pathChecksum

	for _, image := range images {
		fmt.Println("Immateapot")
		spew.Dump(image)
	}
	// spew.Dump(allRecords)

	db, err := sql.Open("sqlite3", "./microfilm.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, path, checksum from image")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var path string
		var checksum string
		err = rows.Scan(&id, &path, &checksum)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, path, checksum)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
