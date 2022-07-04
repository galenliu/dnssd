package responders

import "github.com/galenliu/dnssd/core"

type QueryResponderSettings struct {
	mInfo *QueryResponderInfo
}

func (s *QueryResponderSettings) SetReportAdditional(qName core.FullQName) *QueryResponderSettings {
	if s.IsValid() {
		s.mInfo.alsoReportAdditionalQName = true
		s.mInfo.additionalQName = qName
	}
	return s
}

func (s *QueryResponderSettings) IsValid() bool {
	return s.mInfo != nil
}

func (s *QueryResponderSettings) SetReportInServiceListing(reportService bool) *QueryResponderSettings {
	if s.IsValid() {
		s.mInfo.reportService = reportService
	}
	return s
}
