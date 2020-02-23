package parser_mail

import (
	"net/http"
	"time"
)

type ParserHttpClient struct {
	*http.Client
}

func NewParserHttpClient(maxConn int) *ParserHttpClient {
	transport := &http.Transport{
		DisableKeepAlives:  true,
		MaxIdleConns:       maxConn,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 10,
	}

	return &ParserHttpClient{
		client,
	}
}
