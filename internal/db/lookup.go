package db

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/frealmyr/microfilm/pkg/checksum"
)

type pathChecksum struct {
	filepath string
	checksum string
}

func Lookup(images []string) {
	var allRecords []pathChecksum

	for _, image := range images {
		item := pathChecksum{
			filepath: image,
			checksum: hex.EncodeToString(checksum.Md5(image)),
		}
		allRecords = append(allRecords, item)
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
