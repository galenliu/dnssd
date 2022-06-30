package responders

import (
	"github.com/galenliu/dnssd/core"
	"github.com/galenliu/dnssd/core/QClass"
	"github.com/galenliu/dnssd/core/QType"
)

type ReplyFilter interface {
	Accept(QType.T, QClass.T, *core.FullQName) bool
}
