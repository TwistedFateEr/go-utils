package netlinkService

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func AddVlan(linkname string, vid uint16) error {
	master, err := netlink.LinkByName(linkname)
	if err != nil {
		return fmt.Errorf("link err :%v",err)
	}

	vname := fmt.Sprintf("%v.%v", linkname, vid)

	vl := &netlink.Dummy{LinkAttrs: netlink.LinkAttrs{Name: vname}}

	err = netlink.LinkAdd(vl)
	if err != nil {
		// link exists
		fmt.Println("link exists   :",err,vl)
	}

	err = netlink.LinkSetMaster(vl, master)
	if err != nil {
		return fmt.Errorf("set master err:5V",err)
	}

	err = netlink.BridgeVlanAdd(vl, vid, true, true, false, false)
	if err != nil {
		return fmt.Errorf(" addr err:%v",err)
	}

	return netlink.LinkSetUp(vl)
}

func DelVlan(linkname string, vid uint16) error {
	tl, err := netlink.LinkByName(linkname)
	if err != nil {
		return err
	}

	err = netlink.BridgeVlanDel(tl, vid, true, true, false, false)
	if err != nil {
		return err
	}

	return nil
}

func FindAllVlan() error {
	fs, err := netlink.BridgeVlanList()
	if err != nil {
		return err
	}

	for i, infos := range fs {
		fmt.Println(i, infos)
	}

	return nil
}
