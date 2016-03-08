package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jlmbaka/extconv/extconv"
)

func main() {
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	oldExt := formatExt(os.Args[2])
	newExt := formatExt(os.Args[3])
	extconv.ChangeExts(files, oldExt, newExt)
}

func formatExt(ext string) string {
	if !strings.HasPrefix(ext, ".") {
		return "." + ext
	}
	return ext
}
