package cerve

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"ngen.co.jp/cerve/pkg/cert"
)

type Cerve struct {
	Input string
}

// TODO: implements option
type Option struct {
}

func NewCerve(args []string, option Option) (*Cerve, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("")
	}
	return &Cerve{Input: args[0]}, nil
}

func (c *Cerve) Verify() (*Certs, error) {
	d, err := c.loadInputFile(c.Input)
	if err != nil {
		return nil, err
	}
	return &Certs{Certs: c.verifyCerts(d.Domains)}, nil
}

func (c *Cerve) loadInputFile(input string) (*Define, error) {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		return nil, err
	}
	define := &Define{}
	if yaml.Unmarshal(bytes, define) != nil {
		return nil, err
	}
	return define, nil
}

func (c *Cerve) verifyCerts(targets []string) cert.Certs {
	certs := make(cert.Certs, len(targets))
	for i, t := range targets {
		certs[i] = cert.NewCert(t)
	}
	return certs
}
