package checksum

import (
	"crypto/md5"
	"io"
	"log"
	"os"
)

func Md5(filepath string) []byte {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	sha := h.Sum(nil)

	return sha
}
