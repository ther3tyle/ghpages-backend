package api

import (
	"fmt"
	"ghpages-backend/api/models"
	"github.com/kjk/notionapi"
)

const ROOT_ID = "6e0247a5-1f5e-4d00-a8c2-178725b32480"

type notionWrapper struct {
	client     *notionapi.Client
	rootPage   *notionapi.Page
	categories map[string]*notionapi.Page
	pages      map[string]map[string]*notionapi.Page
	posts      map[string]map[string]*models.Post
}

func NewNotionWrapper() *notionWrapper {
	wrapper := &notionWrapper{
		client:     &notionapi.Client{},
		categories: make(map[string]*notionapi.Page),
		pages:      make(map[string]map[string]*notionapi.Page),
		posts:      make(map[string]map[string]*models.Post),
	}
	wrapper.initRootPage()
	wrapper.initCategoryPages()
	for catName := range wrapper.categories {
		wrapper.makeCategoryPosts(catName)
	}
	return wrapper
}

func (api *notionWrapper) initRootPage() {
	page, err := api.client.DownloadPage(ROOT_ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	api.rootPage = page
}

func (api *notionWrapper) initCategoryPages() {
	catPages := api.fetchSubPages(api.rootPage)

	for _, page := range catPages {
		api.categories[page.Root().Title] = page
		api.fetchCategoryPages(page)
	}
}

func (api *notionWrapper) fetchCategoryPages(catPage *notionapi.Page) {
	if catPage.Root().Type != "page" {
		return
	}

	title := catPage.Root().Title
	categoryPages := make(map[string]*notionapi.Page)

	pages := api.fetchSubPages(catPage)

	for _, subPage := range pages {
		categoryPages[subPage.Root().Title] = subPage
	}

	api.pages[title] = categoryPages
}

func (api *notionWrapper) fetchSubPages(page *notionapi.Page) []*notionapi.Page {
	var pages []*notionapi.Page
	for _, v := range page.GetSubPages() {
		p, err := api.client.DownloadPage(v)
		if err != nil {
			continue
		}
		pages = append(pages, p)
	}
	return pages
}

func (api *notionWrapper) PrintCategories() {
	fmt.Print("[")
	curr, sz := 0, len(api.pages)
	for key := range api.pages {
		fmt.Printf("%s", key)
		if curr < sz-1 {
			fmt.Print(", ")
		}
		curr++
	}
	fmt.Print("]\n")
}

func (api *notionWrapper) PrintCategoryPages(catName string) {
	fmt.Print("[")

	curr, sz := 0, len(api.pages[catName])
	for pageTitle := range api.pages[catName] {
		fmt.Print(pageTitle)
		if curr < sz-1 {
			fmt.Print(", ")
		}
		curr++
	}
	fmt.Print("]\n")
}

func (api *notionWrapper) makeCategoryPosts(catName string) {
	api.fetchCategoryPages(api.categories[catName])

	if api.posts[catName] == nil {
		api.posts[catName] = make(map[string]*models.Post)
	}

	for pageTitle, page := range api.pages[catName] {
		post := BuildPost(page)
		api.posts[catName][pageTitle] = post
	}
}

func (api *notionWrapper) PrintCategoryPosts(catName string) {
	for postTitle, post := range api.posts[catName] {
		fmt.Println(postTitle, *post)
	}
}
