package main

import (
	"io"
	"os"
	"log"
	"flag"
	"net/http"
	"github.com/tie/hseguest/go-requests"
)

func main() {
	action := flag.String(
		"do", "login",
		`Captive Portal action ("login" or "logout")`,
	)
	username := flag.String(
		"user", "hseguest",
		`username for "login" request`,
	)
	password := flag.String(
		"pass", "hsepassword",
		`password for "login" request`,
	)
	logPrefix := flag.String(
		"prefix", "hseguest: ",
		`logger prefix`,
	)
	flag.Parse()

	logger := log.New(os.Stderr, *logPrefix, log.Lshortfile)

	requests := map[string]func() (*http.Request, error){
		"logout": requests.NewLogout,
		"login": func() (*http.Request, error) {
			return requests.NewLogin(*username, *password)
		},
	}

	f, ok := requests[*action]
	if !ok {
		logger.Fatal(*action)
	}

	req, err := f()
	if err != nil {
		logger.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Fatal(err)
	}
	defer resp.Body.Close()

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		logger.Print(err)
	}
}
