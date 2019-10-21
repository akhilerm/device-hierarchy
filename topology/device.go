package topology

import (
	"fmt"
	"path/filepath"
	"strings"
)

var sysFSDirectoryPath = "/sys/"

type Device struct {
	// Path of the blockdevice. eg: /dev/sda, /dev/dm-0
	Path string
}

type DependentDevices struct {
	// Holders is the slice of block-devices that are held by a given
	// blockdevice
	Holders []string

	// Slaves is the slice of blockdevices to which the given blockdevice
	// is a slave
	Slaves []string
}

func (d DependentDevices) String() string {
	holderHeader := "Holders : "
	holders := fmt.Sprintf("%v", d.Holders)
	slaveHeader := "Slaves : "
	slaves := fmt.Sprintf("%v", d.Slaves)
	return holderHeader + holders + "\n" +
		slaveHeader + slaves
}

func (d *Device) GetDependents() (DependentDevices, error) {
	dependents := DependentDevices{}
	blockDeviceName := strings.Replace(d.Path, "/dev/", "", 1)

	blockDeviceSymLink := sysFSDirectoryPath + "class/block/" + blockDeviceName

	sysPath, err := filepath.EvalSymlinks(blockDeviceSymLink)
	if err != nil {
		return dependents, err
	}
	blockDeviceSysPath := deviceSysPath{
		SysPath:    sysPath,
		DeviceName: blockDeviceName,
	}

	// parent device
	if parent, ok := blockDeviceSysPath.getParent(); ok {
		dependents.Slaves = append(dependents.Slaves, parent)
	}

	// get the partitions
	if partitions, ok := blockDeviceSysPath.getPartitions(); ok {
		dependents.Holders = append(dependents.Holders, partitions...)
	}

	// get the holder devices
	if holders, ok := blockDeviceSysPath.getHolders(); ok {
		dependents.Holders = append(dependents.Holders, holders...)
	}

	// get the slaves
	if slaves, ok := blockDeviceSysPath.getSlaves(); ok {
		dependents.Slaves = append(dependents.Slaves, slaves...)
	}

	// adding /dev prefix
	for i, _ := range dependents.Slaves {
		dependents.Slaves[i] = "/dev/" + dependents.Slaves[i]
	}

	// adding /dev prefix
	for i, _ := range dependents.Holders {
		dependents.Holders[i] = "/dev/" + dependents.Holders[i]
	}

	return dependents, nil
}
