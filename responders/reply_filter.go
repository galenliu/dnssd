package responders

import (
	"github.com/galenliu/dnssd/QClass"
	"github.com/galenliu/dnssd/QType"
)

type ReplyFilter interface {
	Accept(QType.T, QClass.T, string) bool
}
