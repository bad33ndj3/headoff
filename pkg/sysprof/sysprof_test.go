package sysprof

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
	  Apple Watch van Casper:
		  DeviceAddress: 60:95:BD:EC:F4:3A
		  RSSI: -44
	  MDR-XB950N1:
		  DeviceAddress: CC:98:8B:62:43:BC
		  Vendor ID: 0x054C
		  Product ID: 0x0BEA
		  Firmware Version: 1.0.3
		  Minor Type: Headset
      Not IsConnected:
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
			IsConnected:     true,
		},
		{
			DeviceName:      "MDR-XB950N1",
			DeviceAddress:   "CC:98:8B:62:43:BC",
			DeviceMinorType: "Headset",
			IsConnected:     true,
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
