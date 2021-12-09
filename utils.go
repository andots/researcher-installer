package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func MakeDownloadUrls() ([]string) {
	urls := []string {
		DR_DICT_URL,
		DR_XML_URL,
		SUDACHI_URL,
		SUDACHI_JSON_URL,
	}

	urls = append(urls, GetESUrl())

	return urls
}

// Return $HOME/researcher
func GetAppPath() string {
	homeDir, err := os.UserHomeDir()
	HandleError(err)
	dir := filepath.Join(homeDir, "researcher")
	return dir
}

// Return $HOME/researcher/src
func GetSrcPath() string {
	return filepath.Join(GetAppPath(), "src")
}

func GetESPath() string {
	return filepath.Join(GetAppPath(), ES_NAME)
}

// Check Elasticsearch already installed in $HOME/researcher/elasticsertch-x.x.x
func CheckInstalled() {
	p := GetESPath()
	if (FileExists(p)) {
		fmt.Printf("%s found.\nYou may have already installed Elasticsearch.\n\n", p)
		BlockForWindows()
		os.Exit(1)
	}
}

func CreateAppDirectories() {
	appPath := GetAppPath()
	CreateDir(appPath)
	srcPath := GetSrcPath()
	CreateDir(srcPath)
}

func CreateDir(path string) {
	fmt.Printf("[CREATE]: %s\n", path)
	err := os.MkdirAll(path, os.ModePerm)
	HandleError(err)
}

func MoveFile(from string, to string) {
	name := filepath.Base(from)
	fmt.Printf("[MOVE]:   %s to %s\n", name, to)
	err := os.Rename(from, to)
	HandleError(err)
}

func GetCurrentDir() string {
	currentPath, err := os.Getwd()
	HandleError(err)
	return currentPath
}

func RemoveAppDir() {
	dir := GetAppPath()
	err := os.RemoveAll(dir)
	HandleError(err)
}

func GetFileNameFromUrl(download_url string) string {
	url_struct, err := url.Parse(download_url)
	HandleError(err)
	return filepath.Base(url_struct.Path)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func BlockForWindows() {
	if runtime.GOOS == "windows" {
		for {
			fmt.Println("[Press `Ctrl+C` key or close window to exit...]")
			time.Sleep(10 * time.Second)
		}
	}
}

func HandleError(err error) {
	if err != nil {
		log.Println("")
		log.Println("[ERROR]", err)
		BlockForWindows()
		os.Exit(1)
	}
}
