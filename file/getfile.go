package file

import (
	"io/ioutil"
	"log"
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
