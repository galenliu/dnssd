package responders

import "time"

type QueryResponderRecordFilter struct {
	mIncludeAdditionalRepliesOnly bool
	mReplyFilter                  ReplyFilter
	mIncludeOnlyMulticastBefore   time.Time
}

func (f *QueryResponderRecordFilter) SetReplyFilter(filter ReplyFilter) *QueryResponderRecordFilter {
	f.mReplyFilter = filter
	return f
}

func (f *QueryResponderRecordFilter) SetIncludeOnlyMulticastBeforeMS(t time.Time) {
	f.mIncludeOnlyMulticastBefore = t
}

func (f *QueryResponderRecordFilter) Accept(record *QueryResponderInfo) bool {
	if record.Responder == nil {
		return false
	}
	if f.mIncludeAdditionalRepliesOnly && !record.reportNowAsAdditional {
		return false
	}

	if f.mIncludeOnlyMulticastBefore.Before(time.Now()) && record.LastMulticastTime.Before(f.mIncludeOnlyMulticastBefore) {
		return false
	}

	if f.mReplyFilter != nil && !f.mReplyFilter.Accept(record.Responder.GetType(), record.Responder.GetClass(), record.Responder.GetName()) {
		return false
	}

	return true
}

func (f *QueryResponderRecordFilter) SetIncludeAdditionalRepliesOnly(b bool) *QueryResponderRecordFilter {
	f.mIncludeAdditionalRepliesOnly = b
	return f
}
