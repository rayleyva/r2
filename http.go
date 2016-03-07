package main

import (
	"bytes"
	"errors"
	"net/http"
	"regexp"
)

type Req struct {
	url      string
	method   string
	header   map[string]string
	body     *bytes.Buffer
	inited   bool
	launched bool
}

func NewReq() *Req {
	return &Req{
		url:      "",
		method:   "",
		header:   nil,
		body:     nil,
		inited:   false,
		launched: false,
	}
}

type Rep struct {
	rawRep *http.Response
}

func NewRep() *Rep {
	return &Rep{
		rawRep: nil,
	}
}

type Report struct {
	//
}

var gCli *http.Client
var gReq *Req
var gRep *Rep
var gReport *Report
var urlReg *regexp.Regexp

var ErrDoNotRedirect = errors.New("Do not redirect")

func init() {
	gReq = NewReq()
	gRep = NewRep()
	urlReg = regexp.MustCompile("^(http|https)://.+")
	gCli = &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			return errors.New("Do not redirect")
		},
	}
}

func (r *Req) Init() {
	if !r.inited {
		r.url = ""
		r.method = ""
		r.header = make(map[string]string)
		r.body = new(bytes.Buffer)
		r.inited = true
		r.launched = false
	}
}

func (r *Req) Cleanup() {
	r.inited = false
}

var httpMethods map[string]bool = map[string]bool{
	"GET":     true,
	"POST":    true,
	"HEAD":    true,
	"TRACE":   true,
	"PUT":     true,
	"DELETE":  true,
	"OPTIONS": true,
	"CONNECT": true,
}

func (r *Req) SetMethod(method string) error {
	if !r.inited {
		return errors.New("Uninitialized request")
	}
	if len(r.method) != 0 {
		return errors.New("Method has already been set")
	}
	if _, ok := httpMethods[method]; !ok {
		return errors.New("Unrecognized method: " + method)
	}
	r.method = method
	return nil
}

func (r *Req) SetUrl(url string) error {
	if !r.inited {
		return errors.New("Uninitialized request")
	}
	if len(r.url) != 0 {
		return errors.New("Url has already been set")
	}
	if !urlReg.MatchString(url) {
		return errors.New("Illegal url of: " + url)
	}
	r.url = url
	return nil
}

func (r *Req) check() error {
	if !r.inited {
		return errors.New("Uninitialized request")
	}
	if len(r.url) == 0 {
		return errors.New("http request missing url (eg: https://github.com/)")
	}
	if len(r.method) == 0 {
		return errors.New("http request missing method (eg: get|post|put|delete)")
	}
	return nil
}

func shouldRedirectGet(statusCode int) bool {
	switch statusCode {
	case http.StatusMovedPermanently, http.StatusFound,
		http.StatusSeeOther, http.StatusTemporaryRedirect:
		return true
	}
	return false
}

func (r *Req) Launch(rep *Rep) error {
	if r.launched {
		return nil
	}
	if err := r.check(); err != nil {
		return err
	}
	if rep.rawRep != nil && !rep.rawRep.Close {
		rep.rawRep.Body.Close()
	}

	req, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		return err
	}

	if len(r.header) > 0 {
		for k, v := range r.header {
			req.Header.Add(k, v)
		}
	}

	rep.rawRep, err = gCli.Do(req)
	r.launched = true

	if err != nil {
		if rep.rawRep != nil && shouldRedirectGet(rep.rawRep.StatusCode) {
			return nil
		}
		return err
	}

	return nil
}

func (rep *Rep) Cleanup() error {
	if rep.rawRep != nil && !rep.rawRep.Close {
		return rep.rawRep.Body.Close()
	}
	return nil
}
