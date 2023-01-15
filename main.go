package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	exists := checkBlueutilAvailability()
	if !exists {
		log.Fatal("blueutil is not found in PATH. try \"brew install blueutil\"")
	}
	output := getBluetoothOutput()
	devices := extractDevices(output)
	for i := range devices {
		if devices[i].MinorType == "Headset" && devices[i].Connected == true {
			err := disconnectDevice(devices[i].Name)
			if err != nil {
				fmt.Println("failed disconnection", err)
			}
		}
	}
}

func checkBlueutilAvailability() bool {
	_, err := exec.LookPath("blueutil")
	if err != nil {
		return false
	}
	return true
}
func getBluetoothOutput() string {
	// todo: this output could also come from blueutil but would be in a different format.
	cmd := exec.Command("system_profiler", "SPBluetoothDataType")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return out.String()
}
func disconnectDevice(address string) error {
	// todo: address is called ID in blueutil so that might be an incorrect assumption.
	return exec.Command("blueutil", "--disconnect", address).Run()
}
