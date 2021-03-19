package gohttp

import (
	"net/http"
	"sync"
)

type client struct{
	builder *clientBuilder
	client *http.Client
	clientOnce sync.Once
}

type Client interface {
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, headers http.Header, body interface{}) (*Response, error)
	Put(url string, headers http.Header, body interface{}) (*Response, error)
	Patch(url string, headers http.Header, body interface{}) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
	SetHeaders(headers http.Header)
}


func (c *client) Get(url string, headers http.Header) (*Response, error){
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *client) Post(url string, headers http.Header, body interface{}) (*Response, error){
	return c.do(http.MethodPost, url, headers, body)
}

func (c *client) Put(url string, headers http.Header, body interface{}) (*Response, error){
	return c.do(http.MethodPut, url, headers, body)
}

func (c *client) Patch(url string, headers http.Header, body interface{}) (*Response, error){
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *client) Delete(url string, headers http.Header) (*Response, error){
	return c.do(http.MethodDelete, url, headers, nil)
}

func (c *client) SetHeaders(headers http.Header) {
	c.builder.headers = headers
}
