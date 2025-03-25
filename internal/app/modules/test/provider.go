package test

import "ecommerce/pkg/api"

type TestProvider struct {
	v string
}

func (p *TestProvider) Test() {
	logger := &api.Logger{Context: "Test"}
	logger.Debug(p.v)
}
