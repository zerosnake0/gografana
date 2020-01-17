package gografana

import (
	"context"
	"encoding/json"
	"net/http"
)

type DataSourceType string

type DataSource struct {
	ID                int             `json:"id"`
	OrgID             int             `json:"orgId"`
	Name              string          `json:"name"`
	Type              DataSourceType  `json:"type"`
	Access            string          `json:"access"`
	URL               string          `json:"url"`
	Password          string          `json:"password"`
	User              string          `json:"user"`
	Database          string          `json:"database"`
	BasicAuth         bool            `json:"basicAuth"`
	BasicAuthUser     string          `json:"basicAuthUser"`
	BasicAuthPassword string          `json:"basicAuthPassword"`
	IsDefault         bool            `json:"isDefault"`
	JsonData          json.RawMessage `json:"jsonData"`
	SecureJsonData    json.RawMessage `json:"secureJsonData"`
}

type DataSourceList []DataSource

func (list DataSourceList) IndexByName(name string) int {
	for i := range list {
		if list[i].Name == name {
			return i
		}
	}
	return -1
}

func (c *Client) GetAllDataSources(ctx context.Context) (arr DataSourceList, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/api/datasources", nil)
	if err != nil {
		return
	}
	err = c.doJsonRequest200(req, &arr)
	return
}
