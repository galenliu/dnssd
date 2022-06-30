package dnssd

import (
	"github.com/galenliu/chip/inet/udp_endpoint"
	"sync"
)

type Resolver struct {
}

func (r Resolver) Init(manager udp_endpoint.UDPEndpoint) {

}

var insResolver *Resolver
var onceResolver sync.Once

func ResolverInstance() *Resolver {
	onceResolver.Do(func() {
		insResolver = newResolver()
	})
	return insResolver
}

func newResolver() *Resolver {
	return &Resolver{}
}
