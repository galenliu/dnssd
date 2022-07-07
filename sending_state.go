package dnssd

import (
	"github.com/galenliu/dnssd/ResourceType"
	"github.com/galenliu/dnssd/mdns"
	"net/netip"
)

type ResponseSendingState struct {
	mQuery        *mdns.QueryData
	mSrcAddress   netip.AddrPort
	mDestAddress  netip.AddrPort
	mMessageId    uint16
	mResourceType ResourceType.T
	mSendError    error
}

func (s *ResponseSendingState) Reset(messageId uint16, query *mdns.QueryData) {
	s.mMessageId = messageId
	s.mQuery = query
}

func (s *ResponseSendingState) SendUnicast() bool {
	if s.mQuery == nil {
		return false
	}
	return s.mQuery.RequestedUnicastAnswer() || s.mSrcAddress.Port() != kMdnsStandardPort
}

func (s *ResponseSendingState) GetError() error {
	return s.mSendError
}

func (s *ResponseSendingState) GetSourcePort() uint16 {
	return s.mSrcAddress.Port()
}

func (s *ResponseSendingState) GetSourceAddress() netip.Addr {
	return s.mSrcAddress.Addr()
}

func (s *ResponseSendingState) SetResourceType(additional ResourceType.T) {
	s.mResourceType = additional
}

func (s *ResponseSendingState) SetSourceAddr(addr string) {
	a, err := netip.ParseAddrPort(addr)
	if err != nil {
		return
	}
	s.mSrcAddress = a
}

func (s *ResponseSendingState) SetDestAddr(addr string) {
	a, err := netip.ParseAddrPort(addr)
	if err != nil {
		return
	}
	s.mDestAddress = a
}
