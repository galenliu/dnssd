package dnssd

import (
	"github.com/galenliu/chip/inet/IP"
	"github.com/galenliu/chip/inet/IPPacket"
	"github.com/galenliu/chip/inet/Interface"
	"github.com/galenliu/chip/inet/udp_endpoint"
	"github.com/galenliu/dnssd/core"
	"github.com/galenliu/dnssd/mdns"
	"github.com/galenliu/dnssd/old"
	"github.com/galenliu/dnssd/responders"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"net"
	"sync"
)

const kMdnsPort = 5353

type BroadcastAdvertiseType int

const (
	kStarted BroadcastAdvertiseType = iota
	kRemovingAll
)

// Advertiser 实现 PacketDelegate 和 ParserDelegate
type Advertiser struct {
	mResponseSender             *ResponseSender
	mCommissionableInstanceName string
	mIsInitialized              bool
	mQueryResponder             *responders.QueryResponder
	mCurrentSource              *IPPacket.Info
	mMessageId                  uint16
}

var insAdvertiser *Advertiser
var advertiserOnce sync.Once

func GetAdvertiserInstance() *Advertiser {
	advertiserOnce.Do(func() {
		insAdvertiser = defaultAdvertiser()
		insAdvertiser.mResponseSender = NewResponseSender()
		mdns.GlobalServer().SetQueryDelegate(insAdvertiser)
		insAdvertiser.mResponseSender.SetQueryResponder(insAdvertiser.mQueryResponder)
	})
	return insAdvertiser
}

func defaultAdvertiser() *Advertiser {
	return &Advertiser{}
}

func (a *Advertiser) Init(udpEndPointManager udp_endpoint.UDPEndpoint) error {
	mdns.GlobalServer().Shutdown()
	if a.mIsInitialized {
		//a.UpdateCommissionableInstanceName()
	}
	err := mdns.DefaultServer().StartServer(udpEndPointManager.Listeners(), kMdnsPort)
	if err != nil {
		return err
	}

	err = a.advertiseRecords(kStarted)
	if err != nil {
		return err
	}

	a.mIsInitialized = true

	return nil
}

func (a *Advertiser) RemoveServices() error {
	return nil
}

func (a *Advertiser) Shutdown() {
	mdns.GlobalServer().Shutdown()
}

