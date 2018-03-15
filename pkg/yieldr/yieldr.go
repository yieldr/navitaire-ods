package yieldr

import (
	"bytes"
	"context"
	"fmt"
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

func (yldr *Yieldr) Upload(project int, b []byte) error {
	url := fmt.Sprintf("%s/projects/%d/flights/upload", yldr.url, project)
	_, err := yldr.client.Post(url, "application/x-jsonlines", bytes.NewReader(b))
	return err
}
