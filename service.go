package dnssd

import (
	"fmt"
	"github.com/galenliu/dnssd/chip"
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

func NewConf(instanceName string) *Config {
	t := fmt.Sprintf("%s.%s", chip.KCommissionableServiceName, chip.KCommissionProtocol)
	return &Config{
		Name:   instanceName,
		Type:   t,
		Domain: chip.KLocalDomain,
		Host:   "",
		Text:   nil,
		IPs:    nil,
		Port:   0,
		Ifaces: nil,
	}
}
