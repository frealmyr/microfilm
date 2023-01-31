package scan

import (
	"encoding/hex"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/frealmyr/microfilm/pkg/checksum"
	"github.com/frealmyr/microfilm/pkg/exif"
)

func walk(path string, d fs.DirEntry, err error) error {
	if err != nil {
		log.Fatal(err)
	}

	if !d.IsDir() {
		dir := filepath.Dir(path)
		parent := filepath.Base(dir)
		file := filepath.Base(path)

		println()
		println("File: " + file)
		println("Album: " + parent)
		println("Checksum: " + hex.EncodeToString((checksum.Md5(path))))

		exif, err := exif.Decode(path)
		if err != nil {
			// log.Fatal(err)
			fmt.Printf("info: exif no data")
		}
		fmt.Printf("%+v\n", exif) // TODO: Output to SQL statement
	}
	return nil
}

func Folder() {
	fmt.Println(filepath.WalkDir(os.Getenv("WATCH_DIR"), walk))
}
