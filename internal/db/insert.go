package db

import (
	"github.com/davecgh/go-spew/spew"
)

type jsonImage struct {
	filename string `json:"filename"`
	basepath string `json:"basepath"`
	album    string `json:"album"`
	checksum string `json:"checksum"`
	exif     any    `json:"exif"`
}

func Insert(images map[string]interface{}) {

	for _, image := range images {
		spew.Dump(image)
	}

	// db, err := sql.Open("sqlite3", "./microfilm.sqlite")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// sqlStmt := `
	// 	INSERT INTO image(
	// 		path, checksum
	// 	) VALUES (
	// 		'/home/fredrick/Repositories/frealmyr/microfilm/assets/pictures/Wallpapers/wallhaven-d6pld3.jpg',
	// 		'786d50481e9c29b09d0ced4c83ca67f0'
	// 	);

	// 	INSERT INTO exif(
	// 		id, make, model, lensMake, lensModel
	// 	) VALUES (
	// 		(SELECT id FROM image WHERE checksum='786d50481e9c29b09d0ced4c83ca67f0'),
	// 		'Sony',
	// 		'A7c',
	// 		'Tamron',
	// 		'24-70mm 2.8'
	// 	);

	// 	INSERT INTO image(
	// 		path, checksum
	// 	) VALUES (
	// 		'/home/fredrick/Repositories/frealmyr/microfilm/assets/pictures/Rome-2022/2022-08-19_173504.JPG',
	// 		'05bd9c9c441f2fa82687f3fc66eabe38'
	// 	);

	// 	INSERT INTO exif(
	// 		id, make, model, lensMake, lensModel
	// 	) VALUES (
	// 		(SELECT id FROM image WHERE checksum='05bd9c9c441f2fa82687f3fc66eabe38'),
	// 		'Sony',
	// 		'A7c',
	// 		'Tamron',
	// 		'24-70mm 2.8'
	// 	);
	// `
	// _, err = db.Exec(sqlStmt)
	// if err != nil {
	// 	log.Printf("%q: %s\n", err, sqlStmt)
	// 	return
	// }
}
