package device

// MinorTypeHeadset is the minor type for a headset
const MinorTypeHeadset = "Headset"

// Info is an interface for a device
type Info interface {
	Name() string
	Address() string
	MinorType() string
	Connected() bool
}

// Lister is an interface for listing devices
type Lister interface {
	List() ([]Info, error)
}

// Disconnecter is an interface for disconnecting devices
type Disconnecter interface {
	Disconnect(device Info) error
}
