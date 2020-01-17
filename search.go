package gografana

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

type SearchType string

const (
	SearchTypeFolder SearchType = "dash-folder"
	SearchTypeDB     SearchType = "dash-db"
)

type SearchOptions struct {
	Query        string
	Tag          []string
	Type         SearchType
	DashboardIDs []int
	FolderIDs    []int
	Starred      bool
	Limit        int
	Page         int
}

func arrIntToStr(in []int) []string {
	out := make([]string, len(in))
	for i := range in {
		out[i] = strconv.Itoa(in[i])
	}
	return out
}

func (opts *SearchOptions) encode() string {
	if opts == nil {
		return ""
	}
	values := url.Values{}
	if opts.Query != "" {
		values.Set("query", opts.Query)
	}
	if len(opts.Tag) > 0 {
		values["tag"] = opts.Tag
	}
	if opts.Type != "" {
		values.Set("type", string(opts.Type))
	}
	if len(opts.DashboardIDs) > 0 {
		values["dashboardIds"] = arrIntToStr(opts.DashboardIDs)
	}
	if len(opts.FolderIDs) > 0 {
		values["folderIds"] = arrIntToStr(opts.FolderIDs)
	}
	if opts.Starred {
		values.Set("starred", "true")
	}
	if opts.Limit > 0 {
		values.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.Page > 0 {
		values.Set("page", strconv.Itoa(opts.Page))
	}
	return values.Encode()
}

type SearchResult struct {
	ID          int        `json:"id"`
	UID         string     `json:"uid"`
	Title       string     `json:"title"`
	URL         string     `json:"url"`
	Type        SearchType `json:"type"`
	Tag         []string   `json:"tags"`
	IsStarred   bool       `json:"isStarred"`
	URI         string     `json:"uri"` // deprecated in Grafana v5.0
	Slug        string     `json:"slug"`
	FolderID    int        `json:"folderId"`
	FolderUID   string     `json:"folderUid"`
	FolderTitle string     `json:"folderTitle"`
	FolderURL   string     `json:"folderUrl"`
}

type SearchResults []SearchResult

func (arr SearchResults) IndexByTitle(title string) int {
	for i := range arr {
		if arr[i].Title == title {
			return i
		}
	}
	return -1
}

func (arr SearchResults) IndexBy(cb func(*SearchResult) bool) int {
	for i := range arr {
		if cb(&arr[i]) {
			return i
		}
	}
	return -1
}

func (c *Client) Search(ctx context.Context, options *SearchOptions) (arr SearchResults, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/api/search?"+options.encode(), nil)
	if err != nil {
		return
	}
	err = c.doJsonRequest200(req, &arr)
	return
}
