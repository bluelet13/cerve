package cert

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"net"
	"strings"
	"time"
)

const (
	Port    string        = "443"
	Timeout time.Duration = 5
)

type Cert struct {
	DomainName string     `json:"domain_name"`
	Issuer     *pkix.Name `json:"issuer,omitempty"`
	Subject    *pkix.Name `json:"subject,omitempty"`
	DNSNames   []string   `json:"dns_names,omitempty"`
	NotBefore  string     `json:"not_before,omitempty"`
	NotAfter   string     `json:"not_after,omitempty"`
	Error      string     `json:"error,omitempty"`
}

type Certs []*Cert

func NewCert(addr string) *Cert {
	host, port, err := splitHostPort(addr)
	if err != nil {
		return &Cert{DomainName: host, Error: err.Error()}
	}

	addr = joinDefaultPort(host, port)
	certs, err := serverCertificates(addr)
	if err != nil {
		return &Cert{DomainName: host, Error: err.Error()}
	}

	cert := certs[0]
	return &Cert{
		DomainName: host,
		Issuer:     &cert.Issuer,
		Subject:    &cert.Subject,
		DNSNames:   cert.DNSNames,
		NotBefore:  cert.NotBefore.In(time.Local).String(),
		NotAfter:   cert.NotAfter.In(time.Local).String(),
	}
}

func splitHostPort(addr string) (host, port string, err error) {
	addr = joinDefaultPort(addr, Port)
	host, port, err = net.SplitHostPort(addr)
	if err != nil {
		return "", "", err
	}
	return host, port, nil
}

func joinDefaultPort(host, port string) string {
	if strings.Contains(host, ":") {
		return host
	}
	return host + ":" + port
}

func serverCertificates(addr string) ([]*x509.Certificate, error) {
	dialer := &net.Dialer{Timeout: Timeout * time.Second,}

	con, err := tls.DialWithDialer(dialer, "tcp", addr, &tls.Config{})
	if err != nil {
		return nil, err
	}

	certs := con.ConnectionState().PeerCertificates
	err = con.Close()
	if err != nil {
		return nil, err
	}

	return certs, nil
}
