package bootstrap

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"time"

	"github.com/frealmyr/microfilm/pkg/checksum"
	"github.com/frealmyr/microfilm/pkg/exif"
)

type JsonImages interface {
	Structure() []JsonImage
}

type JsonImage struct {
	Filename string `json:"filename"`
	Basepath string `json:"basepath"`
	Album    string `json:"album"`
	Checksum string `json:"checksum"`
	Exif     any    `json:"exif"`
}

type JsonExif struct {
	Timestamp time.Time `json:"timestamp"`
	Make      string    `json:"make"`
	Model     string    `json:"model"`
	LensMake  string    `json:"lensMake"`
	LensModel string    `json:"lensModel"`
}

func jsonGenerate(path string) JsonImage {

	album := ""
	if filepath.Dir(path) != os.Getenv("WATCH_DIR") {
		album = filepath.Base(filepath.Dir(path))
	}

	exif, err := exif.Decode(path)
	if err != nil {
		exif = nil
	}

	resultFileInfo := JsonImage{
		Filename: filepath.Base(path),
		Basepath: filepath.Dir(path),
		Album:    album,
		Checksum: hex.EncodeToString((checksum.Md5(path))),
		Exif:     exif,
	}

	return resultFileInfo
}

func GenerateJson(images []string) *[]JsonImage {
	var allRecords []JsonImage

	for _, image := range images {
		allRecords = append(allRecords, jsonGenerate(image))
	}

	// log.Println(&allRecords)
	// spew.Dump(allRecords)

	return &allRecords
}
