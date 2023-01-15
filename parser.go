package main

import (
	"regexp"
	"strings"
)

func extractDevices(input string) []Device {
	// Split the input string by newline
	lines := strings.Split(input, "\n")
	var devices []Device
	var notConnected bool

	// Regex to match the device properties
	nameRegex := regexp.MustCompile(`^([A-Za-z0-9\s]+):$`)
	addressRegex := regexp.MustCompile(`Address: ([A-Za-z0-9:]+)`)
	minorTypeRegex := regexp.MustCompile(`Minor Type: ([A-Za-z]+)`)
	notConnectedRegex := regexp.MustCompile(`Not Connected:`)

	// Iterate through the lines
	var device Device
	for _, line := range lines {
		// Trim leading/trailing whitespace
		line = strings.TrimSpace(line)

		if notConnectedRegex.MatchString(line) {
			notConnected = true
		}
		// Check if the line starts with a device name
		if name := nameRegex.FindStringSubmatch(line); name != nil {
			// If a device is already being processed, add it to the devices slice
			if device.Name != "" {
				devices = append(devices, device)
			}

			// Start processing a new device
			device = Device{Name: name[1], Connected: !notConnected}
		} else if address := addressRegex.FindStringSubmatch(line); address != nil {
			device.Address = address[1]
		} else if minorType := minorTypeRegex.FindStringSubmatch(line); minorType != nil {
			device.MinorType = minorType[1]
		}
	}

	// Append the last device to the devices slice
	if device.Name != "" {
		devices = append(devices, device)
	}

	return filterWrongDevices(devices)
}

// filterWrongDevices is a quick fix for filtering some default devices that are not really devices.
func filterWrongDevices(devices []Device) []Device {
	result := make([]Device, 0, len(devices))
	for _, device := range devices {
		if device.Name != "Not Connected" && device.Name != "Bluetooth" {
			result = append(result, device)
		}
	}

	return result
}
