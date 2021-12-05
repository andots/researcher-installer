package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/deiwin/interact"
	"github.com/fatih/color"
)

func ShowWelcomeMessage() {
	banner := `
+-------------------------------------------------------------------+
|                                                                   |
|                  RE:SEARCHER Backend Installer                    |
|                                                                   |
+-------------------------------------------------------------------+
`
	fmt.Println(banner)
	fmt.Println("This installer will setup Elasticsearch for RE:SEARCHER as search backend.")
	fmt.Println("")
	fmt.Println("Elasticsearch will be installed to the following location.")
	fmt.Println("")
	Red("[ " + GetAppPath() + " ]")
	fmt.Println("")
}

func ShowConfirmation() bool {
	actor := interact.NewActor(os.Stdin, os.Stdout)
	confirmed, err := actor.Confirm("Press y to continue, n to cancel", interact.ConfirmNoDefault)
	HandleError(err)

	return confirmed
}

func ShowEndMessage() {
	color.Set(color.FgGreen)
	fmt.Println("")
	fmt.Println("Setup successfully done!")
	switch runtime.GOOS {
	case "windows":
		fmt.Printf("Please run %s/elasticsearch.bat to start server.\n", GetESPath())
	default:
		fmt.Printf("Please run %s/elasticsearch to start server.\n", GetESPath())
	}
	fmt.Println("")
	color.Unset()
}

func Red(text string) {
	p := color.New(color.FgRed).PrintlnFunc()
	p(text)
}
