package topology

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// deviceSysPath has the a device name and its syspath
type deviceSysPath struct {
	SysPath    string
	DeviceName string
}

// getParent gets the parent of this device if it has parent
func (s deviceSysPath) getParent() (string, bool) {
	parts := strings.Split(s.SysPath, "/")

	parent := parts[len(parts)-2]

	if !isParent(parent) {
		return "", false
	}

	return parent, true
}

// getPartitions gets the partitions of this device if it has any
func (s deviceSysPath) getPartitions() ([]string, bool) {
	/*
		if partition file has value 1, can return from there itself
		partitionPath := s.SysPath + "/partition"
		if _, err := os.Stat(partitionPath); os.IsNotExist(err) {

		}*/
	partitions := make([]string, 0)

	files, err := ioutil.ReadDir(s.SysPath)
	if err != nil {
		return nil, false
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), s.DeviceName) {
			partitions = append(partitions, file.Name())
		}
	}

	return partitions, true
}

// getHolders gets the devices that are held by this device
func (s deviceSysPath) getHolders() ([]string, bool) {
	holderPath := s.SysPath + "/holders"
	holders := make([]string, 0)

	// check if holders are available for this device
	if _, err := os.Stat(holderPath); os.IsNotExist(err) {
		return nil, false
	}

	files, err := ioutil.ReadDir(holderPath)
	if err != nil {
		return nil, false
	}

	for _, file := range files {
		holders = append(holders, file.Name())
	}
	return holders, true
}

// getSlaves gets the devices to which this device is a slave. Or, the devices
// which holds this device
func (s deviceSysPath) getSlaves() ([]string, bool) {
	slavePath := s.SysPath + "/slaves"
	slaves := make([]string, 0)

	// check if slaves are available for this device
	if _, err := os.Stat(slavePath); os.IsNotExist(err) {
		return nil, false
	}

	files, err := ioutil.ReadDir(slavePath)
	if err != nil {
		return nil, false
	}

	for _, file := range files {
		slaves = append(slaves, file.Name())
	}
	return slaves, true
}

func isParent(dir string) bool {

	// if the dir path is block or nvme instance it means the given dir is
	// not a parent

	// check for block
	if dir == "block" {
		return false
	}

	// check for nvme instance
	nvmeInstanceRegex := "nvme[0-9]+$"
	r := regexp.MustCompile(nvmeInstanceRegex)
	if r.MatchString(dir) {
		return false
	}

	return true
}
