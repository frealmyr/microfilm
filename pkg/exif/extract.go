package exif

import (
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

type exifResult struct {
	timestamp time.Time `json:"timestamp"`
	make      string    `json:"make"`
	model     string    `json:"model"`
	lensMake  string    `json:"lensMake"`
	lensModel string    `json:"lensModel"`
}

func Decode(path string) (*exifResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	exif.RegisterParsers(mknote.All...)

	exifData, err := exif.Decode(file)
	if err != nil {
		if exif.IsCriticalError(err) {
			return nil, err
		}
	}
	defer file.Close()

	datetime, err := exifData.DateTime()
	if err != nil {
		return nil, err
	}

	exifCamMake, err := exifData.Get(exif.Make)
	if err != nil {
		return nil, err
	}
	camMake := exifCamMake.String()

	exifCamModel, err := exifData.Get(exif.Model)
	if err != nil {
		return nil, err
	}
	camModel := exifCamModel.String()

	exifLensMake, err := exifData.Get(exif.LensMake)
	if err != nil {
		return nil, err
	}
	lensMake := exifLensMake.String()

	exifLensModel, err := exifData.Get(exif.LensModel)
	if err != nil {
		return nil, err
	}
	lensModel := exifLensModel.String()

	// Generate JSON
	json := exifResult{
		timestamp: datetime,
		make:      camMake,
		model:     camModel,
		lensMake:  lensMake,
		lensModel: lensModel,
	}

	return &json, nil
}
