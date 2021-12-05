package main

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/codeclysm/extract/v3"
)

func ExtractFile(path string, to string) {
	data, err := ioutil.ReadFile(path)
	HandleError(err)

	buffer := bytes.NewBuffer(data)
	err = extract.Archive(context.Background(), buffer, to, nil)
	HandleError(err)
}
