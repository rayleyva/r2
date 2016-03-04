package main

import (
	"io"
	"net/http"
)

type Req struct {
	url    string
	method string
	head   http.Header
	body   io.Reader
}

type Rep struct {
	rawRep *http.Response
}

type Report struct {
	//
}

var defaultCli *http.Client = &http.Client{}

func (r *Req) Launch() (*Rep, error) {
	req, err1 := http.NewRequest(r.method, r.url, r.body)
	if err1 != nil {
		return nil, err1
	}

	// Add Headers here

	rep, err2 := defaultCli.Do(req)
	if err2 != nil {
		return nil, err2
	}
	defer rep.Body.Close()

	return &Rep{
		rawRep: rep,
	}, nil
}
