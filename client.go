package transdevclient

import (
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient() Client {
	httpClient := http.DefaultClient
	return Client{client: httpClient}
}
