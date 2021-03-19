package gohttp

import (
	"time"
	"net/http"
)

type clientBuilder struct {
	maxIdleConnections int
	connectionTimeout time.Duration
	responseTimeout time.Duration
	disableTimeouts bool
	headers http.Header
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetMaxIdleConnections(maxConn int) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	DisableTimeouts(disableTimeouts bool) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (b *clientBuilder) Build() Client {
	client := client{
		builder: b,
	}

	return &client
}

func (b *clientBuilder) SetMaxIdleConnections(maxConn int) ClientBuilder {
	b.maxIdleConnections = maxConn
	return b
}

func (b *clientBuilder) DisableTimeouts(disableTimeouts bool) ClientBuilder {
	b.disableTimeouts = disableTimeouts
	return b
}

func (b *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	b.connectionTimeout = timeout
	return b
}

func (b *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	b.responseTimeout = timeout
	return b
}

func (b *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	b.headers = headers
	return b
}
