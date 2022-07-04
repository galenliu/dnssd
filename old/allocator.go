package old

import (
	"github.com/galenliu/dnssd/responders"
)

type QueryResponderAllocator struct {
	kMaxRecords     int
	mQueryResponder *responders.QueryResponder
}

func (q QueryResponderAllocator) Clear() {
	return
}

//func (q *QueryResponderAllocator) AllocateHostQName(mac string, kLocalDomain string) core.FullQName {
//	return &core.FullQName{
//		Domain: kLocalDomain,
//	}
//}
//
//func (q *QueryResponderAllocator) AllocateQName(serviceType core.ServiceType, chip.KCommissionProtocol core.Protocol, kLocalDomain string, instanceName ...string) core.FullQName {
//	fName := &core.FullQName{
//		ServerType: serviceType,
//		Protocol:   chip.KCommissionProtocol,
//		Domain:     kLocalDomain,
//	}
//	if instanceName != nil {
//		fName.Instance = instanceName[0]
//	}
//	return fName
//}

func (q *QueryResponderAllocator) GetQueryResponder() *responders.QueryResponder {
	return q.mQueryResponder
}

func (q *QueryResponderAllocator) AllocateQNameSpace(size uint) {

}

func (q *QueryResponderAllocator) AddResponder(res responders.RecordResponder) *responders.QueryResponderSettings {
	return q.mQueryResponder.AddResponder(res)
}
