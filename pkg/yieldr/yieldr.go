package yieldr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type YieldrConfig struct {
	Addr,
	ClientID,
	ClientSecret string
}

type Yieldr struct {
	url    string
	client *http.Client
	config *YieldrConfig
}

type YieldrResponse struct {
	Message string `json:"message"`
	Errors  []struct {
		Line   int `json:"line"`
		Errors []struct {
			Code    int         `json:"code"`
			Message string      `json:"message"`
			Field   string      `json:"field"`
			Value   interface{} `json:"value"`
		} `json:"errors"`
	}
}

func New(c *YieldrConfig) *Yieldr {
	cc := &OAuth{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     "https://" + c.Addr + "/api/oauth/token",
	}
	return &Yieldr{
		url:    "https://" + c.Addr + "/api",
		client: cc.Client(context.Background()),
		config: c,
	}
}

func (yldr *Yieldr) Upload(project int, payload []byte) (*YieldrResponse, error) {

	url := fmt.Sprintf("%s/projects/%d/flights/upload", yldr.url, project)

	req, err := http.NewRequest("PUT", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-jsonlines")

	res, err := yldr.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var v *YieldrResponse
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 && res.StatusCode < 500 {
		return v, errors.New("Upload failed")
	}

	return v, nil
}
