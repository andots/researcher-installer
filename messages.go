package main

import (
	"fmt"
	"os"

	"github.com/deiwin/interact"
	"github.com/fatih/color"
)

func ShowWelcomeMessage() {
	color.Set(color.FgBlue)
	fmt.Println("---------------------------------------------------")
	fmt.Println("|         RE:SEARCHER Backend Installer           |")
	fmt.Println("---------------------------------------------------")
	color.Unset()

	fmt.Println("This installer will setup Elasticsearch for RE:SEARCHER as search backend.")
	fmt.Println("")
	fmt.Println("Elasticsearch will be installed to the following location.")
	fmt.Println("[ " + GetAppDirPath() + " ]")
	fmt.Println("")
}

func ShowConfirmation() bool {
	actor := interact.NewActor(os.Stdin, os.Stdout)
	confirmed, err := actor.Confirm("Press y to continue, n to cancel", interact.ConfirmNoDefault)
	HandleError(err)

	return confirmed
}

func Red(text string) {
	p := color.New(color.FgRed).PrintlnFunc()
	p(text)
}
