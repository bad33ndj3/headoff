package system_profiler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExtractDevicesTestSuite struct {
	suite.Suite
}

func (s *ExtractDevicesTestSuite) TestExtractDevices() {
	input := `Bluetooth:

      Bluetooth Controller:
          DeviceAddress: F8:4D:89:8F:46:0D
          State: On
          Chipset: BCM_4387
          Discoverable: Off
          Firmware Version: 20.1.501.6743
          Product ID: 0x4A04
          Supported services: 0x382039 < HFP AVRCP A2DP HID Braille AACP GATT SerialPort >
          Transport: PCIe
          Vendor ID: 0x004C (Apple)
      Not IsConnected:
          Apple Watch van Casper:
              DeviceAddress: 60:95:BD:EC:F4:3A
              RSSI: -44
          Galaxy Note20 5G van Tessa:
              DeviceAddress: 6C:DD:BC:F2:4A:8A
              Vendor ID: 0x0075
              Product ID: 0x0100
              Firmware Version: 2.0.1
              Minor Type: MobilePhone
          iPhone van Casper:
              DeviceAddress: 4C:79:75:4E:F3:59
              RSSI: -61
          Keychron K3:
              DeviceAddress: DC:2C:26:06:2C:69
              Vendor ID: 0x05AC
              Product ID: 0x024F
              Firmware Version: 1.0.6
              Minor Type: Keyboard
          MDR-XB950N1:
              DeviceAddress: CC:98:8B:62:43:BC
              Vendor ID: 0x054C
              Product ID: 0x0BEA
              Firmware Version: 1.0.3
              Minor Type: Headset
          MX Master 3:
              DeviceAddress: E1:8A:2B:08:1B:A0
              Minor Type: Mouse
          Soundcore Liberty 3 Pro:
              DeviceAddress: AC:12:2F:F2:F9:6E
              Minor Type: Headset
          `
	expectedDevices := []*btDevice{
		{
			DeviceName:      "Apple Watch van Casper",
			DeviceAddress:   "60:95:BD:EC:F4:3A",
			DeviceMinorType: "",
			IsConnected:     false,
		},
		{
			DeviceName:      "Galaxy Note20 5G van Tessa",
			DeviceAddress:   "6C:DD:BC:F2:4A:8A",
			DeviceMinorType: "MobilePhone",
			IsConnected:     false,
		},
		{
			DeviceName:      "iPhone van Casper",
			DeviceAddress:   "4C:79:75:4E:F3:59",
			DeviceMinorType: "",
			IsConnected:     false,
		},
		{
			DeviceName:      "Keychron K3",
			DeviceAddress:   "DC:2C:26:06:2C:69",
			DeviceMinorType: "Keyboard",
			IsConnected:     false,
		},
		{
			DeviceName:      "MDR-XB950N1",
			DeviceAddress:   "CC:98:8B:62:43:BC",
			DeviceMinorType: "Headset",
			IsConnected:     false,
		},
		{
			DeviceName:      "MX Master 3",
			DeviceAddress:   "E1:8A:2B:08:1B:A0",
			DeviceMinorType: "Mouse",
			IsConnected:     false,
		},
		{
			DeviceName:      "Soundcore Liberty 3 Pro",
			DeviceAddress:   "AC:12:2F:F2:F9:6E",
			DeviceMinorType: "Headset",
			IsConnected:     false,
		},
	}

	profiler := &SystemProfiler{}

	devices := profiler.extractDevices(input)
	s.Len(devices, len(expectedDevices), "Devices should have the same length as expectedDevices")
	s.ElementsMatch(expectedDevices, devices, "Devices should be extracted correctly")
}

func TestExtractDevices(t *testing.T) {
	suite.Run(t, new(ExtractDevicesTestSuite))
}
