package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/jlmbaka/extconv/extconv"
)

func main() {
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	oldExt := os.Args[2]
	newExt := os.Args[3]
	extconv.ChangeExts(files, oldExt, newExt)
}
