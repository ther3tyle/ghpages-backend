package models

import "time"

type Asset struct {
	Value     string
	CodeLang  string
	AssetType string
}

type Post struct {
	ID         string
	Assets     []Asset
	ModifiedAt time.Time
	CreatedAt  time.Time
}