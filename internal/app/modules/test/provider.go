package test

import "ecommerce/pkg/deuterium"

type TestProvider struct {
	v string
}

func (p *TestProvider) Test() {
	logger := &deuterium.Logger{Context: "Test"}
	logger.Debug(p.v)
}
