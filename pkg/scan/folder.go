package scan

import (
	"io/fs"
	"log"
	"path/filepath"
)

var pathList []string

func pathFinder(path string, d fs.DirEntry, err error) error {
	if err != nil {
		log.Fatal(err)
	}
	if !d.IsDir() {
		pathList = append(pathList, path)
	}
	return nil
}

func Folder(path string) ([]string, error) {
	err := filepath.WalkDir(path, pathFinder)
	if err != nil {
		return nil, err
	}
	return pathList, nil
}
