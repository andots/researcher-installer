// This code is based on https://github.com/tcnksm/ghr/blob/master/version.go

package main

import (
	"bytes"
	"fmt"
	"time"

	latest "github.com/tcnksm/go-latest"
)

const APP_NAME = "RE:SEARCHER Backend Installer"
const VERSION string = "v0.9.0"

// GitCommit describes latest commit hash.
// This is automatically extracted by git describe --always.
var GitCommit string

// OutputVersion checks the current version and compares it against releases
// available on GitHub. If there is a newer version available, it prints an
// update warning.
func OutputVersion() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s version v%s", APP_NAME, VERSION)

	if len(GitCommit) != 0 {
		fmt.Fprintf(&buf, " (%s)", GitCommit)
	}
	fmt.Fprintf(&buf, "\n")

	// Check latest version is release or not.
	verCheckCh := make(chan *latest.CheckResponse)
	go func() {
		githubTag := &latest.GithubTag{
			Owner:      "andots",
			Repository: "researcher-installer",
		}

		res, err := latest.Check(githubTag, VERSION)
		if err != nil {
			// Don't return error
			fmt.Printf("[ERROR] Check latest version is failed: %s\n\n", err)
			return
		}
		verCheckCh <- res
	}()

	select {
	case <-time.After(2 * time.Second):
	case res := <-verCheckCh:
		if res.Outdated {
			fmt.Fprintf(&buf,
				"Latest version is v%s, please upgrade!\n",
				res.Current)
		}
	}

	return buf.String()
}
