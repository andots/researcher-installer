package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ShowWelcomeMessage()

	confirmed := ShowConfirmation()

	if (confirmed) {
		CreateAppDirectories()
		srcPath := GetSrcPath()

		// ! Download files
		urls := MakeDownloadUrls()
		d := NewDownloader(srcPath, urls)
		err := d.Start()
		HandleError(err)

		// ! Extract sudachi dict to src dir
		ExtractFile(filepath.Join(srcPath, SUDACHI_ZIP_NAME), srcPath)

		// ! Extract Elasticsearch to app dir
		esFilename, err := GetESFilename()
		HandleError(err)
		ExtractFile(filepath.Join(srcPath, esFilename), GetAppPath())

	} else {
		fmt.Println("Cancel!")
		os.Exit(1)
	}

	BlockForWindows()
}

// urls := []string{
// 	"https://raw.githubusercontent.com/uschindler/german-decompounder/master/dictionary-de.txt",
// 	"https://raw.githubusercontent.com/uschindler/german-decompounder/master/de_DR.xml",
// 	"https://github.com/codeclysm/extract/raw/master/testdata/archive.zip",
// 	"https://github.com/codeclysm/extract/raw/master/testdata/archive.tar.gz",
// }
