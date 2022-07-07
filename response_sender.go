package dnssd

import (
	"fmt"
	"github.com/galenliu/dnssd/ResourceType"
	"github.com/galenliu/dnssd/mdns"
	"github.com/galenliu/dnssd/responders"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"

	"time"
)

const kMdnsStandardPort = 5353

type Server interface {
	Broadcast(msg *dns.Msg) error
}

// ResponseSender 实现 ResponderDelegate接口
type ResponseSender struct {
	mSendState  *ResponseSendingState
	mResponders *responders.QueryResponder
	mServer     Server
}

func (r *ResponseSender) Int() *ResponseSender {
	r.mResponders = &responders.QueryResponder{
		ResponderInfos: make([]*responders.QueryResponderInfo, 0),
	}
	r.mResponders = &responders.QueryResponder{ResponderInfos: make([]*responders.QueryResponderInfo, 0)}
	r.mSendState = &ResponseSendingState{
		mQuery:        nil,
		mMessageId:    0,
		mResourceType: 0,
		mSendError:    nil,
	}
	return r
}

func (r *ResponseSender) BroadcastRecords(query *mdns.QueryData, client *dns.Client, address string) error {

	msg, err := r.OnQuery(query)
	if err != nil {
		return err
	}
	if client == nil {
		return fmt.Errorf("dns clint empty")
	}
	log.Printf("mDns broadcast Msg: %s", msg.String())
	log.Printf("mDns broadcast Net: %s ,Local Addr: %s Des Addr: %s", client.Net, client.Dialer.LocalAddr.String(), address)
	_, _, err = client.Exchange(msg, address)
	if err != nil {
		return err
	}
	return nil
}

func (r *ResponseSender) Respond(w dns.ResponseWriter, query *mdns.QueryData, configuration *responders.ResponseConfiguration) error {

	r.mSendState.Reset(query.Id, query)
	if w != nil {
		r.mSendState.SetSourceAddr(w.RemoteAddr().String())
		r.mSendState.SetDestAddr(w.LocalAddr().String())
	}
	log.Infof(r.mSendState.mSrcAddress.String())

	r.mResponders.ResetAdditionals()

	log.Infof("Query Message Respond: \t\n %s", query.Msg.String())
	msg := &dns.Msg{}
	msg.SetReply(query.Msg)
	// send all 'Answer' replies
	{
		queryReplyFilter := NewQueryReplyFilter(query)
		responseFilter := responders.QueryResponderRecordFilter{}
		responseFilter.SetReplyFilter(queryReplyFilter)

		if !r.mSendState.SendUnicast() {
			responseFilter.SetIncludeOnlyMulticastBeforeMS(time.Now())
		}
		for _, info := range r.mResponders.ResponderInfos {
			if !responseFilter.Accept(info) {
				continue
			}
			msg.Answer = append(msg.Answer, info.Responder.ResourceRecord())
			if err := r.mSendState.GetError(); err != nil {
				return err
			}
			if !r.mSendState.SendUnicast() {
				info.LastMulticastTime = time.Now()
			}
		}
	}
	// send all 'Additional' replies
	{
		r.mSendState.SetResourceType(ResourceType.Additional)
		queryReplyFilter := NewQueryReplyFilter(query)
		queryReplyFilter.SetIgnoreNameMatch(true).
			SetSendingAdditionalItems(true)

		responseFilter := responders.QueryResponderRecordFilter{}
		responseFilter.SetReplyFilter(queryReplyFilter).
			SetIncludeAdditionalRepliesOnly(true)

		for _, info := range r.mResponders.ResponderInfos {
			if !responseFilter.Accept(info) {
				continue
			}
			msg.Answer = append(msg.Answer, info.Responder.ResourceRecord())
			if err := r.mSendState.GetError(); err != nil {
				return err
			}
		}

	}

	if w == nil {
		return r.mServer.Broadcast(msg)
	}
	return w.WriteMsg(msg)
}

func (r *ResponseSender) OnQuery(query *mdns.QueryData) (*dns.Msg, error) {

	log.Infof("Query Message:\t\n %s", query.Msg.String())
	msg := &dns.Msg{}
	msg.SetReply(query.Msg)
	// send all 'Answer' replies
	{
		queryReplyFilter := NewQueryReplyFilter(query)
		responseFilter := responders.QueryResponderRecordFilter{}
		responseFilter.SetReplyFilter(queryReplyFilter)

		if !r.mSendState.SendUnicast() {
			responseFilter.SetIncludeOnlyMulticastBeforeMS(time.Now())
		}
		for _, info := range r.mResponders.ResponderInfos {
			if !responseFilter.Accept(info) {
				continue
			}
			msg.Answer = append(msg.Answer, info.Responder.ResourceRecord())
			if err := r.mSendState.GetError(); err != nil {
				return nil, err
			}
			if !r.mSendState.SendUnicast() {
				info.LastMulticastTime = time.Now()
			}
		}
	}
	// send all 'Additional' replies
	{
		r.mSendState.SetResourceType(ResourceType.Additional)
		queryReplyFilter := NewQueryReplyFilter(query)
		queryReplyFilter.SetIgnoreNameMatch(true).
			SetSendingAdditionalItems(true)

		responseFilter := responders.QueryResponderRecordFilter{}
		responseFilter.SetReplyFilter(queryReplyFilter).
			SetIncludeAdditionalRepliesOnly(true)

		for _, info := range r.mResponders.ResponderInfos {
			if !responseFilter.Accept(info) {
				continue
			}
			msg.Answer = append(msg.Answer, info.Responder.ResourceRecord())
			if err := r.mSendState.GetError(); err != nil {
				return nil, err
			}
		}

	}
	return msg, nil
}

func (r *ResponseSender) SetServer(server *mdns.Server) {
	r.mServer = server
}

//func (r *ResponseSender) FlushReply() error {
//	if r.mResponseBuilder.HasResponseRecords() {
//		if r.mSendState.SendUnicast() {
//			log.Info("Discovery: Directly sending mDns reply to peer %s on port %d", r.mSendState.GetSourcePort())
//			err := r.mServer.DirectSend(
//				r.mResponseBuilder.ReleasePacket(),
//				r.mSendState.GetSourceAddress(),
//				r.mSendState.GetSourcePort(),
//				r.mSendState.GetSourceInterfaceId())
//			if err != nil {
//				return err
//			}
//		} else {
//			err := r.mServer.BroadcastSend(
//				r.mResponseBuilder.ReleasePacket(),
//				kMdnsStandardPort,
//				r.mSendState.GetSourceInterfaceId(),
//				r.mSendState.GetSourceAddress())
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
