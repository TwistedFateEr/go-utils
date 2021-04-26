package addrService

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"net"
)

func DeviceAndAddr()(map[netlink.Link][]net.Addr,error){

	ls,err:=netlink.LinkList()
	if err != nil {
		return nil,err
	}

	for _, l := range ls {
		if l.Type() == "device"{
			fmt.Println(l)
		}
	}


	return nil,nil
}