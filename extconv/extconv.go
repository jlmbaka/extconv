package extconv

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ChangeExts(tgtDir string, files []os.FileInfo, matchExt, newExt string) {
	for _, file := range files {
		changeExt(tgtDir, file, matchExt, newExt)
	}
}

func changeExt(tgtDir string, file os.FileInfo, matchExt, newExt string) {
	tgtFilepath := path.Join(tgtDir, file.Name())
	myFileExt := filepath.Ext(tgtFilepath)
	if myFileExt == matchExt {
		renameFile(tgtFilepath, myFileExt, newExt)
	}
}

func renameFile(oldpath, oldExt, newExt string) {
	newpath := strings.TrimSuffix(oldpath, oldExt) + newExt
	err := os.Rename(oldpath, newpath)
	if err != nil {
		log.Fatal(err)
	}
}
