package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ShowWelcomeMessage()

	CheckInstalled()

	confirmed := ShowConfirmation()

	if (confirmed) {
		fmt.Println("")
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

		// ! Create config dirs
		esPath := GetESPath()
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

		// mv src/sudachi.json to config/sudachi/sudachi.json
		MoveFile(
			filepath.Join(srcPath, "sudachi.json"),
			filepath.Join(sudachiConfigPath, "sudachi.json"),
		)

		// ! Install plugins
		InstallPlugins()

		ShowEndMessage()
	} else {
		fmt.Println("Cancel!")
		os.Exit(1)
	}

	BlockForWindows()
}
