package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jlmbaka/extconv/extconv"
)

func main() {
	tgtDir := os.Args[1]
	files, err := ioutil.ReadDir(tgtDir)
	if err != nil {
		log.Fatal(err)
	}
	oldExt := formatExt(os.Args[2])
	newExt := formatExt(os.Args[3])
	extconv.ChangeExts(os.Args[1], files, oldExt, newExt)
}

func formatExt(ext string) string {
	if !strings.HasPrefix(ext, ".") {
		return "." + ext
	}
	return ext
}
