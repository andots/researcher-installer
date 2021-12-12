package main

import (
	"testing"
)

func TestShowEndMessage(t *testing.T) {
	ShowEndMessage()
}

func TestGetFileNameFromUrl(t *testing.T) {
	u := "https://raw.githubusercontent.com/uschindler/german-decompounder/master/dictionary-de.txt"
	name := "dictionary-de.txt"
	got := GetFileNameFromUrl(u)
	if (got != name) {
		t.Errorf("Filename is not %s", name)
	}
	u = "https://raw.githubusercontent.com/uschindler/german-decompounder/master/de_DR.xml"
	name = "de_DR.xml"
	got = GetFileNameFromUrl(u)
	if (got != name) {
		t.Errorf("Filename is not %s", name)
	}
	u = "http://sudachi.s3-website-ap-northeast-1.amazonaws.com/sudachidict/sudachi-dictionary-20210802-full.zip"
	name = "sudachi-dictionary-20210802-full.zip"
	got = GetFileNameFromUrl(u)
	if (got != name) {
		t.Errorf("Filename is not %s", name)
	}
	u = "https://raw.githubusercontent.com/andots/researcher-docker/main/sudachi.json"
	name = "sudachi.json"
	got = GetFileNameFromUrl(u)
	if (got != name) {
		t.Errorf("Filename is not %s", name)
	}
	u = "https://github.com/WorksApplications/elasticsearch-sudachi/releases/download/v2.1.0/analysis-sudachi-7.10.1-2.1.0.zip"
	name = "analysis-sudachi-7.10.1-2.1.0.zip"
	got = GetFileNameFromUrl(u)
	if (got != name) {
		t.Errorf("Filename is not %s", name)
	}
}
