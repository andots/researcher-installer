package main

import (
	"bufio"
	"context"
	"os"

	"github.com/codeclysm/extract/v3"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func ExtractFile(path string, to string) {
	ctx := context.Background()

	// data, err := ioutil.ReadFile(path)
	// HandleError(err)

	// fi, err := os.Stat(path)
	// HandleError(err)

	file, err := os.Open(path)
	if err != nil {
		file.Close()
		HandleError(err)
	}

	fi, err := file.Stat()
	if err != nil {
		file.Close()
		HandleError(err)
	}

	pb := mpb.NewWithContext(ctx, mpb.WithWidth(60))
	bar := pb.AddBar(
		fi.Size(),
		mpb.PrependDecorators(
			decor.Name("[Extract]", decor.WCSyncSpaceR),
			// decor.Name(filename, decor.WC{W: len(filename) + 1, C: decor.DidentRight}),
			// decor.Spinner(nil, decor.WCSyncSpace),
		),
		// mpb.AppendDecorators(
			// decor.CountersKibiByte("% .2f / % .2f"),
			// decor.OnComplete(
			// 	decor.EwmaETA(decor.ET_STYLE_GO, 60, decor.WCSyncSpaceR), "done",
			// ),
			// decor.Percentage(decor.WCSyncSpace, decor.WCSyncSpaceR),
		// ),
	)

	reader := bufio.NewReader(file)
	proxyReader := bar.ProxyReader(reader)
	defer proxyReader.Close()

	// buffer := bytes.NewBuffer(data)
	err = extract.Archive(ctx, proxyReader, to, nil)
	if err != nil {
		file.Close()
		HandleError(err)
	}

	file.Close()
	pb.Wait()
}
