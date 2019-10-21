package cerve

import (
	"encoding/json"
	"ngen.co.jp/cerve/pkg/cert"
)

type Certs struct {
	cert.Certs
}

func (c *Certs) JSON() string {
	bytes, err := json.Marshal(c.Certs)
	if err != nil {
		return ""
	}
	return string(bytes)
}
