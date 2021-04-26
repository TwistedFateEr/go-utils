package main

import (
	"github.com/vishvananda/netlink"
)


var (
	linkName = "tap0"
)

func main() {
	link := &netlink.Dummy{LinkAttrs:netlink.LinkAttrs{
		Name: linkName,
	}}
	var  err error

	err = netlink.LinkAdd(link)
	if err != nil {
		panic(err)
	}

	err = netlink.LinkSetUp(link)
	if err != nil {
		panic(err)
	}
}