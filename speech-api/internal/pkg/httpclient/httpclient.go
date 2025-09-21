package httpclient

import (
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"time"
)

type (
	HttpClient interface {
		Execute([]byte) ([]byte, error)
	}

	httpClient struct {
		client *http.Client
	}
)

func NewHttpClient() HttpClient {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,  // connection timeout
			KeepAlive: 30 * time.Second, // keep connections alive
		}).DialContext,
		MaxIdleConns:          100,              // max idle connections
		MaxIdleConnsPerHost:   10,               // per host
		IdleConnTimeout:       90 * time.Second, // idle connection timeout
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second, // overall request timeout
	}

	return &httpClient{client: client}
}

func (h *httpClient) Execute(data []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://localhost:8000/transcribe", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
