package mdns

import (
	"github.com/miekg/dns"
)

const kSizeBytes uint8 = 12
const kMaxValueSize = 63

type QueryData struct {
	*dns.Msg
	mAnswerViaUnicast    bool
	mIsInternalBroadcast bool
}

func NewQueryData(qType, class uint16, unicast bool) *QueryData {
	data := &QueryData{
		Msg: &dns.Msg{
			Question: make([]dns.Question, 1),
		},
		mAnswerViaUnicast:    unicast,
		mIsInternalBroadcast: false,
	}
	data.Question[0] = dns.Question{
		Name:   "",
		Qtype:  qType,
		Qclass: class,
	}
	return data
}

//func NewQueryData(qType QType.T, kClass QClass.T, unicast bool) *QueryData {
//	return &QueryData{
//		Msg: dns.Msg{
//			MsgHdr:   dns.MsgHdr{},
//			Compress: false,
//			Question: nil,
//			Answer:   nil,
//			Ns:       nil,
//			Extra:    nil,
//		},
//		mType:                qType,
//		mClass:               kClass,
//		mAnswerViaUnicast:    unicast,
//		mIsInternalBroadcast: false,
//	}
//}

//func (q *QueryData) Parse(validData *core.BytesRange, start, end uint8) bool {
//	// Structure is:
//	//    QNAME
//	//    TYPE
//	//    CLASS (plus a flag for unicast)
//	if validData.Size() < end {
//		return false
//	}
//	data := validData.Bytes()[start:end]
//	q.mType = validData.Get16At(start)
//	return true
//}

func (q *QueryData) SetIsInternalBroadcast(isInternalBroadcast bool) {
	q.mIsInternalBroadcast = isInternalBroadcast
}

func (q *QueryData) GetType() uint16 {
	return q.Msg.Question[0].Qtype
}

func (q *QueryData) GetClass() uint16 {
	return q.Msg.Question[0].Qclass
}

func (q *QueryData) IsInternalBroadcast() bool {
	return q.mIsInternalBroadcast
}

func (q *QueryData) RequestedUnicastAnswer() bool {
	return q.mAnswerViaUnicast
}

func (q *QueryData) GetName() string {
	return q.Question[0].Name
}

//func ParsePacket(packetData *core.BytesRange, delegate ParserDelegate) bool {
//
//	if packetData.Size() < core.KSizeBytes {
//		return false
//	}
//	var header = &core.ConstHeaderRef{
//		ID:      packetData.Get16At(core.KMessageIdOffset),
//		FLAGS:   packetData.Get16At(core.KFlagsOffset),
//		QDCOUNT: packetData.Get16At(core.KQueryCountOffset),
//		ANCOUNT: packetData.Get16At(core.KAnswerCountOffset),
//		NSCOUNT: packetData.Get16At(core.KAuthorityCountOffset),
//		ARCOUNT: packetData.Get16At(core.KAdditionalCountOffset),
//	}
//
//	if !header.IsValidMdns() {
//		return false
//	}
//
//	// set messageId
//	delegate.OnHeader(header)
//	{
//		queryDataList := packetData.ParseQueryData()
//		for _, queryData := range queryDataList {
//			delegate.OnQuery(queryData)
//		}
//
//		resourceDataList := packetData.ParseQueryResourceData()
//		for _, resourceData := range resourceDataList {
//			delegate.OnResource(ResourceType.Answer, resourceData)
//		}
//
//		resourceAdditionalList := packetData.ParseQueryResourceAdditional()
//		for _, resourceData := range resourceAdditionalList {
//			delegate.OnResource(ResourceType.Additional, resourceData)
//		}
//
//	}
//
//	return true
//}
