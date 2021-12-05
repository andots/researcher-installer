package main

import (
	"errors"
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
	}

	switch runtime.GOOS {
	case "windows":
		urls = append(urls, ES_WIN_URL)
	case "darwin":
		urls = append(urls, ES_MAC_URL)
	case "linux":
		urls = append(urls, ES_LINUX_URL)
	default:
		HandleError(errors.New("Elasticsearch supports Windows, Mac, and Linux."))
	}

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
	appdir := GetAppPath()
	dir := filepath.Join(appdir, "src")
	return dir
}

func CreateAppDirectories() {
	appPath := GetAppPath()
	CreateDir(appPath)
	srcPath := GetSrcPath()
	CreateDir(srcPath)
}

func CreateDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	HandleError(err)
}

// Create app directory $HOME/researcher
// func CreateAppDir() string {
// 	dir := GetAppPath()
// 	err := os.MkdirAll(dir, os.ModePerm)
// 	HandleError(err)
// 	return dir
// }

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
		log.Println("err:", err)
		BlockForWindows()
		os.Exit(1)
	}
}
