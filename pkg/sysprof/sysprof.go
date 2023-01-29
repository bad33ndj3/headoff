// Package sysprof contains an implementation of device.Lister for the macOS system_profiler command.
package sysprof

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/bad33ndj3/headoff/pkg/device"
)

// SystemProfiler is a wrapper for the system_profiler command.
type SystemProfiler struct {
}

// List returns a list of devices.
func (s *SystemProfiler) List() ([]device.Info, error) {
	output := s.bluetoothInfo()
	return s.extractDevices(output), nil
}

// extractDevices extracts the devices from system_profiler output.
func (s *SystemProfiler) extractDevices(input string) []device.Info {
	lines := strings.Split(input, "\n")
	var devices []*btDevice
	var notConnected bool

	// Regex to match the btDevice properties
	nameRegex := regexp.MustCompile(`^([A-Za-z0-9\s]+):$`)
	addressRegex := regexp.MustCompile(`DeviceAddress: ([A-Za-z0-9:]+)`)
	minorTypeRegex := regexp.MustCompile(`Minor Type: ([A-Za-z]+)`)
	notConnectedRegex := regexp.MustCompile(`Not IsConnected:`)

	// Iterate through the lines
	var d *btDevice
	for _, line := range lines {
		// Trim leading/trailing whitespace
		line = strings.TrimSpace(line)

		if notConnectedRegex.MatchString(line) {
			notConnected = true
		}
		// Check if the line starts with a btDevice name
		if name := nameRegex.FindStringSubmatch(line); name != nil {
			// If a btDevice is already being processed, add it to the devices slice
			if d != nil && d.Name() != "" {
				devices = append(devices, d)
			}

			// Start processing a new btDevice
			d = &btDevice{DeviceName: name[1], IsConnected: !notConnected}
		} else if address := addressRegex.FindStringSubmatch(line); address != nil {
			d.DeviceAddress = address[1]
		} else if minorType := minorTypeRegex.FindStringSubmatch(line); minorType != nil {
			d.DeviceMinorType = minorType[1]
		}
	}

	// Append the last btDevice to the devices slice
	if d.Name() != "" {
		devices = append(devices, d)
	}

	return s.filterWrongDevices(devices)
}

// filterWrongDevices is a quick fix for filtering some default devices that are not really devices.
func (s *SystemProfiler) filterWrongDevices(devices []*btDevice) []device.Info {
	result := make([]device.Info, 0, len(devices))
	for _, d := range devices {
		if d.Name() != "Not Connected" && d.Name() != "Bluetooth" {
			result = append(result, d)
		}
	}

	return result
}

// bluetoothInfo returns the output of the system_profiler command.
func (s *SystemProfiler) bluetoothInfo() string {
	cmd := exec.Command("system_profiler", "SPBluetoothDataType")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return out.String()
}
