package main

import (
	"fmt"
	"github.com/kjk/notionapi"
	"log"
)

const PAGEID = "6e0247a5-1f5e-4d00-a8c2-178725b32480"

func main() {
	client := &notionapi.Client{}
	fmt.Println(client.AuthToken)
	client.AuthToken = "TOKEN_VALUE"
	root, err := client.DownloadPage(PAGEID)
	if err != nil {log.Fatal(err.Error())}


	categoryStrings := root.GetSubPages()
	for i := range categoryStrings {
		tmpPage, err := client.DownloadPage(categoryStrings[i])
		if err != nil {continue}
		categoryStrings[i] = tmpPage.Root().Title
	}

	for _, v := range categoryStrings {
		fmt.Println("CATEGORY NAME", v)
		post, err := client.DownloadPage(v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\t%s\n", post.Root().Title)
	}
}