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
		"https://raw.githubusercontent.com/uschindler/german-decompounder/master/dictionary-de.txt",
		"https://raw.githubusercontent.com/uschindler/german-decompounder/master/de_DR.xml",
		"http://sudachi.s3-website-ap-northeast-1.amazonaws.com/sudachidict/sudachi-dictionary-20210802-full.zip",
	}

	esWinUrl := "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.10.1-windows-x86_64.zip"
	esMacUrl := "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.10.1-darwin-x86_64.tar.gz"
	esLinuxUrl := "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.10.1-linux-x86_64.tar.gz"

	switch runtime.GOOS {
	case "windows":
		urls = append(urls, esWinUrl)
	case "darwin":
		urls = append(urls, esMacUrl)
	case "linux":
		urls = append(urls, esLinuxUrl)
	default:
		HandleError(errors.New("Elasticsearch supports Windows, Mac, and Linux."))
	}

	return urls
}

// Return $HOME/researcher
func GetAppDirPath() string {
	homeDir, err := os.UserHomeDir()
	HandleError(err)
	dir := filepath.Join(homeDir, "researcher")
	return dir
}

// Create app directory $HOME/researcher
func CreateAppDir() string {
	dir := GetAppDirPath()
	err := os.MkdirAll(dir, os.ModePerm)
	HandleError(err)
	return dir
}

func GetCurrentDir() string {
	currentPath, err := os.Getwd()
	HandleError(err)
	return currentPath
}

func RemoveAppDir() {
	dir := GetAppDirPath()
	err := os.RemoveAll(dir)
	HandleError(err)
}

func GetFileName(download_url string) string {
	url_struct, err := url.Parse(download_url)
	HandleError(err)
	return filepath.Base(url_struct.Path)
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
