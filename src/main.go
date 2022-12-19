package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func main() {
	// print("hello :D")
	// print("i am backend developer now")

	// api.Dev()
	ExampleDecode()
}

func ExampleDecode() {

	base := "/home/fredrick/Repositories/frealmyr/microfilm"
	relf := "pictures/Rome-2022/2022-08-19_115208.JPG"

	fname := filepath.Join(base, relf)

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	camMake, _ := x.Get(exif.Make) // normally, don't ignore errors!
	fmt.Println(camMake.StringVal())

	camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
	fmt.Println(camModel.StringVal())

	lensModel, _ := x.Get(exif.LensModel) // normally, don't ignore errors!
	fmt.Println(lensModel.StringVal())

	// exposureTime, _ := x.Get(exif.ExposureTime) // normally, don't ignore errors!
	// exNumer, denom, _ := exposureTime.Rat2(0) // retrieve first (only) rat. value
	// fmt.Printf("%v/%v", exNumer, denom)

	apertue, _ := x.Get(exif.ApertureValue) // normally, don't ignore errors!
	exNumer, denom, _ := apertue.Rat2(0) // retrieve first (only) rat. value
	fmt.Printf("%v/%v", exNumer, denom)

	// focal, _ := x.Get(exif.FocalLength)
	// numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
	// fmt.Printf("%v/%v", numer, denom)

	// // Two convenience functions exist for date/time taken and GPS coords:
	// tm, _ := x.DateTime()
	// fmt.Println("Taken: ", tm)

	// lat, long, _ := x.LatLong()
	// fmt.Println("lat, long: ", lat, ", ", long)
}
