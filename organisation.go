package gografana

import (
	"context"
	"net/http"
)

type OrgAddress struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	ZipCode  string `json:"zipCode"`
	State    string `json:"state"`
	Country  string `json:"country"`
}

type Organisation struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Address OrgAddress `json:"address"`
}

func (c *Client) GetCurrentOrganisation(ctx context.Context) (org Organisation, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/api/org", nil)
	if err != nil {
		return
	}
	err = c.doJsonRequest200(req, &org)
	return
}
