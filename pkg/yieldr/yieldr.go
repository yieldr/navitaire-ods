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
	client *http.Client
	config *YieldrConfig
}

func New(c *YieldrConfig) *Yieldr {
	cc := &Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     "https://" + c.Addr + "/oauth/token",
	}
	return &Yieldr{
		client: cc.Client(context.Background()),
		config: c,
	}
}

func (y *Yieldr) Projects() {

}

func (y *Yieldr) Upload(project int, b []byte) error {
	url := fmt.Sprintf("https://%s/projects/%d/flights/upload", y.config.Addr, project)
	_, err := y.client.Post(url, "application/x-jsonlines", bytes.NewReader(b))
	return err
}
