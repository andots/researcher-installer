package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	ShowWelcomeMessage()

	confirmed := ShowConfirmation()

	if (confirmed) {
		CreateAppDirectories()
		dir := GetSrcPath()
		// urls := MakeDownloadUrls()
		urls := []string{
			"https://raw.githubusercontent.com/uschindler/german-decompounder/master/dictionary-de.txt",
			"https://raw.githubusercontent.com/uschindler/german-decompounder/master/de_DR.xml",
		}
		d := NewDownloader(dir, urls)
		err := d.Start()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Cancel!")
		os.Exit(1)
	}

	BlockForWindows()
}
