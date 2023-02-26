package bootstrap

import (
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/frealmyr/microfilm/pkg/checksum"
	"github.com/frealmyr/microfilm/pkg/exif"
)

type jsonImage struct {
	filename string `json:"filename"`
	basepath string `json:"basepath"`
	album    string `json:"album"`
	checksum string `json:"checksum"`
	exif     any    `json:"exif"`
}

func jsonGenerate(path string) jsonImage {

	album := ""
	if filepath.Dir(path) != os.Getenv("WATCH_DIR") {
		album = filepath.Base(filepath.Dir(path))
	}

	exif, err := exif.Decode(path)
	if err != nil {
		exif = nil
	}

	resultFileInfo := jsonImage{
		filename: filepath.Base(path),
		basepath: filepath.Dir(path),
		album:    album,
		checksum: hex.EncodeToString((checksum.Md5(path))),
		exif:     exif,
	}

	return resultFileInfo
}

func GenerateJson(images []string) {
	var allRecords []jsonImage

	for _, image := range images {
		allRecords = append(allRecords, jsonGenerate(image))
	}
	spew.Dump(allRecords)
}
