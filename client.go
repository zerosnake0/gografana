package gografana

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type Client struct {
	addr   string
	key    string
	client *http.Client

	Debug bool
}

func NewClient(addr, key string, client *http.Client) *Client {
	c := Client{
		addr: addr,
		key:  key,
	}
	if client == nil {
		c.client = http.DefaultClient
	} else {
		c.client = client
	}
	return &c
}

func (c *Client) newRequest(ctx context.Context, method, uri string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.addr+uri, body)
	if err != nil {
		return nil, err
	}
	if c.key != "" {
		req.Header.Set("Authorization", "Bearer "+c.key)
	}
	return req, nil
}

func (c *Client) doRequest(req *http.Request) (int, []byte, error) {
	if c.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Printf("request: >>>>>>>>>>\n%s", dump)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	if c.Debug {
		dump, _ := httputil.DumpResponse(resp, true)
		log.Printf("response: <<<<<<<<<<\n%s", dump)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, b, err
}

func (c *Client) doRequest200(req *http.Request) ([]byte, error) {
	code, body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("status code %d instead of 200: %s", code, body)
	}
	return body, nil
}

func (c *Client) doJsonRequest200(req *http.Request, o interface{}) error {
	body, err := c.doRequest200(req)
	if err != nil {
		return err
	}
	if o == nil {
		return nil
	}
	return jsonUnmarshal(body, o)
}
