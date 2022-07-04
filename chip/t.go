package chip

type ServiceType string
type Protocol string

const (
	KSubtypeServiceNamePart                = "_sub"
	KCommissionableServiceName ServiceType = "_matterc"
	KCommissionerServiceName   ServiceType = "_matterd"
	kOperationalServiceName    ServiceType = "_matter"
	KCommissionProtocol                    = "_udp"
	KLocalDomain                           = "local"
	kOperationalProtocol                   = "_tcp"
)

func (s ServiceType) String() string {
	return string(s)
}

func (s Protocol) String() string {
	return string(s)
}
