package util

import (
	"errors"
	"log"
	"os"
)

func GetFileModTime(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, errors.New("open file err!!!")
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return 0, err
	}
	return fi.ModTime().Unix(), nil
}
