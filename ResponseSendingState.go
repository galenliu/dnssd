package dnssd

import (
	"github.com/galenliu/chip/inet/IP"
	"github.com/galenliu/chip/inet/IPPacket"
	"github.com/galenliu/chip/inet/Interface"
	"github.com/galenliu/dnssd/core/ResourceType"
	"github.com/galenliu/dnssd/mdns"
)

type ResponseSendingState struct {
	mQuery        *mdns.QueryData
	mSource       *IPPacket.Info
	mMessageId    uint16
	mResourceType ResourceType.T
	mSendError    error
}

func (state *ResponseSendingState) Reset(messageId uint16, query *mdns.QueryData) {
	state.mMessageId = messageId
	state.mQuery = query
}

func (s *ResponseSendingState) SendUnicast() bool {
	if s.mQuery == nil {
		return false
	}
	return s.mQuery.RequestedUnicastAnswer() || s.mSource.SrcPort != kMdnsStandardPort
}

func (state *ResponseSendingState) GetError() error {
	return state.mSendError
}

func (state *ResponseSendingState) GetSourcePort() int {
	return state.mSource.SrcPort
}

func (state *ResponseSendingState) GetSourceAddress() IP.Address {
	return state.mSource.SrcAddress
}

func (state *ResponseSendingState) GetSourceInterfaceId() Interface.Id {
	return state.mSource.InterfaceId
}

func (state *ResponseSendingState) SetResourceType(additional ResourceType.T) {
	state.mResourceType = additional
}
