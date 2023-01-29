package blueutil

import (
	"fmt"
	"os/exec"

	"github.com/bad33ndj3/headoff/pkg/device"
)

const blueutilCmd = "blueutil"

// BlueUtil is a wrapper for the blueutil command.
type BlueUtil struct {
}

// NewBlueUtil returns a new BlueUtil.
func NewBlueUtil() (*BlueUtil, error) {
	return &BlueUtil{}, checkBlueutilAvailability()
}

// Disconnect disconnects the given device.
func (b *BlueUtil) Disconnect(device device.Info) error {
	return exec.Command(blueutilCmd, "--disconnect", device.Address()).Run()
}

// checkBlueutilAvailability checks if blueutil is available in PATH.
func checkBlueutilAvailability() error {
	_, err := exec.LookPath("blueutil")
	if err != nil {
		return fmt.Errorf("blueutil is not found in PATH. try \"brew install blueutil\": %w", err)
	}
	return nil
}
