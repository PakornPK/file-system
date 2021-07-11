package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GetList(pathFolder string) []string {
	files, err := ioutil.ReadDir(pathFolder)
	var fileList []string
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileList = append(fileList, f.Name())
	}
	return fileList
}

func GetBirthTime(pathFile string) string {
	fi, err := os.Stat(pathFile)
	if err != nil {
		return err.Error()
	}
	mTime := fi.ModTime()

	return fmt.Sprintf("%v", mTime)
}
