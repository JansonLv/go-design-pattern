package main

import (
	"errors"
	"net/http"
)

type proxyClient struct {
	*http.Client
}

func (client *proxyClient) Do(req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	return client.Client.Do(req)
}

func main() {
	req, err := http.NewRequest("POST", "https://baidu.com", nil)
	if err != nil {
		return
	}
	client := &proxyClient{http.DefaultClient}
	client.Do(req)
}
