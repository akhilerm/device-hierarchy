package main

import (
	"fmt"
	"github.com/akhilerm/device-topology/hierarchy"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Device Name not given")
		os.Exit(0)
	}
	deviceName := os.Args[1]
	dev := hierarchy.Device{deviceName}
	dep, err := dev.GetDependents()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println(dep)
	}
}
