package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	fmt.Println("")
	color.Set(color.FgGreen)
	fmt.Println("Setup successfully done!")
	color.Unset()
	fmt.Println("")

	switch runtime.GOOS {
	case "windows":
		cmdPath := filepath.Join(GetESPath(), "elasticsearch.bat")
		fmt.Printf("Execute %s to start Elasticsearch.\n", cmdPath)
		fmt.Println("You can also register Elasticsearch as a Windows service to run in the background or start automatically after login.")
		fmt.Println("")
		fmt.Println("Please see the document to get more information.")
		fmt.Println("https://github.com/andots/researcher-webextension/blob/main/docs/windows/README.md")
	default:
		cmdPath := filepath.Join(GetESPath(), "elasticsearch")
		fmt.Printf("Please run %s to start Elasticsearch.\n", cmdPath)
	}
	fmt.Println("")
}

func Red(text string) {
	p := color.New(color.FgRed).PrintlnFunc()
	p(text)
}
