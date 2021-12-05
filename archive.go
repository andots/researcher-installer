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
	fmt.Printf("Extracting %s ..... ", filename)
	data, err := ioutil.ReadFile(path)
	HandleError(err)

	buffer := bytes.NewBuffer(data)
	err = extract.Archive(context.Background(), buffer, to, nil)
	if err != nil {
		HandleError(err)
	}
	fmt.Print("Done!\n")
}

	// s := spinner.New(spinner.CharSets[36], 100 * time.Millisecond)
	// s.Prefix = fmt.Sprintln("Extracting files...  ")
	// s.Start()

	// fi, err := os.Stat(path)
	// HandleError(err)

	// file, err := os.Open(path)
	// if err != nil {
	// 	file.Close()
	// 	HandleError(err)
	// }

	// fi, err := file.Stat()
	// if err != nil {
	// 	file.Close()
	// 	HandleError(err)
	// }

	// pb := mpb.NewWithContext(ctx, mpb.WithWidth(60))
	// title := fmt.Sprintf("[Extract]: %v", fi.Name())
	// bar := pb.AddBar(
	// 	fi.Size(),
	// 	mpb.PrependDecorators(
	// 		decor.Name(title, decor.WCSyncSpaceR),
	// 		// decor.Spinner(nil, decor.WCSyncSpace),
	// 	),
	// 	// mpb.AppendDecorators(
	// 		// decor.CountersKibiByte("% .2f / % .2f"),
	// 		// decor.OnComplete(
	// 		// 	decor.EwmaETA(decor.ET_STYLE_GO, 60, decor.WCSyncSpaceR), "done",
	// 		// ),
	// 		// decor.Percentage(decor.WCSyncSpace, decor.WCSyncSpaceR),
	// 	// ),
	// )

	// reader := bufio.NewReader(file)
	// proxyReader := bar.ProxyReader(reader)
	// defer proxyReader.Close()
