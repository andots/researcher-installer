package main

import (
	"fmt"
	"os"

	latest "github.com/tcnksm/go-latest"
)

const VERSION string = "v0.9.4"

func CheckVersion() {
	githubTag := &latest.GithubTag{
		Owner: "andots",
		Repository: "researcher-installer",
	}

	fmt.Printf("Checking installer version... ")

	res, err := latest.Check(githubTag, VERSION)
	if (err != nil) {
		fmt.Printf("\n[ERROR] Checking version is failed: %s\n\n", err)
	} else {
		if res.Outdated {
			fmt.Printf("Latest version is %s. Please upgrade!\n\n", res.Current)
			BlockForWindows()
			os.Exit(1)
		} else {
			fmt.Printf("Latest! v%s\n\n", res.Current)
		}
	}
}