func (a *Advertiser) advertiseRecords(t BroadcastAdvertiseType) error {

	responseConfiguration := &responders.ResponseConfiguration{}
	if t == kRemovingAll {
		responseConfiguration.SetTtlSecondsOverride(0)
	}

	for _, inter := range Interface.GetInterfaceIds() {
		for _, addr := range IP.GetAddress(inter) {
			c := new(dns.Client)
			var destAddress string
			c.Dialer = &net.Dialer{}
			if addr.Is6() {
				c.Net = "udp6"
				c.Dialer.LocalAddr = &net.UDPAddr{IP: addr.AsSlice()}
				destAddress = mdns.GetIpv6Into().String()
			}
			if addr.Is4() {
				c.Net = "udp"
				c.Dialer.LocalAddr = &net.UDPAddr{IP: addr.AsSlice()}
				destAddress = mdns.GetIpv4Into().String()
			}

			queryData := mdns.NewQueryData(dns.TypePTR, dns.ClassINET, false)
			queryData.SetIsInternalBroadcast(true)

			err := a.mResponseSender.BroadcastRecords(queryData, c, destAddress)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *Advertiser) FinalizeServiceUpdate() error {
	return nil
}

func (a *Advertiser) OnQuery(w dns.ResponseWriter, queryData *mdns.QueryData) {

	a.mMessageId = queryData.Id
	var defaultResponseConfiguration = &responders.ResponseConfiguration{}
	err := a.mResponseSender.Respond(w, queryData, defaultResponseConfiguration)
	if err != nil {
		log.Info("Failed to reply to query")
	}
}

func (a *Advertiser) FindOperationalAllocator(qname core.QNamePart) *old.QueryResponderAllocator {
	return nil
}

func (a *Advertiser) AddResponder(responder responders.RecordResponder) *responders.QueryResponderSettings {
	//TODO implement me
	panic("implement me")
}

func (a *Advertiser) RemoveRecords() error {
	//TODO implement me
	panic("implement me")
}

func (a *Advertiser) AdvertiseRecords() error {
	//TODO implement me
	panic("implement me")
}

//func (a *Advertiser) AdvertiseCommission(params server.CommissionAdvertisingParameters) error {
//
//	_ = a.advertiseRecords(kRemovingAll)
//
//	var queryResponder *responders.QueryResponder
//	var serviceType chip.ServiceType
//
//	if params.GetCommissionAdvertiseMode() == server.KCommissionableNode {
//		queryResponder = a.mQueryResponderAllocatorCommissionable
//		serviceType = chip.KCommissionableServiceName
//	} else {
//		queryResponder = a.mQueryResponderAllocatorCommissioner
//		serviceType = chip.KCommissionerServiceName
//	}
//
//	serviceName := core.ParseFullQName(serviceType.String(), chip.KCommissionProtocol, chip.KLocalDomain)
//	instanceName := core.ParseFullQName(a.GetCommissionableInstanceName(), serviceType.String(), chip.KCommissionProtocol, chip.KLocalDomain)
//
//	hostName := core.ParseFullQName(chip.KLocalDomain, a.GetCommissionableInstanceName())
//
//	if !queryResponder.AddResponder(responders.NewPtrResponder(serviceName, instanceName)).
//		SetReportAdditional(instanceName).
//		SetReportInServiceListing(true).
//		IsValid() {
//		return errors.New("failed to add service PTR record mDNS responder")
//	}
//
//	if !queryResponder.AddResponder(responders.NewSrvResponder(record.NewSrvResourceRecord(instanceName, hostName, params.GetPort()))).
//		SetReportAdditional(hostName).
//		IsValid() {
//		return errors.New("failed to add SRV record mDNS responder")
//	}
//
//	if !queryResponder.AddResponder(responders.NewIPv6Responder(hostName)).
//		IsValid() {
//		return errors.New("failed to add IPv6 mDNS responder")
//	}
//
//	if params.IsIPv4Enabled() {
//		if !queryResponder.AddResponder(responders.NewIPv4Responder(hostName)).
//			IsValid() {
//			return errors.New("failed to add IPv6 mDNS responder")
//		}
//	}
//
//	if params.GetVendorId() != nil {
//		name := fmt.Sprintf("_V%d", *params.GetVendorId())
//		vendorServiceName := core.ParseFullQName(name, chip.KSubtypeServiceNamePart, serviceType.String(), chip.KCommissionProtocol, chip.KLocalDomain)
//		if !queryResponder.AddResponder(responders.NewPtrResponder(vendorServiceName, instanceName)).
//			SetReportAdditional(instanceName).
//			SetReportInServiceListing(true).
//			IsValid() {
//			return errors.New("failed to add vendor PTR record mDNS responder")
//		}
//	}
//
//	if params.GetDeviceType() != nil {
//		name := fmt.Sprintf("_T%d", *params.GetDeviceType())
//		typeServiceName := core.ParseFullQName(name, chip.KSubtypeServiceNamePart, serviceType.String(), chip.KCommissionProtocol, chip.KLocalDomain)
//		if !queryResponder.AddResponder(responders.NewPtrResponder(typeServiceName, instanceName)).
//			SetReportAdditional(instanceName).
//			SetReportInServiceListing(true).
//			IsValid() {
//			return errors.New("failed to add vendor PTR record mDNS responder")
//		}
//	}
//
//	if params.GetCommissionAdvertiseMode() == server.KCommissionableNode {
//		// TODO
//	}
//
//	if !queryResponder.AddResponder(responders.NewTxtResponder(instanceName, a.GetCommissioningTxtEntries(params))).
//		SetReportAdditional(hostName).
//		IsValid() {
//		return errors.New("failed to add TXT record mDNS responder")
//	}
//
//	err := a.advertiseRecords(kStarted)
//	if err != nil {
//		return err
//	}
//	log.Infof("mDNS service published: %a", instanceName.String())
//	return nil
//}

//func (a *Advertiser) UpdateCommissionableInstanceName() {
//	a.mCommissionableInstanceName = strconv.FormatUint(rand.Uint64(), 16)
//	a.mCommissionableInstanceName = strings.ToUpper(a.mCommissionableInstanceName)
//}
//
//func (a *Advertiser) GetCommissionableInstanceName() string {
//	if a.mCommissionableInstanceName == "" {
//		a.mCommissionableInstanceName = strings.Replace(server.mac48Address(server.randHex()), ":", "", -1)
//	}
//	return a.mCommissionableInstanceName
//}
