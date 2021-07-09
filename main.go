package main

import (
	"fmt"

	"github.com/jimmiepr/file-system/file"
)

func main() {
	Fl := file.GetList("./mockup")

	for i := 0; i < len(Fl); i++ {
		fmt.Printf("File : %s\n", Fl[i])
	}
}
