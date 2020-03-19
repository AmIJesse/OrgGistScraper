package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/github"
)

func main() {

	org := flag.String("org", "", "Organization's Github Name")
	saveOut := flag.Bool("o", false, "Output Results To Text File (Enabled if -d flag is set)")
	dlGists := flag.Bool("d", false, "Automatically Download All Gists")

	flag.Parse()
	if *org == "" {
		flag.Usage()
		return
	}

	if *dlGists {
		newFolderName := *org + "-" + strings.Split(time.Now().String(), " ")[0]

		if _, err := os.Stat(newFolderName); os.IsNotExist(err) {
			os.Mkdir(newFolderName, os.ModePerm)
		}

		err := os.Chdir(newFolderName)
		if err != nil {
			fmt.Println("Error: Unable to move into new working directory")
			fmt.Println(err.Error())
			return
		}
		*saveOut = true
	}

	var outFile *os.File
	var err error
	if *saveOut {
		outFile, err = os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
		if err != nil {
			fmt.Println("Error: Unable to open file for logging - output.txt")
			fmt.Println(err.Error())
			return
		}
		defer outFile.Close()
	}

	// Create new API client with no API key. Due to having no key it allows 10 requests per minute
	g := github.NewClient(nil)

	// Search for members of the organization
	userSearchResults, _, err := g.Organizations.ListMembers(context.TODO(), *org, nil)
	if err != nil {
		fmt.Printf("Error: Unable to list members or the organization")
		fmt.Println(err.Error())
		return
	}
	time.Sleep(6 * time.Second) // Need to wait so we don't timeout the API

	// Set max timeout for HTTP requests (fetching gists) to 10 seconds so we don't hang forever
	client := &http.Client{Timeout: 10 * time.Second}

	for _, user := range userSearchResults {
		fmt.Println(*user.Login)
		if *saveOut {
			outFile.WriteString(*user.Login + "\n")
		}

		// Get all gists from each member of the organization
		gists, _, err := g.Gists.List(context.TODO(), *user.Login, nil)
		if err != nil {
			fmt.Println("Error: Unable to list gists from user")
			fmt.Println(err.Error())
		}
		time.Sleep(6 * time.Second) // Need to wait so we don't timeout the API

		for _, gist := range gists {

			// Check each file inside each gist
			for _, file := range gist.Files {
				fmt.Println("--", *file.Filename, "-", *gist.HTMLURL)
				if *saveOut {
					outFile.WriteString("-- " + *file.Filename + " - " + *gist.HTMLURL + "\n")
				}
				if *dlGists {
					fileName := *user.Login + "-" + *file.Filename

					// Fetch the raw file, and save it
					resp, err := client.Get(*file.RawURL)
					if err != nil {
						fmt.Println("Unable to get", *file.RawURL)
						continue
					}
					body, err := ioutil.ReadAll(resp.Body)
					resp.Body.Close()
					if err != nil {
						fmt.Println("Unable to read", *file.RawURL)
						continue
					}
					ioutil.WriteFile(fileName, body, 0600)
				}
			}
		}
	}
}
