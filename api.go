package poeninjaapi

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	*http.Client
}

type Stats struct {
	ID                       string
	NextChangeID             string
	APIBytesDownloaded       int64
	StashTabsProcessed       int64
	APICalls                 int64
	CharacterBytesDownloaded int64
	CharacterAPICalls        int64
	LadderBytesDownloaded    int64
	LadderAPICalls           int64
}

const (
	apiUrl = "http://api.poe.ninja/api/Data/GetStats"
)

func NewClient(timeout time.Duration) *Client {
	return &Client{
		Client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   timeout,
		},
	}
}

func (c *Client) Fetch(ctx context.Context) (*Stats, error) {
	req, _ := http.NewRequest("GET", apiUrl, nil)
	req = req.WithContext(ctx)

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	defer func() { io.Copy(ioutil.Discard, resp.Body) }()

	var stats Stats
	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
