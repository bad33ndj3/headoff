package main

import (
	"fmt"
	"log"

	"github.com/bad33ndj3/headoff/pkg/blueutil"
	"github.com/bad33ndj3/headoff/pkg/device"
	"github.com/bad33ndj3/headoff/pkg/sysprof"
)

func main() {
	systemProfiler := &sysprof.SystemProfiler{}

	var blueUtil device.Disconnecter
	blueUtil, err := blueutil.NewBlueUtil()
	if err != nil {
		log.Fatal(err)
	}

	devices, err := systemProfiler.List()
	if err != nil {
		log.Fatal(err)
	}

	for i := range devices {
		if devices[i].MinorType() == device.MinorTypeHeadset && devices[i].Connected() {
			err := blueUtil.Disconnect(devices[i])
			if err != nil {
				fmt.Println("failed disconnection", err)
			}
		}
	}
}
