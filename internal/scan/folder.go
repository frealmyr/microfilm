package scan

import (
	"encoding/hex"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/frealmyr/microfilm/pkg/checksum"
	"github.com/frealmyr/microfilm/pkg/exif"
)

type fileInfo struct {
	filename string
	basepath string
	album    string
	checksum string
	exif     any
}

func walk(path string, d fs.DirEntry, err error) error {
	if err != nil {
		log.Fatal(err)
	}

	if !d.IsDir() {
		basepath := filepath.Dir(path)
		parent := filepath.Base(basepath)
		file := filepath.Base(path)

		exif, err := exif.Decode(path)
		if err != nil {
			exif = nil
		}

		resultFileInfo := fileInfo{
			filename: file,
			basepath: basepath,
			album:    parent,
			checksum: hex.EncodeToString((checksum.Md5(path))),
			exif:     exif,
		}

		spew.Dump(resultFileInfo)
	}
	return nil
}

func Folder() {
	filepath.WalkDir(os.Getenv("WATCH_DIR"), walk)
}
