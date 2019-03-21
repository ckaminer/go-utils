package http_helpers

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type MockDoErrorClient struct{}

func (client *MockDoErrorClient) Do(*http.Request) (*http.Response, error) {
	return nil, errors.New("client#Do failed")
}

type MockHttpClient struct {
	Status  int
	Payload []byte
}

func (client *MockHttpClient) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: client.Status,
		Body:       ioutil.NopCloser(bytes.NewReader(client.Payload)),
	}, nil
}

func NewClientMock(status int, payload []byte) MockHttpClient {
	return MockHttpClient{
		Status:  status,
		Payload: payload,
	}
}
