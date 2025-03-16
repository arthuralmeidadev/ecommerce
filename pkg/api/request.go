package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type PathParam struct {
	name  string
	value string
}

func (p *PathParam) String() string {
	return p.value
}

func (p *PathParam) ParseInt() (int, error) {
	return strconv.Atoi(p.value)
}

func (p *PathParam) ParseUint() (uint, error) {
	n, err := strconv.ParseUint(p.value, 10, 0)
	if err != nil {
		return 0, nil
	}

	return uint(n), err
}

func (p *PathParam) ParseBool() (bool, error) {
	return strconv.ParseBool(p.value)
}

func (p *PathParam) ParseTime(layout string) (time.Time, error) {
	return time.Parse(layout, p.value)
}

type request struct {
	r             *http.Request
	reqPathParams map[string]string
}

func (req *request) Header(name string) string {
	return req.r.Header.Get(name)
}

func (req *request) Cookie(name string) (*http.Cookie, error) {
	return req.r.Cookie(name)
}

func (req *request) PathParam(name string) *PathParam {
	return &PathParam{
		name:  name,
		value: req.reqPathParams[name],
	}
}

func (req *request) Body(v any) error {
	return json.NewDecoder(req.r.Body).Decode(v)
}
