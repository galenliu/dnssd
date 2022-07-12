package dnssd

import (
	"net/netip"
)

type Config struct {
	Name   string
	Type   string
	Domain string
	Host   string
	Text   map[string]string
	IPs    []netip.Addr
	Port   int
	Ifaces []string
}
