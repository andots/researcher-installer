package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/codeclysm/extract/v3"
)

func ExtractFile(path string, to string) {
	filename := filepath.Base(path)
	fmt.Printf("[Extract]: %s ..... ", filename)
	data, err := ioutil.ReadFile(path)
	HandleError(err)

	buffer := bytes.NewBuffer(data)
	err = extract.Archive(context.Background(), buffer, to, nil)
	if err != nil {
		HandleError(err)
	}
	fmt.Print("Done!\n")
}
