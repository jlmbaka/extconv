package extconv

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ChangeExts(files []os.FileInfo, matchExt, newExt string) {
	for _, file := range files {
		myFilePath := path.Join(os.Args[1], file.Name())
		myFileExt := filepath.Ext(myFilePath)
		if myFileExt == matchExt {
			renameFile(myFilePath, myFileExt, newExt)
		}
	}
}

func renameFile(oldpath, oldExt, newExt string) {
	newpath := strings.TrimSuffix(oldpath, oldExt) + newExt
	err := os.Rename(oldpath, newpath)
	if err != nil {
		log.Fatal(err)
	}
}
