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
		// ExtractFile(filepath.Join(srcPath, SUDACHI_ZIP_NAME), srcPath)

		// ! Extract Elasticsearch to app dir
		// esFilename, err := GetESFilename()
		// HandleError(err)
		// ExtractFile(filepath.Join(srcPath, esFilename), GetAppPath())

		// ! Create config dirs
		esPath := filepath.Join(GetAppPath(), ES_NAME)
		sudachiConfigPath := filepath.Join(esPath, "config", "sudachi")
		deConfigPath := filepath.Join(esPath, "config", "analysis", "de")
		CreateDir(sudachiConfigPath)
		CreateDir(deConfigPath)

		// ! Move files
		// mv de_DR.xml config/analysis/de
		deXMLName := "de_DR.xml"
		MoveFile(
			filepath.Join(srcPath, deXMLName),
			filepath.Join(deConfigPath, deXMLName),
		)

		// mv dictionary-de.txt config/analysis/de
		deDictName := "dictionary-de.txt"
		MoveFile(
			filepath.Join(srcPath, deDictName),
			filepath.Join(deConfigPath, deDictName),
		)

		// mv sudachi-dictionary-20210802/system_full.dic config/sudachi/system_core.dic
		MoveFile(
			filepath.Join(srcPath, SUDACHI_DIR_NAME, "system_full.dic"),
			filepath.Join(sudachiConfigPath, "system_core.dic"),
		)

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
