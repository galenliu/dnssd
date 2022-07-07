package QName

import (
	"github.com/miekg/dns"
	"strings"
)

type FullQName string

func ParseFullQName(args ...string) FullQName {
	var s string
	for _, a := range args {
		if a == "" {
			continue
		}
		s = s + dns.Fqdn(strings.TrimSpace(a))
	}
	s = dns.Fqdn(s)
	return FullQName(s)
}

func (n FullQName) String() string {
	return dns.Fqdn(string(n))
}
