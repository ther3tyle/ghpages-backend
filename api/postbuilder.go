package api

import (
	"fmt"
	"ghpages-backend/api/models"
	"github.com/kjk/notionapi"
)

func BuildPost (page *notionapi.Page) *models.Post {
	post := models.Post{
		ID:         page.ID,
		Assets:     make([]models.Asset, 0),
		ModifiedAt: page.Root().LastEditedOn(),
		CreatedAt:  page.Root().CreatedOn(),
	}

	page.ForEachBlock(func(block *notionapi.Block) {
		if block == page.Root() {
			return
		}
		asset, err := mapBlockToAsset(block)
		if err != nil {
			return
		}
		post.Assets = append(post.Assets, *asset)
	})
	return &post
}

func mapBlockToAsset(block *notionapi.Block) (*models.Asset, error) {
	switch block.Type {
	case "text":
		if len(block.InlineContent) == 0 {
			return nil, fmt.Errorf("Found empty text block!\n")
		}
		return &models.Asset{
			Value:     block.InlineContent[0].Text,
			AssetType: "text",
		}, nil
	case "code":
		return &models.Asset{
			Value:     block.InlineContent[0].Text,
			CodeLang:  block.CodeLanguage,
			AssetType: "code",
		}, nil
	case "image":
		return &models.Asset{
			Value:     block.ImageURL,
			AssetType: "image",
		}, nil
	case "header", "sub_header", "sub_sub_header":
		text := ""
		if block.InlineContent == nil {
			text = ""
		} else {
			text = block.InlineContent[0].Text
		}
		return &models.Asset{
			Value:     text,
			AssetType: block.Type,
		}, nil
	case "divider":
		return &models.Asset{
			AssetType: "divider",
		}, nil
	default:
		return &models.Asset{
			Value:     block.Title,
			AssetType: block.Type,
		}, nil
	}
}