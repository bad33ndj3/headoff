package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExtractDevicesTestSuite struct {
	suite.Suite
}

func TestExtractDevices(t *testing.T) {
	suite.Run(t, new(ExtractDevicesTestSuite))
}

func (s *ExtractDevicesTestSuite) TestExtractDevices() {
	input := `Bluetooth:

      Bluetooth Controller:
          Address: F8:4D:89:8F:46:0D
          State: On
          Chipset: BCM_4387
          Discoverable: Off
          Firmware Version: 20.1.501.6743
          Product ID: 0x4A04
          Supported services: 0x382039 < HFP AVRCP A2DP HID Braille AACP GATT SerialPort >
          Transport: PCIe
          Vendor ID: 0x004C (Apple)
      Not Connected:
          Apple Watch van Casper:
              Address: 60:95:BD:EC:F4:3A
              RSSI: -44
          Galaxy Note20 5G van Tessa:
              Address: 6C:DD:BC:F2:4A:8A
              Vendor ID: 0x0075
              Product ID: 0x0100
              Firmware Version: 2.0.1
              Minor Type: MobilePhone
          iPhone van Casper:
              Address: 4C:79:75:4E:F3:59
              RSSI: -61
          Keychron K3:
              Address: DC:2C:26:06:2C:69
              Vendor ID: 0x05AC
              Product ID: 0x024F
              Firmware Version: 1.0.6
              Minor Type: Keyboard
          MDR-XB950N1:
              Address: CC:98:8B:62:43:BC
              Vendor ID: 0x054C
              Product ID: 0x0BEA
              Firmware Version: 1.0.3
              Minor Type: Headset
          MX Master 3:
              Address: E1:8A:2B:08:1B:A0
              Minor Type: Mouse
          Soundcore Liberty 3 Pro:
              Address: AC:12:2F:F2:F9:6E
              Minor Type: Headset
          `
	expectedDevices := []Device{
		{
			Name:      "Apple Watch van Casper",
			Address:   "60:95:BD:EC:F4:3A",
			MinorType: "",
			Connected: false,
		},
		{
			Name:      "Galaxy Note20 5G van Tessa",
			Address:   "6C:DD:BC:F2:4A:8A",
			MinorType: "MobilePhone",
			Connected: false,
		},
		{
			Name:      "iPhone van Casper",
			Address:   "4C:79:75:4E:F3:59",
			MinorType: "",
			Connected: false,
		},
		{
			Name:      "Keychron K3",
			Address:   "DC:2C:26:06:2C:69",
			MinorType: "Keyboard",
			Connected: false,
		},
		{
			Name:      "MDR-XB950N1",
			Address:   "CC:98:8B:62:43:BC",
			MinorType: "Headset",
			Connected: false,
		},
		{
			Name:      "MX Master 3",
			Address:   "E1:8A:2B:08:1B:A0",
			MinorType: "Mouse",
			Connected: false,
		},
		{
			Name:      "Soundcore Liberty 3 Pro",
			Address:   "AC:12:2F:F2:F9:6E",
			MinorType: "Headset",
			Connected: false,
		},
	}

	devices := extractDevices(input)
	s.Len(devices, len(expectedDevices), "Devices should have the same length as expectedDevices")
	s.ElementsMatch(expectedDevices, devices, "Devices should be extracted correctly")
}
