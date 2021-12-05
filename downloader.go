package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

type Target struct {
	Url string
	Filename string
}

type Downloader struct {
	wg *sync.WaitGroup
	pool chan *Target
	MaxConcurrents int
	Client http.Client
	SaveDir string
	Targets []Target
}

func NewDownloader(saveDir string, urls []string) *Downloader {
	max := runtime.NumCPU()
	d := &Downloader{
		wg: &sync.WaitGroup{},
		SaveDir: saveDir,
		MaxConcurrents: max,
	}
	for _, url := range urls {
		d.add(url)
	}

	return d
}

func (d *Downloader) Start() error {
	fmt.Println("")
	fmt.Println("Start downloading....")
	fmt.Println("")
	d.pool = make(chan *Target, d.MaxConcurrents)
	pb := mpb.New(mpb.WithWaitGroup(d.wg), mpb.WithWidth(40))
	for _, target := range d.Targets {
		d.wg.Add(1)
		go d.download(target, pb)
	}
	pb.Wait()
	d.wg.Wait()

	fmt.Println("")
	fmt.Println("Successfully downloaded.")
	fmt.Println("")

	return nil
}

func (d *Downloader) add(url string) {
	filename := GetFileNameFromUrl(url)
	d.Targets = append(d.Targets, Target{
		Url: url,
		Filename: filename,
	})
}

func (d *Downloader) download(target Target, pb *mpb.Progress) error {
	defer d.wg.Done()
	d.pool <- &target

	filename := target.Filename
	filePath := filepath.Join(d.SaveDir, target.Filename)

	// ! Check file exists and skip downloading if exists
	if (FileExists(filePath)) {
		fmt.Printf("[SKIP]: Found at %v.\n", filePath)
	} else {
		file, err := os.Create(filePath)
		if err != nil {
			file.Close()
			return err
		}

		req, err := http.NewRequest(http.MethodGet, target.Url, nil)
		if err != nil {
			file.Close()
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			file.Close()
			return err
		}

		defer resp.Body.Close()

		fileSize, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
		bar := pb.AddBar(
			int64(fileSize),
			mpb.PrependDecorators(
				decor.Name(filename, decor.WCSyncSpaceR),
				// decor.Name(filename, decor.WC{W: len(filename) + 1, C: decor.DidentRight}),
				// decor.Spinner(nil, decor.WCSyncSpace),
			),
			mpb.AppendDecorators(
				decor.CountersKibiByte("% .2f / % .2f"),
				// decor.OnComplete(
				// 	decor.EwmaETA(decor.ET_STYLE_GO, 60, decor.WCSyncSpaceR), "done",
				// ),
				// decor.Percentage(decor.WCSyncSpace, decor.WCSyncSpaceR),
			),
		)
		reader := bar.ProxyReader(resp.Body)
		defer reader.Close()

		if _, err := io.Copy(file, reader); err != nil {
			file.Close()
			return err
		}

		file.Close()
	}

	<-d.pool

	return nil
}
