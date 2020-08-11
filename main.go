package main

import (
	"ghpages-backend/api"
	"log"
)

func main() {
	notion := api.NewNotionWrapper()
	notion.PrintCategories()
	notion.PrintCategoryPages("Java")
	notion.PrintCategoryPosts("Java")
}

// TODO: implement handleError function for more generalized cases
func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
