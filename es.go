package main

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
)

const ES_BASE_URL = "https://artifacts.elastic.co/downloads/elasticsearch/"
const ES_NAME = "elasticsearch-7.10.1"

const DR_DICT_URL = "https://raw.githubusercontent.com/uschindler/german-decompounder/master/dictionary-de.txt"
const DR_XML_URL = "https://raw.githubusercontent.com/uschindler/german-decompounder/master/de_DR.xml"

const SUDACHI_URL = "http://sudachi.s3-website-ap-northeast-1.amazonaws.com/sudachidict/sudachi-dictionary-20210802-full.zip"
const SUDACHI_JSON_URL = "https://raw.githubusercontent.com/andots/researcher-docker/main/sudachi.json"
const SUDACHI_PLUGIN_URL = "https://github.com/WorksApplications/elasticsearch-sudachi/releases/download/v2.1.0/analysis-sudachi-7.10.1-2.1.0.zip"
const SUDACHI_ZIP_NAME = "sudachi-dictionary-20210802-full.zip"
const SUDACHI_DIR_NAME = "sudachi-dictionary-20210802"

func GetESFilename() (string, error) {
	// elasticsearch-7.10.1-linux-x86_64.tar.gz
	switch runtime.GOOS {
	case "windows":
		return fmt.Sprintf("%s-%s-x86_64.zip", ES_NAME, "windows"), nil
	case "darwin":
		return fmt.Sprintf("%s-%s-x86_64.tar.gz", ES_NAME, "darwin"), nil
	case "linux":
		return fmt.Sprintf("%s-%s-x86_64.tar.gz", ES_NAME, "linux"), nil
	default:
		return "", errors.New("Unsupported platform!")
	}
}

func GetESUrl() string {
	filename, err := GetESFilename()
	HandleError(err)
	return ES_BASE_URL + filename
}

func InstallPlugins() {
	cmd := filepath.Join(GetESPath(), "bin", "elasticsearch-plugin")
	_, err := exec.LookPath(cmd)
	HandleError(err)

	installPlugin(cmd, "https://github.com/WorksApplications/elasticsearch-sudachi/releases/download/v2.1.0/analysis-sudachi-7.10.1-2.1.0.zip")
	installPlugin(cmd, "analysis-nori")
	installPlugin(cmd, "analysis-smartcn")
}

func installPlugin(cmd string, plugin string) {
	fmt.Printf("[PLUGIN]: Installing %s ..... ", plugin)
	err := exec.Command(cmd, "install", plugin).Run()
	HandleError(err)
	fmt.Printf("Done!\n")
}
