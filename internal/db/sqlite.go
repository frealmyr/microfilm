package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const sqlSchema = `
	CREATE TABLE IF NOT EXISTS image (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT,
		checksum TEXT,
		FOREIGN KEY (id) REFERENCES exif(id)
	);
	CREATE TABLE IF NOT EXISTS exif (
		id INTEGER NOT NULL PRIMARY KEY,
		make TEXT,
		model TEXT,
		lensMake TEXT,
		lensModel TEXT
	);
`

func Nuke() {
	os.Remove("./microfilm.sqlite")
}

func Generate() {
	db, err := sql.Open("sqlite3", "./microfilm.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := sqlSchema

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func Demo() {
	db, err := sql.Open("sqlite3", "./microfilm.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
		INSERT INTO image(
			path, checksum
		) VALUES (
			'/home/fredrick/Repositories/frealmyr/microfilm/assets/pictures/Wallpapers/wallhaven-d6pld3.jpg',
			'786d50481e9c29b09d0ced4c83ca67f0'
		);

		INSERT INTO exif(
			id, make, model, lensMake, lensModel
		) VALUES (
			(SELECT id FROM image WHERE checksum='786d50481e9c29b09d0ced4c83ca67f0'),
			'Sony',
			'A7c',
			'Tamron',
			'24-70mm 2.8'
		);

		INSERT INTO image(
			path, checksum
		) VALUES (
			'/home/fredrick/Repositories/frealmyr/microfilm/assets/pictures/Rome-2022/2022-08-19_173504.JPG',
			'05bd9c9c441f2fa82687f3fc66eabe38'
		);

		INSERT INTO exif(
			id, make, model, lensMake, lensModel
		) VALUES (
			(SELECT id FROM image WHERE checksum='05bd9c9c441f2fa82687f3fc66eabe38'),
			'Sony',
			'A7c',
			'Tamron',
			'24-70mm 2.8'
		);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func Sqlite() {
	Nuke()
	Generate()
	Demo()
}

func example() {
	db, err := sql.Open("sqlite3", "./foo.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちは世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
