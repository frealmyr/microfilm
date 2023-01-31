package exif

import (
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

type exifResult struct {
	timestamp time.Time
	make      string
	model     string
	lensMake  string
	lensModel string
}

func Decode(path string) (*exifResult, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	exif.RegisterParsers(mknote.All...)

	exifData, err := exif.Decode(file)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	datetime, _ := exifData.DateTime()

	exifCamMake, _ := exifData.Get(exif.Make)
	camMake := exifCamMake.String()

	exifCamModel, _ := exifData.Get(exif.Model)
	camModel := exifCamModel.String()

	exifLensMake, _ := exifData.Get(exif.LensMake)
	lensMake := exifLensMake.String()

	exifLensModel, _ := exifData.Get(exif.LensModel)
	lensModel := exifLensModel.String()

	GenExif := exifResult{
		timestamp: datetime,
		make:      camMake,
		model:     camModel,
		lensMake:  lensMake,
		lensModel: lensModel,
	}

	return &GenExif, nil
}
