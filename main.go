package main

import (
	"fmt"
	"os"
)

func main() {
	ShowWelcomeMessage()

	confirmed := ShowConfirmation()

	if (confirmed) {
		CreateAppDirectories()
		srcPath := GetSrcPath()
		urls := MakeDownloadUrls()
		// urls := []string{
		// 	"https://raw.githubusercontent.com/uschindler/german-decompounder/master/dictionary-de.txt",
		// 	"https://raw.githubusercontent.com/uschindler/german-decompounder/master/de_DR.xml",
		// 	"https://github.com/codeclysm/extract/raw/master/testdata/archive.zip",
		// 	"https://github.com/codeclysm/extract/raw/master/testdata/archive.tar.gz",
		// }
		d := NewDownloader(srcPath, urls)
		err := d.Start()
		HandleError(err)

		// ! Extract Elasticsearch
		// to := GetAppPath()
		// ExtractFile(filepath.Join(srcPath, "archive.zip"), to)
		// fmt.Printf("Extracted: %v\n", "archive.zip")

	} else {
		fmt.Println("Cancel!")
		os.Exit(1)
	}

	BlockForWindows()
}
