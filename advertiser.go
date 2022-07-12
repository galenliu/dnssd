package dnssd

import (
	"github.com/galenliu/dnssd/mdns"
	"github.com/galenliu/dnssd/responders"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"net/netip"
)

type BroadcastAdvertiseType int

const (
	KStarted BroadcastAdvertiseType = iota
	KRemovingAll
)

// Advertiser 实现 PacketDelegate 和 ParserDelegate
type Advertiser struct {
	mResponseSender *ResponseSender
	mIsInitialized  bool
	mMessageId      uint16
	mServer         *mdns.Server
}

func (a Advertiser) Init(adders []netip.Addr, port uint16) (*Advertiser, error) {

	a.mServer = new(mdns.Server).Init()
	a.mResponseSender = new(ResponseSender).Int()
	a.mServer.SetQueryDelegate(&a)

	a.mServer = new(mdns.Server)
	a.mServer.Init()
	a.mServer.Shutdown()

	a.mResponseSender.SetServer(a.mServer)

	err := a.mServer.StartServer(adders, port)
	if err != nil {
		return nil, err
	}

	err = a.AdvertiseRecords(KStarted)
	if err != nil {
		return nil, err
	}

	a.mIsInitialized = true

	return &a, nil
}

func (a *Advertiser) RemoveServices() error {
	return nil
}

func (a *Advertiser) Shutdown() {
	a.mServer.Shutdown()
}

func (a *Advertiser) AdvertiseRecords(t BroadcastAdvertiseType) error {

	responseConfiguration := &responders.ResponseConfiguration{}
	if t == KRemovingAll {
		responseConfiguration.SetTtlSecondsOverride(0)
	}
	queryData := mdns.NewQueryData(dns.TypePTR, dns.ClassINET, false)
	queryData.SetIsInternalBroadcast(true)

	err := a.mResponseSender.Respond(nil, queryData, responseConfiguration)
	if err != nil {
		return err
	}
	return nil
}

func (a *Advertiser) FinalizeServiceUpdate() error {
	return nil
}

func (a *Advertiser) OnQuery(w dns.ResponseWriter, queryData *mdns.QueryData) {
	log.Infof("advertiser on query:\t\n %s", queryData.String())
	a.mMessageId = queryData.Id
	var defaultResponseConfiguration = &responders.ResponseConfiguration{}
	err := a.mResponseSender.Respond(w, queryData, defaultResponseConfiguration)
	if err != nil {
		log.Info("Failed to reply to query")
	}
}

func (a *Advertiser) AddResponder(responder responders.RecordResponder) *responders.QueryResponderSettings {
	return a.mResponseSender.mResponders.AddResponder(responder)
}

func (a *Advertiser) RemoveRecords() error {
	//TODO implement me
	return nil
}
