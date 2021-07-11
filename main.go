package main

import (
	"log"
	"net/http"

	"github.com/jimmiepr/file-system/handlefunc"
)

func main() {

	http.HandleFunc("/dirfile", handlefunc.GetFileList)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
