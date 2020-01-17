package gografana

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

type Folder struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	HasAcl    bool      `json:"hasAcl"`
	CanSave   bool      `json:"canSave"`
	CanEdit   bool      `json:"canEdit"`
	CanAdmin  bool      `json:"canAdmin"`
	CreatedBy string    `json:"createdBy"`
	Created   time.Time `json:"created"`
	UpdatedBy string    `json:"updatedBy"`
	Updated   time.Time `json:"updated"`
	Version   int       `json:"version"`
}

func (c *Client) GetAllFolders(ctx context.Context) (arr []Folder, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/api/folders", nil)
	if err != nil {
		return
	}
	err = c.doJsonRequest200(req, &arr)
	return
}

func (c *Client) GetFolderByUID(ctx context.Context, uid string) (folder Folder, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/api/folders/"+uid, nil)
	if err != nil {
		return
	}
	err = c.doJsonRequest200(req, &folder)
	return
}

func (c *Client) CreateFolder(ctx context.Context, title, uid string) (folder Folder, err error) {
	p := struct {
		UID   string `json:"uid,omitempty"`
		Title string `json:"title"`
	}{
		UID:   uid,
		Title: title,
	}
	b, err := jsonMarshal(p)
	if err != nil {
		return
	}
	req, err := c.newRequest(ctx, http.MethodPost, "/api/folders", bytes.NewReader(b))
	if err != nil {
		return
	}
	setRequestContentTypeJson(req)
	err = c.doJsonRequest200(req, &folder)
	return
}
