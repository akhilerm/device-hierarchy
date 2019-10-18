## Device Topology

A simple package that parses the sysfs directory to obtain the blockdevice topology.

The project is still in **Alpha** stage.

Usage:
- To use as a binary, build the code
```
go build -o topo main.go
./topo /dev/sdb
```

- To use as a package. Import the `topology` package and create a `Device` object using the device path.
This will be used to get the dependent devices.

Terms used: 
1. `Holders` : Holders are the list of blockdevices that are held by the given blockdevice.
eg: `/dev/sda1` is held by `/dev/sda` since it is a partition. Therefore the holder list of `/dev/sda` will have `/dev/sda1`
2. `Slaves` : Slaves are the list of blockdevices to which the given device acts as a slave.
eg: `/dev/sda1` is a slave for `/dev/sda`. Therefore slave list of `/dev/sda1` will contain `/dev/sda`
