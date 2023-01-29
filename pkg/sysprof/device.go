package sysprof

import (
	"github.com/bad33ndj3/headoff/pkg/device"
)

// validate btDevice implements device.Info
var _ device.Info = (*btDevice)(nil)

type btDevice struct {
	DeviceName      string
	DeviceAddress   string
	DeviceMinorType string
	IsConnected     bool
}

func (b *btDevice) Name() string {
	return b.DeviceName
}

func (b *btDevice) Address() string {
	return b.DeviceAddress
}

func (b *btDevice) MinorType() string {
	return b.DeviceMinorType
}

func (b *btDevice) Connected() bool {
	return b.IsConnected
}
